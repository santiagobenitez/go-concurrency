package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Person struct {
	name     string
	lastName string
	age      int
}

type readPeople struct {
	resp chan []Person
}

var people []Person
var reads chan *readPeople
var cacheReads int64
var dbReads int64
var ops int64

func main() {

	reads = make(chan *readPeople)

	go getPeople()

	//100 concurrent go routines getting people data
	for p := 0; p < 100; p++ {
		go func() {
			for {
				readP := &readPeople{
					resp: make(chan []Person)}
				reads <- readP
				atomic.AddInt64(&ops, 1)
				<-readP.resp
			}
		}()
	}

	//let the go routines read for a while
	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	dbReadsFinal := atomic.LoadInt64(&dbReads)
	cacheReadsFinal := atomic.LoadInt64(&cacheReads)
	fmt.Println("ops done: ", opsFinal)
	fmt.Println("db reads: ", dbReadsFinal)
	fmt.Println("cacheReads: ", cacheReadsFinal)
}

func getPeople() {
	for {
		read := <-reads
		if len(people) == 0 {
			//simulate reading from data base
			time.Sleep(time.Millisecond * 20)
			people = []Person{
				Person{
					name:     "John",
					lastName: "Snow",
					age:      28},
				Person{
					name:     "Cercie",
					lastName: "Lannister",
					age:      35}}
			atomic.AddInt64(&dbReads, 1)
			read.resp <- people
		} else {
			//data is cached atomic.AddInt64(&cacheReads, 1)
			read.resp <- people
		}
	}
}
