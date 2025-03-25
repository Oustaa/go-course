package main

import "fmt"

type Person82 struct {
	name  string
	age   int8
	email string
}

func main2() {
	// var myArr [3]int = [3]int{1, 2, 3}
	// arr2 := [6]string{"Oussama", "kaoutar", "Moad", "Khadija", "Samir", "Ihsan"}
	// arr3 := [...]int{12: 34, 0: 2, 9: 230, 200: 9}

	// fmt.Printf("%#v\n", myArr)
	// fmt.Printf("%#v\n", arr2)
	// fmt.Printf("%#v\n", arr3)

	// intArr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := 0; i < len(intArr); i++ {
	// 	intArr[i] *= intArr[i]
	// }
	// fmt.Printf("%#v\n", intArr)

	people := [10]Person82{
		{"Oussama Tailba", 27, "otailaba98@gmail.com"},
		{"kaoutar Taki", 22, "ktaki00@gmail.com"},
	}
	fmt.Printf("%#v\n", people)

}
