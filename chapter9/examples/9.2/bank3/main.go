/*
Each time a goroutine accesses the variables of the bank (just balance here), it must call the
mutex’s Lock method to acquire an exclusive lock. If some other goroutine has acquired the
lock, this operation will block until the other goroutine calls Unlock and the lock becomes
available again. The mutex guards the shared variables.
*/
package main

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int { // read the current balance
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
