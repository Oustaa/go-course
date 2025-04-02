package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main2344() {
	f, err := os.OpenFile("errors.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	Logger := log.New(f, "", log.Ldate|log.Ltime|log.Llongfile)

	name := "Thanos"

	defer f.Close()
	defer func() {
		if r := recover(); r != nil {
			Logger.Println("SOME FUCKING ERROR HAPPENED")
		}

	}()

	intToString := strconv.Itoa(23.3)
	fmt.Printf("%#v\n", intToString)

	Logger.Println("Demo app")
	Logger.Printf("%s is here!", name)
	Logger.Panicln("WHAAT THE FUCK HAPPENED")
	Logger.Print("Run")
}
