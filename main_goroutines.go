package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

func rev(fn func(int) int, x int, ch chan<- int) {
	ch <- fn(x)
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		wg := new(sync.WaitGroup)
		array := make([]int, n)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			wg.Add(1)
			go func(in1 <-chan int, in2 <-chan int, wg *sync.WaitGroup, i int) {
				defer wg.Done()
				ch1 := make(chan int)
				ch2 := make(chan int)

				defer close(ch1)
				defer close(ch2)

				go rev(fn, x1, ch1)
				go rev(fn, x2, ch2)

				array[i] = <-ch1 + <-ch2
			}(in1, in2, wg, i)
		}
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- array[i]
		}
	}()
}
