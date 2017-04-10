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

func variadicFanIn(channels ...<-chan string) <-chan string {
	broadcast := make(chan string)
	go func() {
		for {
			for i := range channels {
				c := channels[i]
				select {
				case s := <-c:
					broadcast <- s
				default:
				}
			}
		}
	}()
	return broadcast
}

func sliceFanIn(channels []<-chan string) <-chan string {
	return variadicFanIn(channels...)
}

func main() {
	// variadic slice
	names := []string{
		"Joe",
		"Ann",
		"Bob",
		"Liz",
	}
	channels := []<-chan string{}
	for _, s := range names {
		channels = append(channels, boring(s))
	}

	broadcast := variadicFanIn(channels...)
	for i := 0; i < 15; i++ {
		println(<-broadcast)
	}
	println("You're all boring; I'm leaving")
	println()

	// and proper slice
	channels = []<-chan string{}
	for _, s := range names {
		channels = append(channels, boring(s))
	}
	broadcast = sliceFanIn(channels)
	for i := 0; i < 15; i++ {
		println(<-broadcast)
	}
	println("You're all boring; I'm leaving")
}
