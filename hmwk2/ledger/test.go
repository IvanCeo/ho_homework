package main

import (
	"fmt"
	"time"
)

func Test() {
	storage := NewStorage()

	err := storage.AddTransaction(Transaction{
		Amount:      1500,
		Category:    "Food",
		Description: "Lunch at cafe",
		Date:        time.Now(),
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = storage.AddTransaction(Transaction{
		Amount:      3200,
		Category:    "Transport",
		Description: "Taxi ride",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = storage.AddTransaction(Transaction{
		Amount:      7000,
		Category:    "Shopping",
		Description: "Bought new headphones",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("=== Transactions list ===")
	for _, tx := range storage.ListTransactions() {
		fmt.Printf("ID: %d | Amount: %d | Category: %s | Description: %s | Date: %s\n",
			tx.ID, tx.Amount, tx.Category, tx.Description, tx.Date.Format(time.RFC3339))
	}
}