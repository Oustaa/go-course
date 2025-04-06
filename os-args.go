// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	args := os.Args
// 	if len(args) < 2 {
// 		fmt.Println("Usage: go run main.go <name>")
// 		return
// 	}
// 	name := args[1]

// 	greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", name)
// 	fmt.Println(greeting)
// }

package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Oussama", "Name of the person to say hello to")
	quietFlag = flag.Bool("quiet", false, "Toggle to be quiet when saying hello")
)

func mai12344n() {
	flag.Parse()
	if !*quietFlag {
		greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", *nameFlag)
		fmt.Println(greeting)
	}
}
