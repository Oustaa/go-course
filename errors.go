// package main

// import (
// 	"errors"
// 	"fmt"
// )

// var (
// 	ErrHourlyRate  = errors.New("invalid hourly rate")
// 	ErrHoursWorked = errors.New("invalid hours worked per week")
// )

// func main() {
// 	pay, err := payDay(81, 50)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	pay, err = payDay(80, 5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	pay, err = payDay(80, 50)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(pay)
// }
// func payDay(hoursWorked, hourlyRate int) (int, error) {
// 	if hourlyRate < 10 || hourlyRate > 75 {
// 		return 0, ErrHourlyRate
// 	}
// 	if hoursWorked < 0 || hoursWorked > 80 {
// 		return 0, ErrHoursWorked
// 	}
// 	if hoursWorked > 40 {
// 		hoursOver := hoursWorked - 40
// 		overTime := hoursOver * 2
// 		regularPay := hoursWorked * hourlyRate
// 		return regularPay + overTime, nil
// 	}
// 	return hoursWorked * hourlyRate, nil
// }

package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func main() {
	pay := payDay(100, 25)
	fmt.Println(pay)
	pay = payDay(100, 200)
	fmt.Println(pay)
	pay = payDay(60, 25)
	fmt.Println(pay)
}

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n",
					hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n",
					hoursWorked, r)
			}
		}

	}()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)

	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	a()
// 	fmt.Println("This line will now get printed from main() function")
// }
// func a() {
// 	b("good-bye")
// 	fmt.Println("Back in function a()")
// }
// func b(msg string) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("error in func b()", r)
// 		}
// 	}()
// 	if msg == "good-bye" {
// 		panic(errors.New("something went wrong"))
// 	}
// 	fmt.Print(msg)
// }
