package domains

type Budget struct {
	Category string	`json:"category"`
	Limit int64		`json:"limit"`
	Period int64 	`json:"period"`
}