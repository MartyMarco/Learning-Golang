/*

TESTO : Scrivere un programma che legge da standard input un numero intero n (specificato senza segno se maggiore o uguale a 0) e stampi a video il
		numero con segno.

ESEMPIO :
$ go run interoconsegno.go
Inserisci numero: 5
+5

$ go run interoconsegno.go
Inserisci numero: 0
0

$ go run interoconsegno.go
Inserisci numero: -5
-5

*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile necessaria

	fmt.Print("Inserisci numero : ") // messaggio di output a livello di terminale
	fmt.Scan(&numero)                // richiesta del dato a livello di terminale

	/*
		GOOD : --------------------------------------------------------------------------------------------------------------------------------------------------

			if numero <= 0 { // condizione se il numero è negativo o uguale a zero
				fmt.Println(numero) // lo stampo esattamento come è
			} else if numero > 0 { // condizione se il numero è positivo (RIDONDANTE meglio usare solo "else")
				fmt.Print("+", numero) // lo stampo aggiungendo un segno positivo davanti al numero
			}

		---------------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	if numero != 0 {
		fmt.Printf("%+d\n", numero)
	} else {
		fmt.Println(numero)
	}

	/*

		Il modo più idiomatico in Go per stampare un intero con segno è usare il verb %+d, che aggiunge automaticamente il segno, questo gestisce tutti e tre i casi
		(positivo, zero, negativo) in una sola riga.

		Il miglioramento principale è %+d: sfrutta direttamente le funzionalità della libreria standard invece di reimplementare manualmente la logica del segno.
		Questo è esattamente lo stile che Go incoraggia.

	*/
}
