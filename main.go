package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const delay time.Duration = 10 * time.Millisecond

func main() {
	quit := make(chan bool)
	letter := make(chan string)

	fmt.Println("Press enter...")
	go letterLoop(quit, letter)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

	quit <- true
	fmt.Println("Oh noes! You stopped on " + <-letter)
}

func letterLoop(quit <-chan bool, letter chan<- string) {
	i := 0
	for {
		select {
		case <-quit:
			letter <- intToChar(i % 26)
			return
		default:
			i++
		}

		time.Sleep(delay)
	}
}

func intToChar(num int) string {
	return string('A' + num)
}
