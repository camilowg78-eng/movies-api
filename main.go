package main

import (
	"fmt"
	"movies-api/collector"
)

func main() {

	books, err := collector.CollectorB()
	if err != nil {
		fmt.Println("Error collecting books:", err)
		return
	}

	for _, book := range books {
		fmt.Printf("Title: %s, Price: %.2f\n", book.Title, book.Price)
	}
}
