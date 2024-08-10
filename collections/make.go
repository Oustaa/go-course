package main

import "fmt"

func main4() {
	userNames := make([]string, 2,3)

	userNames[0] = "ktaki"
	userNames[1] = "ousta6"

	userNames = append(userNames, "Goo", "Ts", "js", "node", "c", "c++")

	for index, value := range userNames{
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
		fmt.Println("---------------")
	}
}