package ledger

import (
	"encoding/json"
	"fmt"
	"strconv"

	"ledger/pkg/domain"

	"github.com/gofiber/fiber/v2"
)

type CreateBudgetRequest struct {
	Category string `json:"category"`
	Limit    string `json:"limit"`
	Period   string `json:"period,omitempty"`
}

type BudgetResponse struct {
	Category string `json:"category"`
	Limit    int64  `json:"limit"`
	Period   int64  `json:"period,omitempty"`
}

func CreateBudgetRequestToDomainBudget(req *CreateBudgetRequest) (*domain.Budget, error) {
	limit, err := strconv.ParseInt(req.Limit, 10, 64)
	if err != nil {
		return nil, err
	}

	budget := &domain.Budget{
		Category: req.Category,
		Limit:    limit,
	}

	if req.Period != "" {
		budget.Period, err = strconv.ParseInt(req.Period, 10, 64)
		if err != nil {
			return nil, err
		}
		return budget, nil
	}
	return budget, nil
}

func DomainBudgetToCreateBudgetRequest(dto *domain.Budget) *BudgetResponse {
	return &BudgetResponse{
		Category: dto.Category,
		Limit:    dto.Limit,
		Period:   dto.Period,
	}
}

func (h *Handler) CreateBudgetHandle(c *fiber.Ctx) error {
	bReq := &CreateBudgetRequest{}
	if err := c.BodyParser(&bReq); err != nil {
		return c.Status(400).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
	}

	req, err := CreateBudgetRequestToDomainBudget(bReq)
	if err != nil {
		return c.Status(500).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
	}

	err = h.server.SetBudget(*req)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf(`{"error": "%v"}`, err))
	}

	res := DomainBudgetToCreateBudgetRequest(req)
	j, _ := json.Marshal(res)

	return c.Status(201).Send(j)
}

func (h *Handler) GetBudgetsHandle(c *fiber.Ctx) error {
	budgets := h.server.ListBudgets() //[]domain.Budgets
	var res []*BudgetResponse
	for _, row := range budgets {
		res = append(res, DomainBudgetToCreateBudgetRequest(&row))
	}
	return c.JSON(res)
}
