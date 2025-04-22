package main

import (
	_ "embed"

	"fmt"
)

//go:embed errors.txt
var errorText string

func main() {

	fmt.Println(errorText)

}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("errors.txt")
// 	if err != nil {
// 		fmt.Println("An error opening the file")
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Printf("%s\n", line)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("An error reading from the file")
// 		fmt.Printf("%v\n", err)
// 	}
// }
