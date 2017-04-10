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
/// init
//////////////////////////////////////////////

func init() {
	s := time.Now().UTC().UnixNano()
	rand.Seed(s)
}

//////////////////////////////////////////////
/// functions
//////////////////////////////////////////////

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- sprintf("%s %d", msg, i):
				ms := time.Duration(rand.Intn(1e3))
				time.Sleep(ms * time.Millisecond)
			case s := <-quit:
				println(s)
				quit <- sprintf("%s %s", msg, "See you!")
				return
			}
		}
	}()
	return c
}

func main() {
	quit := make(chan string)
	joe := boring("Joe", quit)
	for i := 0; i < 10; i++ {
		println(<-joe)
	}
	quit <- "Bye!"
	println(<-quit)
}
