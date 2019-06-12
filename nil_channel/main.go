package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) chan int {
	var result = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range a {
			result <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range b {
			result <- v
		}
	}()
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
