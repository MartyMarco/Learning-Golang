/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Il programma legge da tastiera una serie di numeri maggiori di -1 e dopo ogni lettura stampa "crescente" se il nuovo valore è
			maggiore o uguale al precedente e "decrescente" altrimenti. Si assuma che il primo numero inserito sia sempre preceduto da uno 0.
			Il programma si ferma quando il numero in input è <= -1. Utilizzare fmt.Scan() per la lettura.

	ESEMPIO :
		$ go run sequenza.go
		inserisci una sequenza di numeri:

		1 3 2 5 6 3 1 -1

		crescente
		crescente
		decrescente
		crescente
		crescente
		decrescente
		decrescente
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inserisci una sequenza di numeri:") // messaggio output a livello di terminale

	primo := 0 // dichiarazione e inizializzazione del primo numero
	for {
		var numero int    // dichiarazione della variabile nel blocco in cui è necessaria
		fmt.Scan(&numero) // richiesta di inserimento della variabile a livello del terminale

		if numero < 0 {
			return // casistica di interruzione del ciclo
		}
		if primo <= numero { // ✅ Dopo - Il testo dice "crescente se maggiore o uguale"
			fmt.Println("crescente")
		} else {
			fmt.Println("decrescente")
		}
		primo = numero // sostituzione del numero nella posizione del primo
	}
}
