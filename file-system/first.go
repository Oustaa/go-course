package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type FlagFuncs interface {
	Bool(name string, value bool, usage string) *bool
	Duration(name string, value time.Duration, usage string) *time.
		Duration
	Float64(name string, value float64, usage string) *float64
	Int(name string, value int, usage string) *int
	Int64(name string, value int64, usage string) *int64
}

func mai2n() {
	var (
		debug    bool
		duration time.Duration
		price    float64
		age      int
	)

	name := flag.String("name", "", "you name here")

	flag.BoolVar(&debug, "debug", false, "this is to enable debuging")
	flag.DurationVar(&duration, "duration", 0, "the duration you have run")
	flag.Float64Var(&price, "price", 0.0, "The price of the product")
	flag.IntVar(&age, "age", 21, "Your age")
	flag.Parse()

	fmt.Printf("your name is %s\n", *name)
	fmt.Printf("you want debug mode %t\n", debug)
	fmt.Printf("you have run for %v\n", duration)
	fmt.Printf("you age is %d\n", age)
	fmt.Printf("the product's price you buyed is %.2f\n", price)

	f, err := os.OpenFile("file2.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("There is an error while opening the ffile.txt file")
		os.Exit(1)
	}
	defer f.Close()

	output := fmt.Sprintf("Name: %s\nDebug: %t\nDuration: %v\nAge: %d\nPrice: %.2f\n\n",
		*name, debug, duration, age, price)

	_, err = f.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	f.Write([]byte{'o', 'u', 's', 's', 'a', 'm', 'a'})
}
