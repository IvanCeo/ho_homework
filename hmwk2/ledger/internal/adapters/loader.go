package adapters

import (
	"fmt"
	"io"
	"encoding/json"

	"ledger/pkg/domain"
)

func (s *Storage) LoadBudgets(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("Data reading error: %w", err)
	}

	var budgets []domain.Budget
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