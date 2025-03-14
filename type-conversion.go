package main

import (
	"fmt"
	"math"
)

func main1() {
	var num int64 = math.MaxInt64
	var smallerNum int8 = int8(num)

	name := []int{88, 52, 66, 965}

	// fmt.Println(string(name))
	fmt.Println(name)

	fmt.Println("num ", num)
	fmt.Println("smallerNum ", smallerNum)
}
