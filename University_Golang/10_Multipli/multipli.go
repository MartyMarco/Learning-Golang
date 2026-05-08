/*

TESTO : Scrivere un programma che legge da standard input un numero intero n e verifica se il numero è multiplo di 10.

ESEMPIO :
$ go run multiplo10.go
Inserisci numero: 15
15 non è multiplo di 10

$ go run multiplo10.go
Inserisci numero: 20
20 è multiplo di 10

*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile in ingresso necessaria

	fmt.Print("Inserisci numero : ") // messaggio di output al terminale
	fmt.Scan(&numero)                // richiesta del dato in ingresso

	if numero%10 == 0 { // condizione se il resto della divisione per 10 è uguale a zero
		fmt.Printf("%d è multiplo di 10\n", numero) // stampa a terminale una risposta positiva
	} else {
		fmt.Printf("%d non è multiplo di 10\n", numero) // altrimenti stampa a terminale una risposta negativa
	}
}

/*
RICORDA :
	-	Il segno newline alla fine del fmt.Printf(), senza il cursore del terminale rimane sulla stessa riga dopo l'output, comportamento indesiderato
		su molti sistemi.
*/
