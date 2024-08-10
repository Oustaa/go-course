package main

import (
	"fmt"
	"os"
)

func main_test() {
	revunue := getData("Enter Revunue: ")
	validateAndExit(revunue)
	expenses := getData("Enter Expenses: ")
	validateAndExit(expenses)
	taxRate := getData("Enter Tax Rate: ")
	validateAndExit(taxRate)

	ebt := revunue - expenses
	profit := float64(ebt) * (1 - taxRate / 100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

	valueToStore := fmt.Sprintf("\n%.2f,%.2f,%d,%.2f,%.2f,%.2f", revunue, expenses, int(taxRate), ebt,profit,ratio)

	os.WriteFile("profit.csv", []byte(valueToStore), 0644)

}

func getData(message string) float64 {
	var value float64;
	fmt.Print(message)
	fmt.Scan(&value)

	return value;
}

func validateAndExit(value float64) {
	if value <= 0 {
        fmt.Println("Invalid value. Please enter a positive number.")
        os.Exit(1)
    }
}