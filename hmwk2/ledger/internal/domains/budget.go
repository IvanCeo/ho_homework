package domains

import (
	"fmt"
)

type Budget struct {
	Category string	`json:"category"`
	Limit int64		`json:"limit"`
	Period int64 	`json:"period"`
}

func (b *Budget) Validate() error {
	if b.Limit < 0 {
		return fmt.Errorf("Negative limit error")
	}
	if b.Category == "" {
		return fmt.Errorf("Empty category error")
	}
	return nil
}