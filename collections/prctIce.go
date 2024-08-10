package main

import "fmt"


type Product struct {
	id int
	title string
	price float32
}

func main2() {
	hobbies := [3]string{"gaming", "coding", "reading"}

	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	courseGools := []string{"go", "python"}
	courseGools[1] = "python3"
	courseGools = append(courseGools, "java")
	fmt.Println(courseGools)

	productList := []Product{
		{
			id: 1,
			title: "book",
			price: 99.99,
		},
		{
			id: 2,
			title: "mouse g502 Hero",
			price: 19.99,
		},
	}

	fmt.Println(productList)
}
