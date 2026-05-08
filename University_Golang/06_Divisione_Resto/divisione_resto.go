/*

TESTO : Scrivere un programma che prenda in ingresso due numeri e stampi il quoziente ed il resto della divisione del primo numero per il secondo.

ESEMPIO :
$ go run convertisec.go
Inserisci due numeri: 5 3
Quoziente: 1
Resto: 2

*/

package main

import (
	"fmt"
)

func main() {
	var dividendo, divisore int // dichiarazione delle due variabili necessarie

	fmt.Print("Inserisci due numeri : ") // messaggio output a terminale
	fmt.Scan(&dividendo, &divisore)      // richiesta dei due dati dal terminale

	quoziente := dividendo / divisore // calcolo del quoziente
	resto := dividendo % divisore     // calcolo del resto

	fmt.Println("Quoziente : ", quoziente) // restituzione del valore del Quoziente al terminale
	fmt.Println("Resto : ", resto)         // restituzione del valore del Resto al terminale
}
