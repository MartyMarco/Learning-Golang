/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un numero intero n e stampi a video un quadrato di n asterischi intervallati
			da spazi come mostrato nell'Esempio di esecuzione.

	SUGGERIMENTO : potete utilizzare due cicli for annidati.

	ESEMPIO :
		$ go run quadratoasterischi.go
		Inserisci un numero: 3
		* * *
		* * *
		* * *
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var dimensione int                 // dichiarazione della variabile che rappresenta la dimensione del quadrato
	fmt.Print("Inserisci un numero: ") // messaggio output a livello di terminale
	fmt.Scan(&dimensione)              // richiesta del dato

	// ✅ Validare l'input : Il codice non gestisce input invalidi (zero, negativi, errori di lettura).
	if dimensione <= 0 {
		fmt.Println("Errore: inserisci un numero positivo.")
		return
	}

	for i := 0; i < dimensione; i++ {
		for j := 0; j < dimensione; j++ {
			/*
				❌ PRIMA : stampa uno spazio alla fine della riga
				fmt.Print("* ") // stampa senza newline per creare la linea del quadrato

				Il tuo codice stampa ("* ") anche dopo l'ultimo asterisco di ogni riga, quindi ogni riga ha uno spazio in coda.
				Non si vede, ma è tecnicamente scorretto.
			*/

			// ✅ DOPO :
			if j < dimensione-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("*") // ultimo asterisco senza spazio
			}
		}
		/*
			❌ PRIMA : Meno idiomatico
			fmt.Print("\n") // stampa della newline per iniziare una nuova linea al prossimo ciclo
		*/

		/*
			✅ DOPO : Idiomatico in Go
			Perché: fmt.Println() senza argomenti stampa esattamente un \n ed è la forma convenzionale in Go.
		*/
		fmt.Println()
	}
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
N.B.	-	Usare ( strings.Repeat ) al posto del ciclo interno, questo perchè il ciclo interno fa una sola cosa ripetitiva:
			concatenare ("* ") per n volte. Go ha uno strumento nativo per questo.

			Perché è meglio: evita n chiamate a fmt.Print per riga, costruisce la stringa in memoria e la stampa in un colpo solo —
			meno I/O, più efficiente.


		-	fmt.Printf("%c ", asterisco) // %c stampa una rune come carattere
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
