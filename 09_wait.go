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

type Message struct {
	str  string
	wait chan bool
}

func boringWait(msg string) <-chan Message {
	c := make(chan Message)
	go func() {
		waitForIt := make(chan bool)
		for i := 0; ; i++ {
			s := sprintf("%s %d", msg, i)
			m := Message{s, waitForIt}
			c <- m
			ms := time.Duration(rand.Intn(1e3))
			time.Sleep(ms * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

func sliceFanIn(channels []<-chan Message) <-chan Message {
	broadcast := make(chan Message)
	for i := range channels {
		// copy rather than use loop variable
		c := channels[i]
		go func() {
			for {
				broadcast <- <-c
			}
		}()
	}
	return broadcast
}

func main() {
	// prep
	names := []string{
		"Joe",
		"Ann",
		"Bob",
		"Liz",
	}
	channels := []<-chan Message{}
	for _, s := range names {
		channels = append(channels, boringWait(s))
	}
	broadcast := sliceFanIn(channels)

	// execution
	for i := 0; i < 15; i++ {
		msg := <-broadcast
		println(msg.str)
		msg.wait <- true
	}
	println("You're all boring; I'm leaving")
}
