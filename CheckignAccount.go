package banking

type CheckingAccount struct {
	account
}

func NewCheckingAccount(number string, customer Customer, balance float64) *CheckingAccount {
	return &CheckingAccount{
		account: account{
			number:   number,
			customer: customer,
			balance:  balance,
		},
	}
}
