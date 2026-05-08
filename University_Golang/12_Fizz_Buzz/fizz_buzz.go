/*

TESTO : Scrivere un programma che legge da standard input un numero intero e stampa "Fizz" se il numero è multiplo di 3, "Buzz" se il numero è multiplo
		di 5, "Fizz Buzz" se è multiplo sia di 3 sia di 5, niente altrimenti.

ESEMPIO :
$ go run fizzbuzz.go
Inserisci un numero: 5
Buzz
$ go run fizzbuzz.go
Inserisci un numero: 4

$ go run fizzbuzz.go
Inserisci un numero: 15
Fizz Buzz

$ go run fizzbuzz.go
Inserisci un numero: 6
Fizz

*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile in ingresso

	fmt.Print("Inserisci un numero : ") // messaggio in output a terminale
	fmt.Scan(&numero)                   //richiesta del dato a livello del terminale

	/*

		-----------------------------------------------------------------------------------------------------------------------------------------------------

		GOOD :

			if numero%3 == 0 && numero%5 == 0 { // numeri dividibili sia per 3 che per 5
				fmt.Println("Fizz Buzz")
			} else if numero%3 == 0 { // numeri dividibili solo per 3
				fmt.Println("Fizz")
			} else if numero%5 == 0 { // numero dividibili solo per 5
				fmt.Println("Buzz")
			}

			// Usando else-if con le condizioni escludiamo anche il caso in cui il numero non sia dividibile per i numeri da ma scelti.

		-----------------------------------------------------------------------------------------------------------------------------------------------------

	*/

	// BETTER :

	if numero%15 == 0 { // numeri dividibili sia per 3 che per 5
		fmt.Println("Fizz Buzz")
	} else if numero%3 == 0 { // numeri dividibili solo per 3
		fmt.Println("Fizz")
	} else if numero%5 == 0 { // numero dividibili solo per 5
		fmt.Println("Buzz")
	} else {
		// il numero non è multiplo né di 3 né di 5, non stampiamo nulla
	}

	/*

		-----------------------------------------------------------------------------------------------------------------------------------------------------

		-	Usare ( n%15 == 0 ) invece del doppio ( && ) è più conciso e idiomatico, senza cambiare la struttura if-else-if.

		-	Aggiungere il caso else esplicito. Il tuo commento finale spiega già bene il ragionamento, ma renderlo un else esplicito (anche vuoto) rende
			il codice più leggibile e mostra che il caso è stato considerato, non dimenticato.

		-----------------------------------------------------------------------------------------------------------------------------------------------------

	*/
}
