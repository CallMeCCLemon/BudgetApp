package main

import (
	"BudgetingApp/service"
	"log"
)

func main() {
	err := service.Start()
	if err != nil {
		log.Fatal("Received error from gin server!", err)
	}
}
