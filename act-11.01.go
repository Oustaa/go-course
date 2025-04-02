package main

import (
	"errors"
	"fmt"
	_ "log"
	_ "os"
	"strings"
)

var (
	ErrInvalidSSNLength     = errors.New("SSN Length is invalid")
	ErrInvalidSSNNumbers    = errors.New("SSN Must has at least one digit")
	ErrInvalidSSNPrefix     = errors.New("SSN Should not be prefixed with 000")
	ErrInvalidSSNDigitPlace = errors.New("SSN Should not be prefixed with 000")
)

func maiwwwn() {
	// Logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)

	name := "Oussama Tailba is me"

	newString := strings.Replace(name, "a", "//", 2)

	fmt.Println(newString)
}

func ssnLength(ssn string) error {
	if len(ssn) < 9 {
		return ErrInvalidSSNLength
	}

	return nil
}
