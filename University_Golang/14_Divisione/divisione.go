/*

TESTO : Scrivere un programma che legga da standard input due numeri a e b all'interno di due variavili di tipo int e calcoli il risultato
		della divisione a/b. Se b è uguale a 0, il programma stampa Impossibile.

ESEMPIO :
$ go run divisione.go
Inserisci due numeri:
5 2
Quoziente = 2.5

$ go run divisione.go
Inserisci due numeri:
5 0
Impossibile

*/

package main

import (
	"fmt"
)

func main() {
	var dividendo, divisore int // dichiarazione delle due varabili necessarie per la divisione

	fmt.Println("Inserisci due numeri : ") // messaggio in uscita a livello di terminale
	fmt.Scan(&dividendo, &divisore)        // richiesta di dati a livello di terminale

	if divisore != 0 { // escludiamo l'unica divisione non possibile che è quella in cui si divide per 0
		quoziente := float64(dividendo) / float64(divisore)

		/*

			fmt.Println("Quoziente = ", quoziente)
			-	Meglio utilizzare fmt.Printf() perchè non crea un doppio spazio inutile.
			-	Il verbo %g rimuove gli zeri decimali inutili

		*/

		fmt.Printf("Quoziente = %g\n", quoziente)

	} else { // messaggio di errore nel caso l'utente inserisca 0 come divisore della divisione
		fmt.Println("Impossibile")
	}
}
