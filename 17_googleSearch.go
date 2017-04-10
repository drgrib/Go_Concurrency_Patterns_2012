package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
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
/// search
//////////////////////////////////////////////

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		msg := sprintf("%s result for %q\n", kind, query)
		return Result(msg)
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

//////////////////////////////////////////////
/// Google
//////////////////////////////////////////////

type googleFunc func(query string) (results []Result)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

func Google1(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func Google2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func Google2_1(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)

func Google3(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// more replicas!
func Google4(query string) (results []Result) {
	c := make(chan Result)
	webs, images, videos := []Search{}, []Search{}, []Search{}
	for i := 0; i < 10; i++ {
		webs = append(webs, fakeSearch(sprintf("web%v", i)))
		images = append(images, fakeSearch(sprintf("image%v", i)))
		videos = append(videos, fakeSearch(sprintf("videos%v", i)))
	}
	go func() { c <- First(query, webs...) }()
	go func() { c <- First(query, images...) }()
	go func() { c <- First(query, videos...) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

//////////////////////////////////////////////
/// runGoogle
//////////////////////////////////////////////

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func runGoogle(Google googleFunc) {
	fname := GetFunctionName(Google)
	printf("\n== %v ==\n", fname)
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	println(results)
	println(elapsed)
}

func main() {
	googles := []googleFunc{
		Google1,
		Google2,
		Google2_1,
		Google3,
		Google4,
	}
	for _, g := range googles {
		runGoogle(g)
	}
}
