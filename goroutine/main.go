package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func cf(ch chan int) {
	msg := <-ch
	fmt.Printf("channel内の値:%v\n", msg)
	close(ch)

}

func s(num int) {
	fmt.Printf("go print:%d\n", num)
}

func goroutine(ch chan string) {
	for s := range ch {
		log.Println("Hello", s)
		time.Sleep(time.Second)
	}
}

func main() {

	go s(1)
	//capacity　10
	ch0 := make(chan int, 10)
	ch0 <- 111
	go cf(ch0)

	//現在のgoroutine の数を表示
	fmt.Printf("現在のgoroutine数:%d\n", runtime.NumGoroutine())

	go func() {
		fmt.Println("終了")
		runtime.Goexit()
	}()
	time.Sleep(time.Second)

	ch1 := make(chan string, 5)

	go goroutine(ch1)

	ch1 <- "World0"
	ch1 <- "World1"
	ch1 <- "World2"
	ch1 <- "World3"
	ch1 <- "World4"
	ch1 <- "World5"
	ch1 <- "World6"
	ch1 <- "World7"

	fmt.Println("end")
	time.Sleep(time.Second * 6)
}
