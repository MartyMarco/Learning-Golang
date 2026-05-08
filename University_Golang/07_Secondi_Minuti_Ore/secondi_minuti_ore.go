/*

TESTO : Scrivere un programma che prenda in ingresso un valore intero rappresentante una quantità di tempo in secondi e restituisca il numero di ore, minuti e secondi
		corrspondenti.

ESEMPIO :
$ go run convertisec.go
Secondi? 15634
h:m:sec - 4:20:34

$ go run convertisec.go
Secondi? 134
h:m:sec - 0:2:14

*/

package main

import (
	"fmt"
)

func main() {
	var secondi int // dichiarazione della variabile necessaria

	fmt.Print("Inserisci il numero di secondi : ") // messaggio output a terminale
	fmt.Scan(&secondi)                             // richiesta del valore del dato al terminale

	/*
		GOOD :
		ore := secondi / 3600               // calcolo delle ore (60*60)
		restanti := secondi - (ore * 3600)  // calcolo dei secondi restanti tolte le ore
		minuti := restanti / 60             // calcolo dei minuti
		restanti = restanti - (minuti * 60) // calcolo dei secondi rimantenti tolti i minuti

	*/

	// Utilizziamo l'operatore del modulo "%" invece delle sottrazioni manuali, è più idiomatico e leggibile e elimina le variabili intermedie inutili.
	// BETTER :
	ore := secondi / 3600
	minuti := (secondi % 3600) / 60
	sec := secondi % 60

	fmt.Printf("h:m:sec - %d:%d:%d \n", ore, minuti, sec) // stampa a livello del terminale dei dati
}
