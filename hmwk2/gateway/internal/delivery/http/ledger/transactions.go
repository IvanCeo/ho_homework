package ledger

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"ledger/pkg/domain"
	"ledger/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	server *usecase.Ledger
}

type CreateTransactionRequest struct {
	Amount      string    `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date"`
}

type TransactionResponse struct {
	ID          int64     `json:"id"`
	Amount      int64     `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date"`
}

func CreateTransactionRequestToDomainTransaction(req *CreateTransactionRequest) (*domain.Transaction, error) {
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil {
		return nil, err
	}

	// d, err := time.Parse(time.RFC3339, req.Date)
	// if err != nil {
	// 	return nil, err
	// }

	return &domain.Transaction{
		Amount:      amount,
		Category:    req.Category,
		Description: req.Description,
		Date:        req.Date,
	}, nil
}

func DomainTransactionToTransactionResponse(dto *domain.Transaction) *TransactionResponse {
	return &TransactionResponse{
		ID:          dto.ID,
		Amount:      dto.Amount,
		Category:    dto.Category,
		Description: dto.Description,
		Date:        dto.Date,
	}
}

func NewHandler(server *usecase.Ledger) *Handler {
	return &Handler{
		server: server,
	}
}

func (h *Handler) CreateTransactionHandle(c *fiber.Ctx) error {
	tReq := &CreateTransactionRequest{}
	if err := c.BodyParser(&tReq); err != nil {
		return c.Status(400).SendString(fmt.Sprintf(`{"error":"%v"}`, err))
	} // TODO: json создать структуру ошибки

	req, err := CreateTransactionRequestToDomainTransaction(tReq)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf(`{"error":"%v"}`, err))
	}

	err = h.server.AddTransaction(*req)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrZeroAmount):
			return c.Status(400).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
		case errors.Is(err, domain.ErrEmptyCategory):
			return c.Status(400).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
		case errors.Is(err, domain.ErrDate):
			return c.Status(400).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
		case errors.Is(err, domain.ErrBudgetExceeded):
			return c.Status(409).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
		default:
			return c.Status(500).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
		}
	}

	res := DomainTransactionToTransactionResponse(req)
	j, _ := json.Marshal(res)

	return c.Status(201).Send(j)
}

func (h *Handler) GetTransactionsHandle(c *fiber.Ctx) error {
	transactions := h.server.ListTransactions() //[]domain.Transaction
	var res []*TransactionResponse
	for _, row := range transactions {
		res = append(res, DomainTransactionToTransactionResponse(&row))
	}
	return c.JSON(res)
}
