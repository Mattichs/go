package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	nome string
}

type Veicolo struct {
	tipo string
}

func noleggia(cliente Cliente, veicoli []Veicolo, wg *sync.WaitGroup, mutex *sync.Mutex, conteggio map[string]int) {
	defer wg.Done()

	// Genera un numero casuale per selezionare un veicolo
	rand.Seed(time.Now().UnixNano())
	indice := rand.Intn(len(veicoli))

	// Ottieni il veicolo selezionato
	veicolo := veicoli[indice]

	// Aggiorna il conteggio del veicolo noleggiato
	mutex.Lock()
	conteggio[veicolo.tipo]++
	mutex.Unlock()

	fmt.Printf("Il cliente %s ha noleggiato il veicolo %s\n", cliente.nome, veicolo.tipo)
}

func stampa(conteggio map[string]int) {
	fmt.Println("Numero di veicoli noleggiati:")
	for tipo, count := range conteggio {
		fmt.Printf("%s: %d\n", tipo, count)
	}
}

func main() {
	clienti := []Cliente{
		{nome: "Cliente 1"},
		{nome: "Cliente 2"},
		{nome: "Cliente 3"},
		{nome: "Cliente 4"},
		{nome: "Cliente 5"},
		{nome: "Cliente 6"},
		{nome: "Cliente 7"},
		{nome: "Cliente 8"},
		{nome: "Cliente 9"},
		{nome: "Cliente 10"},
	}

	veicoli := []Veicolo{
		{tipo: "Berlina"},
		{tipo: "SUV"},
		{tipo: "Station Wagon"},
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	conteggio := make(map[string]int)

	for _, cliente := range clienti {
		wg.Add(1)
		go noleggia(cliente, veicoli, &wg, &mutex, conteggio)
	}

	wg.Wait()

	stampa(conteggio)
}
