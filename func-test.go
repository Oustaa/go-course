// package main

// import "fmt"

// func main() {
// 	itemsSold()

// 	myName, myAge := greeting()

// 	fmt.Println(myName, myAge)

// 	message := "Hello I am Oussama"

// 	f := func(str string) {
// 		fmt.Println(str)
// 	}

// 	f(message)

// 	withCallback(func() {
// 		fmt.Println("I am Oussama Tailba, and i am a good develpper")
// 	})
// }

// func greeting() (name string, age int) {
// 	name = "Oussama Tailba"
// 	age = 27

// 	return
// }

// func itemsSold() {
// 	items := map[string]int{
// 		"Oussama": 23,
// 	}
// 	items["kaoutar"] = 506
// 	items["Khalid"] = 55

// 	for k, v := range items {
// 		fmt.Printf("%s sold %d items, and ", k, v)

// 		if v < 40 {
// 			fmt.Println("Under")
// 		} else if v > 40 && v < 100 {
// 			fmt.Println("Meet it")
// 		} else {
// 			fmt.Println("Exceeded")
// 		}
// 	}

// 	fmt.Printf("%T\n", [3]int{})
// 	fmt.Printf("%T\n", []int{})
// }

// func withCallback(callback func()) {
// 	callback()
// }

package main

import "fmt"

// func main() {
// 	i := 0
// 	incrementor := func() int {
// 		i += 1
// 		return i
// 	}
// 	fmt.Println(incrementor())
// 	fmt.Println(incrementor())
// 	i += 10
// 	fmt.Println(incrementor())
// }

// func main() {
// 	increment := incrementor()

// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

// func incrementor() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

func main() {
	x := decriment(34)

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func decriment(counter int) func() int {
	return func() int {
		counter--
		return counter
	}
}
