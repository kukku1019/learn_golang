package main

import (
	"fmt"
	"runtime"
	"time"
)

func f(ch chan int) {
	msg := int{}
	msg <- ch
	msg = msg * 10
	fmt.Printf("channel:%v\n", msg)
}
func main() {
	go f(1)
	go f(2)
	go f(3)
	//現在のgoroutine の数を表示
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
	go func() {
		fmt.Println("終了")
		runtime.Goexit()
	}()

}
