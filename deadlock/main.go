package main

import "fmt"

// func main() {
// 	var wg sync.WaitGroup
// 	var mu1, mu2 sync.Mutex

// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		mu1.Lock()
// 		fmt.Println("goroutine 1 acquired lock 1")
// 		mu2.Lock()
// 		fmt.Println("goroutine 1 acquired lock 2")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		mu2.Lock()
// 		fmt.Println("goroutine 2 acquired lock 2")
// 		mu1.Lock()
// 		fmt.Println("goroutine 2 acquired lock 1")
// 	}()
// 	wg.Wait()
// }

func main() {
	ch := make(chan int)

	ch <- 42

	go func() {
		value := <-ch
		fmt.Println("Received:", value)
	}()

	fmt.Println("Sent: 42")
}
