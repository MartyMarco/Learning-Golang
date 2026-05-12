/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un numero intero n e stampi a video i divisori propri del numero letto,
			ovvero tutti i suoi divisori escluso il numero stesso. Ad esempio, i divisori del numero 12 sono: 1, 2, 3, 4, 6, 12; quindi
			i divisori propri di 12 sono: 1, 2, 3, 4, 6.

	ESEMPIO :
		$ go run divisori.go
		Inserisci numero: 6
		Divisori di 6: 1 2 3

		$ go run divisori.go
		Inserisci numero: 10
		Divisori di 10: 1 2 5
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var numero int                  // dichiarazione della variabile necessaria per l'inseriemento
	fmt.Print("Inserisci numero: ") // messaggio outuput a livello di terminale
	fmt.Scan(&numero)               // richiesta di inserimento di un numero intero

	// ✅ DOPO : Validazione dell'input
	// Senza questo controllo, con numero = 0 il ciclo non stampa nulla (senza spiegazione),
	// e con numero < 0 il comportamento è indefinito.
	if numero <= 0 {
		fmt.Println("Errore: inserisci un numero intero positivo")
		return
	}

	fmt.Printf("Divisori di %d: ", numero) // messaggio output a livello di terminale
	for i := 1; i < numero; i++ {          // ciclo condizionato escluso lo zero perchè non si può dividere per zero
		if numero%i == 0 { // se la divisione intera non produce resto
			/*
				❌ PRIMA :
					fmt.Print(i)   // allora stampo il numero per cui ho diviso
					fmt.Print(" ") // stampa dello spazio per la formattazione
			*/

			/*
				✅ DOPO : Migliorato — una chiamata sola
				Unifica le due fmt.Print in una sola, ogni chiamata a fmt.Print ha un costo.
				È una piccola ottimizzazione, ma è buona abitudine.
			*/
			fmt.Printf("%d ", i)
		}
	}
}
