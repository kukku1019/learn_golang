package main

import (
	"fmt"
	"time"
)

func test(n int) int {
	var num int = 10
	n = n * num
	return n
}

func main() {
	fmt.Println("Hello Wold,你好 世界！")
	s := test(2)
	fmt.Printf("funcの結果:%d", s)
	time.Sleep(time.Second)
}
