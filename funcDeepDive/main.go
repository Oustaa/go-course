package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{2, 3, 4, 5}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformFunc1 := getTransformerFunc(&numbers)
	transformFunc2 := getTransformerFunc(&numbers2)

	fmt.Println(transformNumber(&numbers, transformFunc1))
	fmt.Println(transformNumber(&numbers, transformFunc2))

	multiplyByTwo := createTransformer(2)

	twoTwo := multiplyByTwo(34)

	fmt.Println(twoTwo)

}

func transformNumber(numbers *[]int, transform func(int) int) []int {
	dnumbers := []int{}

	for _, val := range *numbers {
		dnumbers = append(dnumbers, transform(val))
	}

	return dnumbers
}

func getTransformerFunc(numbers *[]int) func (int) int {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func testReturnFunc() func(int) int {
	return func(num int) int {
		return num * 2
	}
}


func createTransformer(factor int) func (int) int {
	return func (number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number - 1)
}