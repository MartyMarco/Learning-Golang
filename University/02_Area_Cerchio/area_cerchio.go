/*

TESTO : Scrivere un programma che legga da standard input il raggio di un cerchio e ne calcoli circonferenza e area.

SUGGERIMENTO : l'area del cerchio si calcola facendo raggio x raggio x pi_greco, mentre la circonferenza facendo 2 x raggio x pi_greco.
Il valore numerico di pi_greco è memorizzato nella costante Pi del package math, a cui ci si può riferire con math.Pi.

ESEMPIO :
$ go run cerchio.go
Raggio = 2.5
Circonferenza = 15.707963267948966
Area = 19.634954084936208

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var raggio float64 // Dichiarazione della variabile raggio

	fmt.Print("Inserisci il raggio del cerchio : ") // messaggio in output
	fmt.Scan(&raggio)                               // richiesta dei dato in input

	circonferenza := 2 * raggio * math.Pi // calcolo della circonferenza del cerchio
	area := raggio * raggio * math.Pi     // calcolo dell'area del cerchio

	fmt.Println("Circonferenza : ", circonferenza) // Restituzione del valore della circonferenza al terminale
	fmt.Println("Area : ", area)                   // Restituzione del valore del'area al terminale
}
