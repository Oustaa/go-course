// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // type safeCounter struct {
// // 	counts map[string]int
// // 	mu     *sync.Mutex
// // }

// // func (sc safeCounter) inc(key string) {
// // 	sc.mu.Lock()
// // 	defer sc.mu.Unlock()
// // 	sc.slowIncrement(key)
// // }

// // func (sc safeCounter) val(key string) int {
// // 	sc.mu.Lock()
// // 	defer sc.mu.Unlock()
// // 	return sc.slowVal(key)
// // }

// // // don't touch below this line

// // func (sc safeCounter) slowIncrement(key string) {
// // 	tempCounter := sc.counts[key]
// // 	time.Sleep(time.Microsecond)
// // 	tempCounter++
// // 	sc.counts[key] = tempCounter
// // }

// // func (sc safeCounter) slowVal(key string) int {
// // 	time.Sleep(time.Microsecond)
// // 	return sc.counts[key]
// // }

// func main() {
// 	ch := make(chan string)

// 	wg := sync.WaitGroup{}

// 	wg.Add(3)
// 	go test(ch, "Hi i love kaoutar", &wg)
// 	go test(ch, "And She loves me too", &wg)
// 	go test(ch, "I will find a better paying job, inshaalah", &wg)

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for message := range ch {
// 		fmt.Println(string(message))
// 	}

// }

// func test(ch chan string, message string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	ch <- message
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mu := &sync.Mutex{}

	go writeLoop(m, mu)
	go readLoop(m, mu)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.Mutex) {

	for {
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.Unlock()
	}
}
