/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input una stringa e, considerando l’insieme delle lettere dell'alfabeto
			inglese, ristampi a video la stringa due volte: la prima volta in maiuscolo e la seconda volta in minuscolo.

	ESEMPIO :
		$ go run trasforma.go
		TestoDiProva!!!
		Testo maiuscolo: TESTODIPROVA!!!
		Testo minuscolo: testodiprova!!!

		$ go run trasforma.go
		Testo_Di_PrOvA....
		Testo maiuscolo: TESTO_DI_PROVA....
		Testo minuscolo: testo_di_prova....

-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var testo string                  // dichiarazione della variabile necessaria
	fmt.Print("Inserisci un testo: ") // messaggio di output a livello di terminale
	fmt.Scan(&testo)                  // richiesta di inserimento di una striga senza spazi

	fmt.Print("Testo maiuscolo: ")    // messaggio di output a livello del terminale
	for _, carattere := range testo { // trasformazione di tutto in maiuscolo
		if carattere >= 'a' && carattere <= 'z' {
			fmt.Print(string(carattere - 32))
		} else {
			fmt.Print(string(carattere))
		}
	}

	fmt.Println() // stampa a terminale della newline

	fmt.Print("Testo minuscolo: ")    // messaggio di output a livello del terminale
	for _, carattere := range testo { // trasformazione di tutto in minuscolo
		if carattere >= 'A' && carattere <= 'Z' {
			fmt.Print(string(carattere + 32))
		} else {
			fmt.Print(string(carattere))
		}
	}

	fmt.Println() // stampa a terminale della newline
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
ATTENZIONE :	-	L'aritmetica (±32) è accettabile nel contesto solo perchè il testo dell'esercizio dice esplicitamente "considerando
					l'insieme delle lettere dell'alfabeto inglese", quindi il vincolo ASCII è intenzionale e escludiamo caratteri
					particolari.

				-	Morale: leggere il testo del problema cambia completamente cosa è un "miglioramento" e cosa è invece una scelta
					progettuale consapevole.
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
