/*

TESTO : Scrivere un programma che legge da standard input un intero n e stampa a video se il numero è pari o dispari.

ESEMPIO :
$ go run paridispari.go
Inserisci un numero: 10
10 è pari

$ go run paridispari.go
Inserisci un numero: 11
11 è dispari

*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile necessaria in input

	fmt.Print("Inserisci un numero : ") // messaggio in output al terminale
	fmt.Scan(&numero)                   // richiesta del dato in ingresso a terminale

	if numero%2 == 0 { // condizione che la divisione per due non produce resto e quindi il numero è pari
		fmt.Printf("%d è pari\n", numero)
	} else { // altrimenti il numero per forza è dispari
		fmt.Printf("%d è dispari\n", numero)
	}
}
