package domain

// Validate находится в домене, потому что
// это не бизнес логика, а логика корректности данных

import (
	"time"
)

type Transaction struct {
	ID int64
	Amount int64
	Category string
	Description string
	Date time.Time
}

const sinceYear int = 2023

func (t *Transaction) Validate() error {
	if t.Amount == 0 {
		return ErrZeroAmount
	}
	if t.Category == "" {
		return ErrEmptyCategory
	}
	// дается разница в 1 милисекунду, потому что при создании транзакции если поле Date не задано явно, присваивается time.Now()
	if time.Now().Before(t.Date.Add(-1)) && t.Date.Before(time.Date(sinceYear, time.January, 1, 0, 0, 0, 0, time.UTC)){
		return ErrDate
	}
	return nil
}