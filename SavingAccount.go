package banking

type SavingAccount struct {
	account
	interest float64
}

func NewSavingAccount(number string, customer Customer, balance float64) *SavingAccount {
	return &SavingAccount{
		account: account{
			number:   number,
			customer: customer,
			balance:  balance,
		},
		interest: 0,
	}
}

func (s *SavingAccount) Accrue(rate float64, resultChan chan<- float64) {
	interest := s.balance * rate
	s.interest += interest
	s.balance += interest
	resultChan <- interest
}
