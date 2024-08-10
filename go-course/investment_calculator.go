package main

import (
	"fmt"
	"math"
)

func calculate () {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64;

	fmt.Print("Enter The Investement Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter How Many Years You Want To Invest: ")
	fmt.Scan(&years)

	fmt.Print("Enter The Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100 , years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(`The future value is: `, futureValue)
	fmt.Println(`The real future value is: `, futureRealValue)
}