package main

import (
	"fmt"

	"example.com/pointers/ages"
)


func main( ) {
  age := 26

  fmt.Println("Age: ", age)

  ages.EditAgeToAdultAge(&age)
  fmt.Println("Adult age are: ", age)
  fmt.Println("Age after: ", age)
}

