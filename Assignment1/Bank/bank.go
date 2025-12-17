package Bank

import "fmt"

type BankAccount struct {
	AccountNumber string
	AccountOwner  string
	Balance       float64
	Transactions  []string
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount should be more than 0")
		return
	}

	b.Balance += amount
	b.Transactions = append(b.Transactions, fmt.Sprintf("Deposit: %.2f", amount))
	fmt.Println("Deposited successfully")
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdraw amount should be more than 0")
		return
	}

	if b.Balance < amount {
		fmt.Println("Withdraw amount should be more than Balance")
		return
	}

	b.Balance -= amount
	b.Transactions = append(b.Transactions, fmt.Sprintf("Withdraw: %.2f", amount))
	fmt.Println("Withdrawed successfully")
}

func (b *BankAccount) GetBalance() float64 {
	fmt.Printf("Balance: %.2f\n", b.Balance)
	return b.Balance
}
