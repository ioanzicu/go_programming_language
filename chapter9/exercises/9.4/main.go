/*
Construct a pipeline that connects an arbitrary number of goroutines with chan-
nels. What is the maximum number of pipeline stages you can create without running out of
memory? How long does a value take to transit the entire pipeline?
*/
package main

func pipeline(stages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < stages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}
