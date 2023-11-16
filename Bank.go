package banking

import (
	"fmt"
	"sync"
)

type Bank struct {
	accounts map[Account]struct{}
	mutex    sync.Mutex
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[Account]struct{}),
	}
}

func (b *Bank) Add(account Account) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.accounts[account] = struct{}{}
}

func (b *Bank) Accrue(rate float64) {
	var totalInterest float64
	resultChan := make(chan float64)

	var wg sync.WaitGroup
	b.mutex.Lock()
	for account := range b.accounts {
		wg.Add(1)
		go func(acc Account) {
			defer wg.Done()
			acc.Accrue(rate, resultChan)
		}(account)
	}
	b.mutex.Unlock()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for interest := range resultChan {
		totalInterest += interest
	}

	fmt.Printf("Total interest accrued across all accounts: %.2f\n", totalInterest)
}

func (b *Bank) String() string {
	var result string
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for account := range b.accounts {
		result += account.String() + "\n"
	}
	return result
}

func main() {
	bank := banking.NewBank()

	customer := banking.NewCustomer("Ann")
	bank.Add(banking.NewCheckingAccount("01001", *customer, 100.00))
	bank.Add(banking.NewSavingAccount("01002", *customer, 200.00))

	bank.Accrue(0.02)
	fmt.Println(bank)
}