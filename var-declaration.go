// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // type Person struct {
// // 	name         string
// // 	age_of_birth int
// // }

// // func getConfig() (bool, string, time.Time) {
// // 	return false, "Info", time.Now()
// // }

// // func incremant(num *int64) {
// // 	*num++
// // }

// // func main() {
// // 	var name, myType string = "Oussama", "Dev"
// // 	debug, logLevel, startUpTime := true, "Info", time.Now()
// // 	var اسمي string = "Oussama"

// // 	fmt.Println(name, myType)

// // 	fmt.Println(debug, logLevel, startUpTime)

// // 	fmt.Println(اسمي)

// // 	age := 27

// // 	if age < 18 {
// // 		fmt.Println("You are young")
// // 	} else if age <= 30 {

// // 		fmt.Println("About to be old asf")
// // 	} else {
// // 		fmt.Println("Expired")
// // 	}

// // 	var price float32 = 99.9987

// // 	fmt.Printf("%.2f, %T\n", price, price)
// // 	fmt.Printf("%#v\n", price)
// // 	fmt.Printf("%+v\n", price)

// // 	ousta := Person{"Oussama Tailba", 99}

// // 	fmt.Printf("%+v\n", ousta)
// // 	fmt.Println(ousta)

// // 	var num int64 = 100

// // 	fmt.Println(num)
// // 	incremant(&num)
// // 	incremant(&num)
// // 	fmt.Println(num)

// // 	var pp *int64 = nil

// // 	pp = &num

// // 	fmt.Println(*pp)

// // 	const (
// // 		Saturday = iota
// // 		_
// // 		_
// // 		_
// // 		_
// // 		_
// // 		Monday
// // 		Tuesday
// // 		Wednesday
// // 	)

// // 	fmt.Println(Monday)

// // }

// // package main

// // import "fmt"

// // func main() {
// // 	count := 5
// // 	var message string
// // 	if count > 5 {
// // 		message = "Greater than 5"
// // 	} else {
// // 		message = "Not greater than 5"
// // 	}
// // 	fmt.Println(message)
// // }

// // package main

// // import "fmt"

// // func main() {
// // 	count := 0
// // 	if count < 5 {
// // 		count = 10
// // 		count++
// // 	}
// // 	fmt.Println(count == 11)
// // }

// // package main

// // import "fmt"

// // func main() {
// // 	data := map[string]string{
// // 		"name": "Oussama Tailba",
// // 		"age":  "23",
// // 	}

// // 	for key, value := range data {

// // 		fmt.Printf("%s= %s\n", key, value)
// // 	}

// // }

// // package main

// // import "fmt"

// // func main() {
// // 	words := map[string]int{
// // 		"Gonna": 3,
// // 		"You":   3,
// // 		"Give":  2,
// // 		"Never": 1,
// // 		"Up":    4,
// // 	}

// // 	word := ""
// // 	count := 0

// // 	for key, value := range words {
// // 		if value > count {
// // 			word = key
// // 			count = value
// // 		}
// // 	}

// // 	fmt.Println("Most Populare word is: ", word)
// // 	fmt.Println("With a count of: ", count)

// // }

// // package main

// // import "fmt"

// // func main() {
// // for i := 1; i <= 100; i++ {

// // 	switch j := i; {
// // 	case j%3 == 0 && j%5 == 0:
// // 		fmt.Println("FizzBuzz")
// // 	case i%3 == 0:
// // 		fmt.Println("Fizz")
// // 	case i%5 == 0:
// // 		fmt.Println("Bizz")
// // 	default:
// // 		fmt.Println(i)
// // 	}
// // 	// if i%3 == 0 && i%5 == 0 {
// // 	// 	fmt.Println("FizzBuzz")
// // 	// } else if i%3 == 0 {
// // 	// 	fmt.Println("Fizz")

// // 	// } else if i%5 == 0 {
// // 	// 	fmt.Println("Buzz")

// // 	// } else {
// // 	// 	fmt.Println(i)

// // 	// }
// // }

// // arr := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6, -1}

// // fmt.Printf("Before %#v\n", arr)
// // for i := 0; i < len(arr)-1; i++ {
// // 	for j := 0; j < len(arr)-i-1; j++ {
// // 		if arr[j] > arr[j+1] {
// // 			arr[j], arr[j+1] = arr[j+1], arr[j]
// // 		}
// // 	}

// // }
// // fmt.Printf("After %#v\n", arr)
// // _goto(1223)

// // }

// // func _goto(number int16) {
// // 	if number < 123 {
// // 		goto RETURN
// // 	}

// // RETURN:
// // 	{
// // 		fmt.Println("RETURN SOMETHINE")
// // 	}
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"runtime"
// // )

// // func main() {
// // 	// var list []int
// // 	var list []int8
// // 	for i := 0; i < 1000000000; i++ {
// // 		list = append(list, 100)
// // 	}
// // 	var m runtime.MemStats
// // 	runtime.ReadMemStats(&m)
// // 	fmt.Printf("TotalAlloc (Heap) = %v GB\n", m.TotalAlloc/1024/1024/1024)
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // 	"math/big"
// // )

// // func main() {

// // 	intA := math.MaxInt64

// // 	intA += 1

// // 	bigA := big.NewInt(math.MaxInt64)

// // 	bigA.Add(bigA, big.NewInt(1))

// // 	fmt.Println("MaxInt64: ", math.MaxInt64)
// // 	fmt.Println("Int   :", intA)
// // 	fmt.Println("Big Int : ", bigA.String())
// // }

// // package main

// // import "fmt"

// // func main() {
// // 	comment1 := `This is the BEST
// // thing ever!`
// // 	comment2 := `This is the BEST\nthing ever!`
// // 	comment3 := "This is the BEST\nthing ever!"
// // 	fmt.Print(comment1, "\n\n")
// // 	fmt.Print(comment2, "\n\n")
// // 	fmt.Print(comment3, "\n")
// // }

// package main

// import "fmt"

// func main() {
// 	// username := "Sir_King_Über"

// 	// runes := []rune(username)

// 	// for i := 0; i < len(runes); i++ {
// 	// 	fmt.Print(string(runes[i]), " ")
// 	// }
// 	// fmt.Println()

// 	// for i := 0; i < len(username); i++ {
// 	// 	fmt.Print(string(username[i]), " ")
// 	// }
// 	// fmt.Println()

// 	// fmt.Println("WITH RANGE")

// 	// for _, v := range runes {
// 	// 	fmt.Print(string(v), " ")
// 	// }
// 	// fmt.Println()

// 	// myName := "Oussama Tailba"

// 	// fmt.Println(myName[5:7])

// 	// fmt.Println("Whaaat from NVIM")

// 	nums := [3]int{1, 2, 3}
// 	var nums2 [3]int = [3]int{1, 2, 3}

// 	fmt.Printf("%#v\n", nums)
// 	fmt.Printf("%#v\n", nums2)

// 	fmt.Println("nums == nums2", nums == nums2)

// 	arr := [7]int{1, 2, 3, 4, 5}

// 	fmt.Println(len(arr))

// }
