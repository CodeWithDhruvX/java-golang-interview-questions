package users

type Account struct {
	Owner   string
	balance int
}

func NewAccount(owner string) *Account {
	return &Account{
		Owner:   owner,
		balance: 0,
	}
}

func (a *Account) Deposit(amount int) {
	if amount > 0 {
		a.balance += amount
	}
}

func (a *Account) BalanceCount() int {
	return a.balance
}
