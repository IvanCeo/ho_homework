package domains

import (
	"time"
	"fmt"
)

type Transaction struct {
	ID int64			`json:"id"`
	Amount int64		`Json:"amount"`
	Category string		`json:"category"`
	Description string	`json:"description"`
	Date time.Time		`json:"date"`
}

// бизнес логика - тразакции не могут быть раньше заданного года (конкретного вермени, при необходимости) и не могут быть из будущего
const sinceYear int = 2023

func (t *Transaction) Validate() error {
	if t.Amount == 0 {
		return fmt.Errorf("Zero amount error")
	}
	if t.Category == "" {
		return fmt.Errorf("Empty category error")
	}
	// дается разница в 1 милисекунду, потому что при создании транзакции если поле Date не задано явно, присваивается time.Now()
	if time.Now().Before(t.Date.Add(-1)) && t.Date.Before(time.Date(sinceYear, time.January, 1, 0, 0, 0, 0, time.UTC)){
		return fmt.Errorf("Date error")
	}
	return nil
}