package main

import (
	"fmt"
)

func main() {
	strs := make(chan string)
	testStrs := []string{"a", "ab", "abc", "abcd"}
	done := make(chan bool, len(testStrs))

	//发送的goroutine
	go func() {
		for _, item := range testStrs {
			strs <- item
			//fmt.Println("send msg item=" + item)
		}
		close(strs)
	}()
	//接收的goroutine
	go func() {
		for item := range strs {
			fmt.Println(item)
			done <- true
		}
	}()

	for true {

	}
}
