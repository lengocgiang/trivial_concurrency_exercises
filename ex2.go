package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var meals = []string{"chorizo", "chopitos", "pimientos de padrón", "croquetas", "atatas bravas"}
var friends = []string{"Alice", "Bop", "Charlie", "Dave"}

func main() {
	fmt.Println("Bon appétit!")
	var wg sync.WaitGroup
	for _, friend := range friends {
		enjoin(friend, &wg)
	}
	wg.Wait()
}

func enjoin(name string, wg *sync.WaitGroup) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	wg.Add(1)

	go func() {

		defer wg.Done()

		var wg2 sync.WaitGroup

		for _, val := range meals {
			nums := 5 + r1.Intn(5)
			for i := 0; i < nums; i++ {
				morse(name, val, &wg2)
			}
		}
		wg2.Wait()
	}()
}

func morse(name string, meal string, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("%s is enjoing some %s \n", name, meal)
	go func() {
		t := seed2()
		time.AfterFunc(t, func() {
			wg.Done()
		})
	}()
}

func seed2() time.Duration {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	t := 30 + r1.Intn(3*60-30)
	return time.Duration(t) * time.Millisecond
}
