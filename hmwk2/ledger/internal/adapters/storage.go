package adapters

import (
	"fmt"
	"time"

	"ledger/pkg/domain"
)

type Storage struct {
	transactions []domain.Transaction
	budgets      map[string]domain.Budget
}

func NewStorage() *Storage {
	return &Storage{
		transactions: make([]domain.Transaction, 0),
		budgets:      make(map[string]domain.Budget, 0),
	}
}

type Validatable interface {
	Validate() error
}

func CheckValid(v Validatable) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) SetBudget(bd domain.Budget) error {
	if err := CheckValid(&bd); err != nil {
		return err
	}
	s.budgets[bd.Category] = bd
	fmt.Printf("Budget %s successfully setted\n", bd.Category)
	return nil
}

func (s *Storage) AddTransaction(tx domain.Transaction) error {
	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	if err := CheckValid(&tx); err != nil {
		return err
	}

	if budget, ok := s.budgets[tx.Category]; ok {
		var sum int64
		for _, t := range s.transactions {
			if t.Category == tx.Category {
				sum += t.Amount
			}
		}
		if (sum + tx.Amount) > budget.Limit {
			return fmt.Errorf("%s%w", tx.Category, domain.ErrBudgetExceeded)
		}
	}

	tx.ID = int64(len(s.transactions) + 1)
	s.transactions = append(s.transactions, tx)

	return nil
}

func (s *Storage) ListTransactions() []domain.Transaction {
	res := make([]domain.Transaction, len(s.transactions))
	copy(res, s.transactions)
	return res
}

func (s *Storage) ListBudgets() []domain.Budget {
	var res []domain.Budget
	for _, row := range s.budgets {
		res = append(res, row)
	}
	return res
}
