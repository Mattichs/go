package main

import (
	"fmt"
	"sync"
)

func countCharacter(str string, char byte, wg *sync.WaitGroup, countChan chan int) {
	defer wg.Done()
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++
		}
	}
	countChan <- count
}

func main() {
	str := "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
	char := byte('f')

	var wg sync.WaitGroup
	countChan := make(chan int)

	for i := 0; i < len(str); i++ {
		wg.Add(1)
		go countCharacter(string(str[i]), char, &wg, countChan)
	}

	go func() {
		wg.Wait()
		close(countChan)
	}()

	totalCount := 0
	for count := range countChan {
		totalCount += count
	}

	fmt.Printf("Il conteggio finale del carattere '%c' è %d\n", char, totalCount)
}
