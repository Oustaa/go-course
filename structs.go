package main

import "fmt"

type Priceable struct {
	price                float32
	price_after_discount float32
}

type Product struct {
	name string
	Priceable
}

func mainwwe() {
	priceable := Priceable{price_after_discount: 99.99, price: 199.99}
	fmt.Printf("%#v\n", priceable)

	prod1 := Product{"book", Priceable{99.99, 89.99}}
	prod2 := Product{"book", Priceable{99.99, 89.99}}

	fmt.Println(prod1.price)
	fmt.Println(prod1.Priceable.price_after_discount)

	if prod1 == prod2 {
		fmt.Println("They are the same")
	} else {
		fmt.Println("They are not the fucking same")
	}
}
