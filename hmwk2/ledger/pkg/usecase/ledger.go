package usecase

import (
	"io"
	"ledger/pkg/domain"
)

type Storage interface {
	LoadBudgets(io.Reader) error
	SetBudget(domain.Budget) error
	AddTransaction(domain.Transaction) error
	ListTransactions() []domain.Transaction
	ListBudgets() []domain.Budget
}

type Ledger struct {
	storage Storage
}

func NewLedger(s Storage) *Ledger {
	return &Ledger{storage: s}
}

func (l *Ledger) SetBudget(bd domain.Budget) error {
	err := l.storage.SetBudget(bd)
	if err != nil {
		return err
	}
	return nil
}

func (l *Ledger) AddTransaction(tx domain.Transaction) error {
	err := l.storage.AddTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

func (l *Ledger) ListTransactions() []domain.Transaction {
	return l.storage.ListTransactions()
}

func (l *Ledger) ListBudgets() []domain.Budget {
	return l.storage.ListBudgets()
}
