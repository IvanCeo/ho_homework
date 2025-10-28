package domain

// Validate находится в домене, потому что
// это не бизнес логика, а логика корректности данных

type Budget struct {
	Category string
	Limit int64
	Period int64
}

func (b *Budget) Validate() error {
	if b.Limit < 0 {
		return ErrNegativeLimit
	}
	if b.Category == "" {
		return ErrEmptyCategory
	}
	return nil
}