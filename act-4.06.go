package main

import "fmt"

type MyStruct struct {
	name string
}

func printTyper(value interface{}) {

	switch t := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		fmt.Printf("%v is of type Integer\n", t)
	case float32, float64:
		fmt.Printf("%v is of type Float\n", t)
	case string:
		fmt.Printf("%v is of type String\n", t)
	case bool:
		fmt.Printf("%v is of type Bool\n", t)
	default:
		fmt.Printf("%v is of type Unknown\n", t)

	}

}

func main234() {
	arr := [5]any{1, 22.3, "123244", false, MyStruct{"Oussama Tailba"}}

	for _, value := range arr {
		printTyper(value)
	}
}
