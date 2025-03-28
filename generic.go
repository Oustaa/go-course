// package main

// func Maximum[Num int | float32 | float64](nums []Num) Num {
// 	if len(nums) == 0 {
// 		return -1
// 	}

// 	max := nums[0]

// 	for _, num := range nums {
// 		if num > max {
// 			max = num
// 		}
// 	}

// 	return max
// }

// func min[T int | float32 | float64](nums []T) T {
// 	if len(nums) == 0 {
// 		return -1
// 	}

// 	min := nums[0]

// 	for _, num := range nums {
// 		if num < min {
// 			min = num
// 		}
// 	}

// 	return min
// }

package main

import (
	"fmt"
)

func FindLargestRanchStock[K comparable, V int | float64](m map[K]V) K {
	var stock V
	var name K
	for k, v := range m {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name
}

func main() {
	animalStock := map[string]int{
		"Chicken": 5,
		"Cattle":  20,
		"Horses":  4,
	}
	miscStock := map[string]float64{
		"Hay":        5.5,
		"Feed":       1.2,
		"Fertilizer": 4.5,
	}
	bookStock := map[int]float64{
		1: 5.5,
		2: 1.2,
		3: 422.5,
	}
	largestStockOnRanchInt := FindLargestRanchStock(animalStock)
	fmt.Printf("The largest stocked item on the ranch is %s\n",
		largestStockOnRanchInt)
	largestStockOnRanchFloat := FindLargestRanchStock(miscStock)
	fmt.Printf("The largest stocked item on the ranch is %s\n",
		largestStockOnRanchFloat)
	largestStockOnRanchKInt := FindLargestRanchStock(bookStock)
	fmt.Printf("The largest stocked item on the ranch is %d\n",
		largestStockOnRanchKInt)
}
