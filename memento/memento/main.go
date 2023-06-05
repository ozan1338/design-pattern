package main

import "log"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*Memento, *BankAccount) {
	return &Memento{balance}, &BankAccount{balance}
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func main() {
	m0, ba := NewBankAccount(100)

	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)

	log.Println("Total deposit ", ba)

	ba.Restore(m1)
	log.Println(ba)

	ba.Restore(m2)
	log.Println(ba)

	ba.Restore(m0)
	log.Println(ba)
}
