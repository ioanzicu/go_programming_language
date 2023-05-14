// This package provides a concurrency-unsafe
// memoization of a function of type Func.
package memo5

// A request is a message requesting that the Func be applied to key
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	requests chan request
}

// Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Concurrent, duplicate-suppressing, non-blocking cache, alternative to memo4
func (memo *Memo) Get(key string) (value interface{}, err error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

/*
The cache variable is confined to the monitor goroutine (*Memo).server, shown below. The
monitor reads requests in a loop until the request channel is closed by the Close method. For
each request, it consults the cache, creating and inserting a new entry if none was found.
*/
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e

			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
		/*
			The call and deliver methods must be called in their own goroutines
			to ensure that the monitor goroutine does not stop processing new requests.
		*/
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition
	<-e.ready // the cannel to be close
	// Send the result to the client
	response <- e.res
}
