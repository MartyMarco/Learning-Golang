/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Create un programma che legge da standard input un intero positivo n e stampi i primi n elementi della successione di
			Fibonacci. La successione di Fibonacci (detta anche successione aurea) è una successione di numeri interi in cui ciascun numero
			è la somma dei due precedenti.

	ESEMPIO :
		$ go run fibonacci.go
		Inserisci un numero intero:
		4

		1 1 2 3

		$ go run fibonacci.go
		Inserisci un numero intero:
		5

		1 1 2 3 5

		$ go run fibonacci.go
		Inserisci un numero intero:
		12

		1 1 2 3 5 8 13 21 34 55 89 144
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var numero int                              // dichiarazione della variabile necessaria
	fmt.Println("Inserisci un numero intero: ") // messaggio output a livello di terminale
	fmt.Scan(&numero)                           // richiesta di inserimento del dato necessario

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ ERRATO :
			for i := 0; i <= numero; i++ {
				somma := i + (i - 1)

				if i == numero { // evito lo spazio superfluo dopo l'ultimo elemento
					fmt.Printf("%d", somma)
				} else { // formatto lo spazio che si crea tra ogni elemento
					fmt.Printf("%d ", somma)
				}
			}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	// ✅ GIUSTO :

	a, b := 0, 1 // i due accumulatori tengono traccia dei due termini precedenti

	for i := 0; i < numero; i++ { // Condizione del ciclo corretta
		a, b = b, a+b // ad ogni step: a diventa il termine corrente, b il prossimo

		if i < numero-1 {
			fmt.Printf("%d ", a)
		} else {
			fmt.Printf("%d\n", a)
		}
	}
}
