package main

import (
	"fmt"
	"time"
)

func disp(s string) {
	fmt.Println(s)
}

func main() {
	go disp("hello world")
	time.Sleep(1000)
	disp("hello world")
}
