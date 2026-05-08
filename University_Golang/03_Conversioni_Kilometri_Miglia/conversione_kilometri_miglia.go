/*

TESTO : Scrivere un programma che legga da standard input una distanza in Km ed effettui la conversione di tale distanza in miglia (1 Km = 0.62137 mi).

ESEMPIO :
$ go run convertitore.go
Distanza (Km) = 12
Distanza (mi) = 7.45644

*/

package main

import (
	"fmt"
)

const conversione float64 = 0.62137 // Dichiarazione + inizializzazione della costante conversione kilometri-miglia

func main() {
	var kilometri float64 // Dichiarazione variabile kilometri

	fmt.Print("Inserisci il numero di kilometri da convertire : ") // messagio in output al terminale
	fmt.Scan(&kilometri)                                           // richiesta del dato kilometri

	miglia := kilometri * conversione // calcolo della conversione da kilometri a miglia

	fmt.Println("Miglia : ", miglia) // restituzione del valore delle miglia a livello del terminale
}
