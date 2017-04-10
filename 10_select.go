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
			s := sprintf("%s %d", msg, i)
			c <- s
			ms := time.Duration(rand.Intn(1e3))
			time.Sleep(ms * time.Millisecond)
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan string) <-chan string {
	broadcast := make(chan string)
	go func() {
		for {
			select {
			case s := <-c1:
				broadcast <- s
			case s := <-c2:
				broadcast <- s
			}
		}
	}()
	return broadcast
}

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	broadcast := fanIn(joe, ann)
	for i := 0; i < 15; i++ {
		println(<-broadcast)
	}
	println("You're both boring; I'm leaving")
}
