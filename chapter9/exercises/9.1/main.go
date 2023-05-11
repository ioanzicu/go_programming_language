/*
Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program.
The result should indicate whether the transaction succeeded or failed due
to insufficient funds. The message sent to the monitor goroutine must
contain both the amount to withdraw and a new channel over which the
monitor goroutine can send the boolean result back to Withdraw.
*/

// concurrency-safe bank with one account
package main

// var deposits = make(chan int) // send amount to deposit
// var withdraw = make(chan int) // send amount to withdraw
// var balances = make(chan int) // receive balance

// func Deposit(amount int) {
// 	deposits <- amount
// }

// func Balance() int { // read the current balance
// 	return <-balances
// }

// func Withdraw(amount int) bool {
// 	if <-balances-amount < 0 {
// 		return false
// 	}

// 	withdraw <- amount
// 	return true
// }

// func teller() {
// 	var balance int // balance is confined to teller goroutine
// 	for {
// 		select {
// 		case amount := <-deposits:
// 			balance += amount
// 		case amount := <-withdraw:
// 			balance -= amount
// 		case balances <- balance: // send current balance
// 		}
// 	}
// }

// func init() {
// 	go teller() // start the monitor goroutine
// }

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type Withdrawal struct {
	amount  int
	success chan bool
}

var withdrawals = make(chan Withdrawal) // send amount to withdraw

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int { // read the current balance
	return <-balances
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- Withdrawal{amount, ch}
	return <-ch
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdrawals:
			if withdraw.amount > balance {
				withdraw.success <- false
				continue
			}
			balance -= withdraw.amount
			withdraw.success <- true
		case balances <- balance: // send current balance
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
