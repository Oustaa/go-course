package main

import "fmt"

type Person2 struct {
	name string
	age  int8
}

func printPeople(people map[string]Person2) {
	for key, value := range people {
		fmt.Printf("%s: named %s and is %d years old\n", key, value.name, value.age)
	}
}

func mainw() {

	// myMap := map[string]float64{}

	first_map := map[string]Person2{
		"DEV-206-1":   {"Kaoutar Takiww", 22},
		"DEV-206wers": {"Oussama Tailba", 58},
		"DEV-206-11":  {"Oussama Tailba", 58},
	}

	delete(first_map, "DEV-206-11")
	delete(first_map, "DEV-206-1")

	first_map["DEV-206"] = Person2{"Oussama Tailba", 58}

	fmt.Printf("%#v\n", first_map)
	fmt.Printf("%#v\n", first_map["dfdf"])

	value, exists := first_map["DEV-206-1"]

	fmt.Printf("%#v\n", value.age)
	fmt.Println(exists)

	printPeople(first_map)
}
