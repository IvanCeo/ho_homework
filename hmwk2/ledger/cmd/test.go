package main

import (
	"fmt"
	"time"
	// "os"

	"ledger/pkg/domain"
	"ledger/pkg/factory"
)

type CreateBudgetRequest = domain.Budget
type CreateTransactionRequest = domain.Transaction

func Test() {
	fmt.Println("test started")

	st := factory.NewLedgerFactory()

	err := st.SetBudget(CreateBudgetRequest{
		Category: "Food",
		Limit: 3000,
		Period: 13221892,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	// f, err := os.Open("budgets.json")
	// if err != nil {
	// 	fmt.Printf("Open error: %v", err)
	// 	return
	// }
	// defer f.Close()

	// st := d.NewStorage()
	
	// if err := st.LoadBudgets(f); err != nil {
	// 	fmt.Printf("LoadBudgets error: %v", err)
	// }

	err = st.SetBudget(CreateBudgetRequest{
		Category: "Food",
		Limit: 3000,
		Period: 13221892,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = st.AddTransaction(CreateTransactionRequest{
		Amount:      1500,
		Category:    "Food",
		Description: "Lunch at cafe",
		Date:        time.Now(),
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = st.AddTransaction(CreateTransactionRequest{
		Amount:      3200,
		Category:    "Transport",
		Description: "Taxi ride",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = st.AddTransaction(CreateTransactionRequest{
		Amount:      7000,
		Category:    "Shopping",
		Description: "Bought new headphones",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = st.AddTransaction(CreateTransactionRequest{
		Amount:      2000,
		Category:    "Food",
		Description: "Dinner in restaurant",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Transactions list ===")
	for _, tx := range st.ListTransactions() {
		fmt.Printf("ID: %d | Amount: %d | Category: %s | Description: %s | Date: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format(time.RFC3339))
	}
}