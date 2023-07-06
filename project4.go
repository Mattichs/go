package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CurrencyPair struct {
	Name  string
	Price float64
}

func simulateMarketData(currencyPairs chan<- CurrencyPair, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		eurUsdPrice := 1.0 + rand.Float64()*0.5
		gbpUsdPrice := 1.0 + rand.Float64()*0.5
		jpyUsdPrice := 0.006 + rand.Float64()*0.003

		currencyPairs <- CurrencyPair{Name: "EUR/USD", Price: eurUsdPrice}
		currencyPairs <- CurrencyPair{Name: "GBP/USD", Price: gbpUsdPrice}
		currencyPairs <- CurrencyPair{Name: "JPY/USD", Price: jpyUsdPrice}

		time.Sleep(1 * time.Second)
	}
}

func selectPair(currencyPair CurrencyPair, wg *sync.WaitGroup) {
	defer wg.Done()

	switch currencyPair.Name {
	case "EUR/USD":
		if currencyPair.Price > 1.20 {
			fmt.Println("Vendita di EUR/USD in corso...")
			time.Sleep(4 * time.Second)
			fmt.Println("Vendita di EUR/USD confermata")
		}
	case "GBP/USD":
		if currencyPair.Price < 1.35 {
			fmt.Println("Acquisto di GBP/USD in corso...")
			time.Sleep(3 * time.Second)
			fmt.Println("Acquisto di GBP/USD confermato")
		}
	case "JPY/USD":
		if currencyPair.Price < 0.0085 {
			fmt.Println("Acquisto di JPY/USD in corso...")
			time.Sleep(3 * time.Second)
			fmt.Println("Acquisto di JPY/USD confermato")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	currencyPairs := make(chan CurrencyPair)
	var wg sync.WaitGroup

	wg.Add(1)
	go simulateMarketData(currencyPairs, &wg)

	go func() {
		wg.Wait()
		close(currencyPairs)
	}()

	for currencyPair := range currencyPairs {
		wg.Add(1)
		go selectPair(currencyPair, &wg)
	}

	time.Sleep(5 * time.Second)
}
