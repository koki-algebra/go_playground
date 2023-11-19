package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"Golang", "Ruby", "Java", "Python"}

	for _, word := range words {
		fmt.Println(word)
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("Time up!")
			continue
		case ans := <-input():
			if ans == word {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect...")
			}
		}
	}
}

func input() <-chan string {
	ch := make(chan string)
	go func() {
		var ans string
		fmt.Print("> ")
		fmt.Scanln(&ans)
		ch <- ans
	}()

	return ch
}
