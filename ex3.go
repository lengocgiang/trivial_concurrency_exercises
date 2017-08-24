package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_COMPUTERS = 8

func main() {

	a := generateArray()

	fmt.Println(a)

	enterchan := make(chan int, 1)
	leavechan := make(chan int, 1)
	stopchan := make(chan bool)

	go managerComputer(enterchan, leavechan, stopchan, a)

	for i := range a {
		enterchan <- a[i]
	}

	<-stopchan

	fmt.Println("The place is empty, let's close up and go to the beach!")

}

func managerComputer(enterchan, leavechan chan int, stopchan chan bool, a []int) {
	nbUser := 0
	var queue []int

	for {
		select {
		case e := <-enterchan:
			if nbUser >= MAX_COMPUTERS {
				fmt.Printf("Tourist %d waiting for turn. \n", e)
				queue = append(queue, e)
			} else {
				nbUser++
				go using(e, leavechan)
			}

		case <-leavechan:
			nbUser--
			if nbUser == 0 {
				stopchan <- true
				break
			}
			if len(queue) > 0 {
				nextUser := queue[0]
				queue = queue[1:]
				go func() {
					enterchan <- nextUser
				}()
			}
		}
	}
}

func using(id int, leavechan chan int) {
	fmt.Printf("Tourist %d is online \n", id)
	usingtime := seedTime()
	time.AfterFunc(usingtime, func() {
		fmt.Printf("Tourist %d is done, having spend %v online \n", id, usingtime)
		leavechan <- id
	})
}

func seedNum() *rand.Rand {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1
}
func seedTime() time.Duration {
	// time 15 minutes - 2 hour
	// using milisecond to shorten simulator time
	r1 := seedNum()
	d := 15*60 + r1.Intn(2*60*60)
	return time.Duration(d) * time.Millisecond
}

func generateArray() []int {
	var a []int
	for i := 0; i < 25; i++ {
		a = append(a, i+1)
	}
	r := seedNum()

	for i := range a {
		j := r.Intn(24)
		a[i], a[j] = a[j], a[i]
	}
	return a
}
