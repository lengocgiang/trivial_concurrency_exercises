package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	getReady("Bob", &wg)
	getReady("Alice", &wg)
	wg.Wait()

	var wg2 sync.WaitGroup
	setArming(&wg2)
	takeShoes("Bob", &wg)
	takeShoes("Alice", &wg)

	waitShoes(&wg)

	waitArming(&wg2)

}

func getReady(name string, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("%s started getting ready\n", name)
	t := Seed(60, 90)
	time.AfterFunc(t, func() {
		defer wg.Done()
		fmt.Printf("%s spent %v getting ready\n", name, t)
	})
}

func takeShoes(name string, wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Printf("%s started putting on shoes\n", name)
	t := Seed(35, 45)
	time.AfterFunc(t, func() {
		defer wg.Done()
		fmt.Printf("%s spent %v putting on shoes\n", name, t)
	})
}

func waitShoes(wg *sync.WaitGroup) {
	wg.Wait()
	fmt.Println("exit door")
}

func setArming(wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("Arming alarm.")
	time.AfterFunc(600*time.Millisecond, func() {
		defer wg.Done()
		fmt.Println("Alarm is armed")
	})
}

func waitArming(wg *sync.WaitGroup) {
	wg.Wait()
}

func Seed(a, b int) time.Duration {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r := a + r1.Intn(b-a)

	return time.Duration(r) * time.Millisecond
}
