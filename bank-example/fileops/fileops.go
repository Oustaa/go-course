package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const BalenceFileName = "balance.txt"

func WriteBalanceToFile(balance float64) {
	os.WriteFile(BalenceFileName, []byte(fmt.Sprintf("%.2f", balance)), 0644)
}

func ReadBalanceFromFile() (float64, error)  {
	data, _ := os.ReadFile(BalenceFileName)
	dataText := string(data)

	balance, error := strconv.ParseFloat(dataText, 64)

	if error != nil {
		fmt.Println("An error occured: ", error)
		return 0.0, errors.New("an error occured while reding the file")
	}

	return balance, nil
}