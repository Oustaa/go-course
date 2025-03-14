package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Local struct {
	language string
	country  string
}

func isLocalSupported(local Local, locals []Local) bool {
	for i := 0; i < len(locals); i++ {
		if local == locals[i] {
			return true
		}
	}

	return false
}

func passwordChecker(psw string) (string, error) {
	if len(psw) < 8 {
		return "", errors.New("Your password must be 8 character long")
	}
	return psw, nil
}

func printArgsLen(numsArr ...int) {
	fmt.Printf("The number of intered args is %v\n", len(numsArr))
}

func printAny(args ...any) {
	fmt.Println(args...)
}

func main() {
	locals := []Local{{"en", "US"}, {"en", "CN"}, {"fr", "CN"}, {"fr", "FR"}, {"ru", "RU"}}

	if len(os.Args) < 2 {
		fmt.Println("Please enter a local")
		os.Exit(1)
	}

	localParts := strings.Split(os.Args[1], "_")

	if len(localParts) != 2 {
		fmt.Println("The passed local format is invalid")
		os.Exit(1)
	}

	passedLocalStruct := Local{country: localParts[1], language: localParts[0]}

	if !isLocalSupported(passedLocalStruct, locals) {
		fmt.Println("The passed local is unsuported")
		os.Exit(1)
	}

	fmt.Println("The passsed local is supported")

	printArgsLen(1, 2, 3, 4, 5, 6, 7)
	printArgsLen()

	fmt.Printf("string len is %v\n", len("Oussama tailba"))

	printAny(1, 2, 32, 34, 45, 6, 546, 5462346, 236546, 2456)

	psw, err := passwordChecker("1233456677passwordChecker")
	if err == nil {
		fmt.Printf("Passowrd %s is valid\n", psw)
	} else {
		fmt.Println(err)
	}

	psw, err = passwordChecker("3234")
	if err == nil {
		fmt.Printf("Passowrd %s is valid\n", psw)
	} else {
		fmt.Println(err)
	}
}
