package domain

import (
	"errors"
)

var ErrZeroAmount = errors.New("Zero amount error")
var ErrEmptyCategory = errors.New("Empty category error")
var ErrDate = errors.New("Date value error")
var ErrNegativeLimit = errors.New("Negative limit error")

var ErrBudgetExceeded = errors.New(": budget exceeded")
