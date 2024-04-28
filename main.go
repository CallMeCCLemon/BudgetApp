package main

import (
	"BudgetingApp/service"
	"log"
)

func main() {
	err := service.Run()
	if err != nil {
		log.Fatal("Received error from gin server!", err)
	}
}
