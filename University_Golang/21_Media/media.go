/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input una sequenza di numeri reali strettamente positivi (un numero è
			strettamente positivo se è maggiore di 0; se un numero è minore o uguale 0 non è strettamente positivo). La lettura
			termina quando viene letto un numero minore o uguale a 0.

			Il programma deve stampare a video il risultato della media aritmetica dei valori inseriti.

	ESEMPIO :
		$ go run medie.go
		Inserisci una sequenza di numeri (interrompi con numero<=0): 4 6 8 0
		Media aritmetica: 6

		$ go run medie.go
		Inserisci una sequenza di numeri (interrompi con numero<=0): 3 5 2 6 1 -1
		Media aritmetica: 3.4
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	// dichiarazione delle due variabili locali necessarie che non si possono resettare con il ciclo
	var somma float64
	var contatore int

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ") // messaggio di output a livello di terminale

	for { // ciclo infinito
		var numero float64 // dichiarazione della variabile che si resetta ogni ciclo
		fmt.Scan(&numero)  // richiesta del dato a livello del terminale

		if numero <= 0 { // condizionale che controlla se inserisco la condizione per finire il ciclo
			break // interruzione del ciclo
		}

		somma += numero // somma progressiva dei numeri
		contatore++     // contatore di quanti numeri l'utente sta inserendo
	}

	// ✅ DOPO : da aggiungere prima di stampare la media
	if contatore == 0 {
		fmt.Println("Nessun numero valido inserito.")
		return
	}

	media := somma / float64(contatore) // calcolo della media

	fmt.Printf("Media aritmetica: %g\n", media) // stampa a terminale del valore della media
}
