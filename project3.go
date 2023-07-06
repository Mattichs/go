package main

import (
	"fmt"
	"sync"
	"time"
)

const numCakes = 5

var (
	cakeCounter   int
	cookedSpace   = make(chan struct{}, 2)
	iceredSpace   = make(chan struct{}, 2)
	cooks         = make(chan struct{}, 1)
	decorators    = make(chan struct{}, 1)
	icers         = make(chan struct{}, 1)
	cookTime      = 1 * time.Second
	icerTime      = 4 * time.Second
	decoratorTime = 8 * time.Second
)

func main() {
	fmt.Println("Inizio a produrre le torte")
	var wg sync.WaitGroup

	for i := 0; i < numCakes; i++ {
		cakeID := getNextCakeID()
		wg.Add(1)
		go cookCake(cakeID, &wg)
	}

	wg.Wait()
	fmt.Println("Tutte le torte prodotte")
}

func cookCake(cakeID int, wg *sync.WaitGroup) {
	cooks <- struct{}{}

	fmt.Printf("Cucino le torta %d\n", cakeID)

	time.Sleep(cookTime)

	fmt.Printf("Finito di Cuocere la torta %d\n", cakeID)

	cookedSpace <- struct{}{}
	wg.Add(1)
	go iceCake(cakeID, wg)

	<-cooks
	defer wg.Done()
}

func iceCake(cakeID int, wg *sync.WaitGroup) {
	icers <- struct{}{}
	fmt.Printf("Glasso la torta %d\n", cakeID)

	time.Sleep(icerTime)

	fmt.Printf("Finito glassatura torta %d\n", cakeID)

	iceredSpace <- struct{}{}
	wg.Add(1)
	go decorateCake(cakeID, wg)
	<-cookedSpace
	<-icers
	defer wg.Done()
}

func decorateCake(cakeID int, wg *sync.WaitGroup) {
	decorators <- struct{}{}
	fmt.Printf("Decoro la torta %d\n", cakeID)

	time.Sleep(decoratorTime)

	fmt.Printf("Finito decorazione torta %d\n", cakeID)

	<-iceredSpace
	<-decorators
	defer wg.Done()
}

func getNextCakeID() int {
	cakeCounter++
	return cakeCounter
}
