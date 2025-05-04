// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func myAdd(from int, to int, wg *sync.WaitGroup, result *int) {
// 	for i := from; i <= to; i++ {
// 		*result += i
// 	}

// 	wg.Done()
// }

// func main() {
// 	s1 := 0

// 	wg := sync.WaitGroup{}

// 	wg.Add(2)
// 	go myAdd(1, 25, &wg, &s1)
// 	go myAdd(26, 50, &wg, &s1)

// 	wg.Wait()
// 	fmt.Println(s1)
// }

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func next(num *int32, wg *sync.WaitGroup) {
	// *num += 1
	atomic.AddInt32(num, 1)
	wg.Done()
}

func maisdn() {
	counter := int32(0)
	wg := sync.WaitGroup{}

	wg.Add(3)
	go next(&counter, &wg)
	go next(&counter, &wg)
	go next(&counter, &wg)

	wg.Wait()
	fmt.Println(counter)

}
