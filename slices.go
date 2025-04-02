package main

import (
	F "fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	// F.Printf("%#v\n", os.Args[1:])
	if len(os.Args) < minArgs {
		F.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}
	var args []string
	args = append(args, os.Args[1:]...)
	// for i := 1; i < len(os.Args); i++ {
	// args = append(args, os.Args[i])
	// }
	return args
}

func findLongest(args []string) string {
	longest := args[0]
	for i := 1; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func fromArrToSlice(arr [5]int) []int {
	return arr[:]
}

func maiwn() {
	// if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
	// 	F.Println("The longest word passed was:", longest)
	// } else {

	// 	F.Println("There was an error")
	// 	os.Exit(1)
	// }

	arr := [3]int{1, 2, 3}
	slice := arr[:]

	arr[1] = 1_000_000
	slice = append(slice, 123, 123, 123, 123, 123, 123, 123, 123)
	slice = append(slice, 123_000)
	slice = append(slice, 123_000)
	arr[2] = 23

	F.Printf("%#v\n", arr)
	F.Printf("%#v\n", slice)
	F.Printf("%#v\n", cap(slice))

	makeSlice := make([]string, 12) // (type, len, capacity) => len for the slice, capacity for the underlying array
	F.Printf("%#v\n", makeSlice)

	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	// copied := copy(s2, s1)

	F.Printf("%p\n", s1)
	F.Printf("%p\n", s2)

}
