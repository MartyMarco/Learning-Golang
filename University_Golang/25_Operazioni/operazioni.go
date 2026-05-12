/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO : Create un programma che legge da standard input due numeri interi, che chiameremo x e y. Letti i due numeri, il programma stampa:
				-	il maggiore tra x e y
				-	il minore tra x e y
				-	il risultato della somma tra x e y
				-	il risultato della differenza tra il maggiore e il minore
				-	il risultato della divisione x/y
				-	il risultato del prodotto x*y
				-	il valore medio tra x e y
				-	il risultato di x elevato alla y (utilizzando sia un ciclo for sia la funzione math.Pow)

	ESEMPIO :
		$ go run operazioni.go
		Inserisci due numeri interi:
		2 4
		Maggiore: 4
		Minore: 2
		Somma: 6
		Differenza: 2
		Prodotto: 8
		Divisione: 0.5
		Valore medio: 3
		Potenza (ciclo for): 16
		Potenza (math.Pow): 16
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var numeroX, numeroY, differenza int       // dichiarazione delle due variabili necessarie per l'inseriemento
	fmt.Print("Inserisci due numeri interi: ") // messaggio di output a livello di terminale
	fmt.Scan(&numeroX, &numeroY)               // richiesta di dati in input

	if numeroX > numeroY {
		fmt.Printf("Maggiore: %d\n", numeroX) // il primo numero è maggiore del secondo
		fmt.Printf("Minore: %d\n", numeroY)   // il secondo numero è minore del primo
		differenza = numeroX - numeroY        // calcolo della differenza nel caso il primo sia maggiore
	} else {
		fmt.Printf("Maggiore: %d\n", numeroY) // il secondo numero è maggiore del primo
		fmt.Printf("Minore: %d\n", numeroX)   // il primo numero è minore del secondo
		differenza = numeroY - numeroX        // calcolo della differenza nel caso il secondo sia maggiore
	}

	somma := numeroX + numeroY                 // calcolo della somma
	fmt.Printf("Somma: %d\n", somma)           // stampa della somma a livello di terminale
	fmt.Printf("Differenza: %d\n", differenza) // stampa della differenza a livello di terminale

	prodotto := numeroX * numeroY          // calcolo del prodotto
	fmt.Printf("Prodotto: %d\n", prodotto) // stampa del prodotto a livello di terminale

	/*
		divisione := float64(numeroX) / float64(numeroY) // calcolo della divisione (numero con la virgola)
		fmt.Printf("Divisione: %g\n", divisione)         // stampa della divisione a terminale togliendo gli zeri extra in fondo con (%g)
	*/

	/*
		✅ DOPO : Da aggiungere prima della divisione
		Evitare la divisione per zero
		Perché: dividere per zero manda in crash il programma (o produce +Inf). È un caso limite classico da sempre controllare.
	*/
	if numeroY == 0 {
		fmt.Println("Divisione: impossibile (divisione per zero)")
	} else {
		divisione := float64(numeroX) / float64(numeroY)
		fmt.Printf("Divisione: %g\n", divisione)
	}

	media := somma / 2                      // calcolo della media tramite divisione intera
	fmt.Printf("Valore medio: %d\n", media) // stampa a terminale della media

	/*
		❌ PRIMA : non prende in considerazione il caso y=0
			potenza := numeroX             // dichiarazione e inizializzazione della variabile potenza con il valore della base della potenza
			for i := 1; i < numeroY; i++ { // parto da 1 perchè è il valore inizializzato
				potenza = potenza * numeroX // moltiplico a ogni ciclo il valore del ciclo precedente per il numero della base
			}
	*/

	// ✅ DOPO :
	potenza := 1
	for i := 0; i < numeroY; i++ { // faccio y cicli
		potenza *= numeroX // Totale moltiplicazioni: y
	}
	fmt.Printf("Potenza (ciclo for): %d\n", potenza)

	// Entrambi producono lo stesso risultato per y > 0. Il punto che sollevavo io riguarda però il caso y = 0.
	// Quindi il tuo ragionamento sul ciclo è corretto — l'unico problema reale è quel caso limite.
	// Se il tuo esercizio non prevede y = 0, il tuo codice funziona perfettamente.

	potenzaDopo := math.Pow(float64(numeroX), float64(numeroY)) // calcolo della potenza anche attraverso il package "math"
	fmt.Printf("Potenza (math.Pow): %g\n", potenzaDopo)         // stampa a terminale anche della possibilità tramite calcolo del package "math"
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	N.B :	-	Separare calcoli e stampe (leggibilità).
				Perché: separare la logica dalla presentazione rende il codice più facile da leggere, testare e modificare in futuro.

			-	Dichiarazione delle variabili più vicina all'uso.
				Perché: in Go è idiomatico dichiarare le variabili con := il più vicino possibile al loro primo utilizzo, riducendo il rischio
				di usarle non inizializzate.
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
