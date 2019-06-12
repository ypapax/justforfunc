package main

import (
	"log"
	"math/rand"
	"time"
)

// https://youtu.be/t9bEg2A4jsw?t=54

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		log.Println(v)
	}
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(1000)))
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	var result = make(chan int)
	go func() {
		var aClosed, bClosed bool
		defer close(result)
		for !aClosed || !bClosed {
			select {
			case v, ok := <-a:
				if !ok {
					aClosed = true
					continue
				}
				result <- v
			case v, ok := <-b:
				if !ok {
					bClosed = true
					continue
				}
				result <- v
			}
		}
	}()
	return result
}
