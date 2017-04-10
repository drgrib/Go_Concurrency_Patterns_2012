package main

import (
	"fmt"
	"math/rand"
	"time"
)

//////////////////////////////////////////////
/// aliases
//////////////////////////////////////////////

var println = fmt.Println
var sprintf = fmt.Sprintf
var printf = fmt.Printf

//////////////////////////////////////////////
/// functions
//////////////////////////////////////////////

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			s := sprintf("%s %d", msg, i)
			c <- s
			ms := time.Duration(rand.Intn(1e3))
			time.Sleep(ms * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		s := <-c
		printf("You say: %q\n", s)
	}
	println("You're boring; I'm leaving")
}
