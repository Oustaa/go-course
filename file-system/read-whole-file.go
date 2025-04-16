// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	content, err := os.ReadFile("clean-up.go")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("File contents: ")

// 	fmt.Println(string(content))
// }

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ma23in() {
	f, err := os.Open("clean-up.go")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
	fmt.Println()
}
