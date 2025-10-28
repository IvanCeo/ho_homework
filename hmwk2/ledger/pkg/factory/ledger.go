package factory

import (
	"ledger/pkg/usecase"
	"ledger/internal/adapters"
)

func NewLedgerFactory() *usecase.Ledger {
    var st usecase.Storage = adapters.NewStorage()
    return usecase.NewLedger(st)
}