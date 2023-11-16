package banking

import "fmt"

type Account interface {
	Accrue(rate float64, resultChan chan<- float64)
	Balance() float64
	Deposit(amount float64)
	Withdraw(amount float64)
	String() string
}

type account struct {
	number   string
	customer Customer
	balance  float64
}

func (a *account) Accrue(rate float64, resultChan chan<- float64) {
	interest := a.balance * rate
	a.balance += interest
	resultChan <- interest
}

func (a *account) Balance() float64 {
	return a.balance
}

func (a *account) Deposit(amount float64) {
	a.balance += amount
}

func (a *account) Withdraw(amount float64) {
	a.balance -= amount
}

func (a *account) String() string {
	return fmt.Sprintf("%s:%v:%.2f", a.number, a.customer, a.balance)
}
