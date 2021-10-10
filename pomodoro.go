package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		pomoSesh()
		pomoBreak()
	}
}

func pomoBreak() {
	fmt.Println("Break begins.")
	time.Sleep(time.Minute * 5)
	fmt.Println("Break's up!")
}

func pomoSesh() {
	fmt.Println("Sesh starts.")
	time.Sleep(time.Minute * 25)
	fmt.Println("Sesh ends.")
}
