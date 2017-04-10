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
	rand.Seed(time.Now().UTC().UnixNano())
}

//////////////////////////////////////////////
/// functions
//////////////////////////////////////////////

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- sprintf("%s %d", msg, i):
				ms := time.Duration(rand.Intn(1e3))
				time.Sleep(ms * time.Millisecond)
			}
		}
	}()
	return c
}

func main() {
	joe := boring("Joe")
	for i := 0; i < 10; i++ {
		println(<-joe)
	}
}
