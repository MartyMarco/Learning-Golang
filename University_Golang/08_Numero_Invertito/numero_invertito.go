/*

TESTO : Scrivi un programma che inverta le cifre di un numero intero a tre cifre inserito da standard input. Esempio: dato n=123 il programma dovrebbe restituire 321.

SUGGERIMENTO : Puoi ottenere le singole cifre di un numero con operazioni di modulo e divisione intera.

ESEMPIO :
$ go run invertinum.go
Numero? 123
Risultato: 321

$ go run invertinum.go
Numero? 453
Risultato: 354

*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile necessaria

	fmt.Print("Inserisci un numero da 3 cifre : ") // messaggio output a terminale
	fmt.Scan(&numero)                              // richiesta del valore iniziale

	// GOOD :
	ultimo := numero / 100                             // calcolo dell'ultima cifra del numero invertito
	mezzo := (numero % 100) / 10                       // calcolo della cifra di mezzo del numero invertito
	primo := numero % 10                               // calcolo della prima cifra del numero invertito
	Invertito := (primo * 100) + (mezzo * 10) + ultimo // calcolo del nuovo numero invertito

	// BETTER :
	cifraUnita := numero / 100                                           // calcolo dell'ultima cifra del numero invertito
	cifraDecine := (numero % 100) / 10                                   // calcolo della cifra di mezzo del numero invertito
	cifraCentinaia := numero % 10                                        // calcolo della prima cifra del numero invertito
	Invertito = (cifraCentinaia * 100) + (cifraDecine * 10) + cifraUnita // calcolo del nuovo numero invertito

	fmt.Println("Risultato : ", Invertito) // stampa a terminale del numero invertito
}
