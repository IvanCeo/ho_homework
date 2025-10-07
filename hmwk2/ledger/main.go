package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID int64
	Amount int64
	Category string
	Description string
	Date time.Time
}

type Storage struct {
	transactions []Transaction
}

func NewStorage() *Storage {
	return &Storage{
		transactions: make([]Transaction, 0),
	}
}

func (s *Storage) AddTransaction(tx Transaction) error {
	if tx.Amount == 0 {
		return fmt.Errorf("transaction amount cannot be zero")
	}

	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	tx.ID = int64(len(s.transactions) + 1)
	s.transactions = append(s.transactions, tx)

	return nil
}

func (s *Storage) ListTransactions() []Transaction {
	res := make([]Transaction, len(s.transactions))
	copy(res, s.transactions)
	return res
}

func main() {
	fmt.Println("started")
	Test()
}