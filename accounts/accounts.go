package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner}

	return &account
}
