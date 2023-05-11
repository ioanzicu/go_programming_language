// concurrency-safe bank with one account
package main

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int { // read the current balance
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance: // send current balance
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
