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

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		s := sprintf("%s %d", msg, i)
		c <- s
		ms := time.Duration(rand.Intn(1e3))
		time.Sleep(ms * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		s := <-c
		printf("You say: %q\n", s)
	}
	println("You're boring; I'm leaving")
}
