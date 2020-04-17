package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var arr [300000000]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	start := time.Now()

	a := make(chan int, 1000)
	group := len(arr) / 500000
	mod := len(arr) % 500000
	var count int
	if mod == 0 {
		count = group
	} else {
		count = group + 1
	}
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		end := 500000 * (i + 1)
		if end < len(arr) {
			go add(arr[end-500000:end], &wg, a)
		} else {
			go add(arr[end-500000:len(arr)], &wg, a)
		}
	}
	wg.Wait()
	close(a)
	var sum int
	for i := range a {
		sum += i
	}
	fmt.Println(sum)

	cost := time.Since(start)
	fmt.Println(cost)

	btime := time.Now()
	addsec(arr[:])
	bcost := time.Since(btime)
	fmt.Println(bcost)

}

func add(arr []int, wg *sync.WaitGroup, a chan int) {
	var sum int
	for _, i := range arr {
		sum += i
	}
	wg.Done()
	a <- sum
}

func addsec(arr []int) {
	var sum int
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
}
