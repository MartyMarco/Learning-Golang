/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input due numeri interi a e b e stampi a video la somma dei numeri pari
			compresi tra a e b (estremi esclusi).

	ESEMPIO :
		$ go run sommaintervallo.go
		1 9
		Somma = 20
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var numeroA, numeroB, somma int // dichiarazione dei due estremi dell'intervallo

	fmt.Print("Inserisci due numeri : ") // messaggio in output a livello di terminale
	fmt.Scan(&numeroA, &numeroB)         // richiesta dei due dati necessari

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		NOT-SO-GOOD :
			if numeroA < numeroB { // controllo se il primo numero è minore del secondo
				for i := numeroA + 1; i < numeroB; i++ { // ciclo che parte dal numero dopo il numero minore e arriva un numero primo al numero maggiore
					if i%2 == 0 { // controllo se il numero è pari
						somma += i // somma dei numeri pari
					}
				}
			} else { // se il primo numero è maggiore del secondo inverto i numeri rispetto al primo procedimento
				for i := numeroB + 1; i < numeroA; i++ {
					if i%2 == 0 { // controllo se il numero è pari
						somma += i // somma dei numeri pari
					}
				}
			}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/
	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		GOOD :
		// Prima del ciclo, assicurati che numeroA sia sempre il minore
			if numeroA > numeroB {
				numeroA, numeroB = numeroB, numeroA // swap in Go: elegante e diretto
			}

			for i := numeroA + 1; i < numeroB; i++ {
				if i%2 == 0 {
					somma += i
				}
			}

			IMPORTANTE :	il problema principale del codice precedente è aver duplicato lo stesso codice due volte cambiando solo l'ordine delle
							due variabili (numeroA, numeroB) nel ciclo-For, quando si può benissimo evitarlo facendo prima un controllo e nel caso scambiando
							le due variabili in modo che siano in ordine.

							Lo swap a, b = b, a è idiomatico in Go e non richiede una variabile temporanea.
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	// BETTER :
	// Possiamo ottimizzare comunque il codice precendete
	if numeroA > numeroB { // Garantisce che numeroA sia sempre il minore
		numeroA, numeroB = numeroB, numeroA
	}

	start := numeroA + 1 // Parte dal primo numero pari dopo numeroA
	if start%2 != 0 {
		start++
	}
	// Se stiamo cercando solo pari, puoi partire direttamente dal primo pari nell'intervallo invece di controllare ogni numero.

	for i := start; i < numeroB; i += 2 { // una volta che abbiamo controllato che il primo sia pari possiamo andari avanti per 2
		somma += i
	}

	fmt.Printf("Somma = %d\n", somma) // stampa della somma a livello del terminale
}
