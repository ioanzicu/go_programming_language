// A semaphore that counts only to 1 is called a binary semaphore.
package main

var (
	sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int { // read the current balance
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
