// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f, err := os.Create("test.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		fmt.Println("DEFER WAS CALLED")
// 		f.Close()
// 	}()

// 	f.WriteString("Ousta is a foftware enginier\n")
// 	f.Write([]byte("Ousta is a foftware enginier\n"))
// 	os.WriteFile("test.txt", []byte("Oussama is the fucking goat"), 0644)
// }

package main

import (
	"flag"
	"fmt"
	"os"
)

func mdain() {
	// var name string
	// flag.StringVar(&name, "name", "", "File name")

	nameFlag := flag.String("name", "", "File name")
	flag.Parse()

	name := *nameFlag

	file, err := os.Stat(name)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File does not exist!\n", name)
			fmt.Println(file)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("file name: %s\nIsDir: %t\nModTime: %v\nMode: %v\nSize: %d\n", file.Name(),
		file.IsDir(), file.ModTime(), file.Mode(), file.Size())
}
