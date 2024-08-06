package main

import "fmt"

func entry() {
	prices := []float64{}

	prices = append(prices, 100.00, 45.02)
	prices = append(prices, 99.00)
	prices = append(prices, 49.0000001)


	fmt.Println(prices)
}