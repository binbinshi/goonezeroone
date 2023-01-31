package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int){
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Println("put :", i)
	}
}

func Producer2(queue chan<- int){
	for i := 10; i < 20; i++ {
		queue <- i
		fmt.Println("2put :", i)
	}
}

func Consumer(queue <-chan int){
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("get:", v)
	}
}

func Consumer2(queue <-chan int){
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("2get:", v)
	}
}

func main() {
	queue := make(chan int, 100)
	go Producer(queue)
	go Consumer(queue)
	go Producer2(queue)
	go Consumer2(queue)
	time.Sleep(3 * time.Second)
}
