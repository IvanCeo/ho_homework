package domains

import "time"

type Transaction struct {
	ID int64
	Amount int64
	Category string
	Description string
	Date time.Time
}
