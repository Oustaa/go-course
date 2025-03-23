package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	firstName     string
	lastName      string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if strings.TrimSpace(d.lastName) == "" {
		panic(ErrInvalidLastName)
	}

}

func (d *directDeposit) report() {
	fmt.Println("*************************************************************************")
	fmt.Println("Last Name: ", d.lastName)
	fmt.Println("First Name: ", d.firstName)
	fmt.Println("Bank Name: ", d.bankName)
	fmt.Println("Routing NUmber: ", d.routingNumber)
	fmt.Println("Account Number: ", d.accountNumber)
}

func main() {
	account := directDeposit{"Abe", "", "XYZ Inc", 17, 1809}

	account.validateLastName()
	account.validateRoutingNumber()
	account.report()

}
