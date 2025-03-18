package main

import "fmt"

type job struct {
	title  string
	salary float32
}

type person struct {
	name   string
	age    int8
	maried bool
	job
}

func salarySum(people []person) float32 {
	var sum float32 = 0
	for _, person := range people {
		sum += person.salary
	}
	return sum
}

func main() {
	me := person{"Oussama tailba", 26, true, job{"Front end dev", 4000}}
	kaoutar := person{"Kaouta Taki", 22, true, job{"Front end dev", 3500}}

	fmt.Println(me)

	// Slice
	people := []person{kaoutar, me}
	fmt.Printf("%#v\n", people)

	// Family array
	meFamily := []person{kaoutar, me}
	fmt.Printf("%#v\n", meFamily)

	me.salary = 9000

	var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 80, 9}
	nums = append(nums[:7], nums[8:]...)

	fmt.Printf("%#v\n", nums)

	fmt.Printf("Our total salary is %.2f\n", salarySum(meFamily))

}
