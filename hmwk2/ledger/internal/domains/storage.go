package domains

import (
	"fmt"
	"time"
	"io"
	"encoding/json"
)

type Validatable interface {
	Validate() error
}

func CheckValid(v Validatable) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}

type Storage struct {
	transactions []Transaction
	budgets map[string]Budget
}

func (s *Storage) LoadBudgets(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("Data reading error: %w", err)
	}

	var budgets []Budget
	if err := json.Unmarshal(data, &budgets); err != nil {
		return fmt.Errorf("Parsing error: %w", err)
	}

	for _, b := range budgets {
		if err := s.SetBudget(b); err != nil {
			return fmt.Errorf("in %s SetBudget error: %w", b.Category, err)
		}
	}
	return nil
}

func NewStorage() *Storage {
	return &Storage{
		transactions: make([]Transaction, 0),
		budgets: make(map[string]Budget, 0), // ключ это название категории, которая тоже есть в структуру budget
	}
}

func (s *Storage) SetBudget(bd Budget) error {
	if err := CheckValid(&bd); err != nil {
		return err
	}
	s.budgets[bd.Category] = bd
	fmt.Printf("Budget %s successfully setted\n", bd.Category)
	return nil
}

func (s *Storage) AddTransaction(tx Transaction) error {
	if tx.Date.IsZero() {
		tx.Date = time.Now()
	}

	if err := CheckValid(&tx); err != nil {
		return fmt.Errorf("AddTrasaction error: %w", err)
	}

	if budget, ok := s.budgets[tx.Category]; ok {
		var sum int64
		for _, t := range s.transactions {
			if t.Category == tx.Category {
				sum += t.Amount
			}
		}
		if (sum + tx.Amount) > budget.Limit {
			return fmt.Errorf("budget for %s exceeded", tx.Category)
		}
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