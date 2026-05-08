/*

TESTO : Scrivere un programma che legga da standard input un voto v da 0 a 100 e stampi:
			-	Insufficiente se il voto è inferiore a 60.
			-	Sufficiente se il voto è compreso tra 60 (incluso) e 70 (escluso).
			-	Buono se il voto è compreso tra 70 (incluso) e 80 (escluso).
			-	Distinto se il voto è comrpeso tra 80 (incluso) e 90 (escluso).
			-	Ottimo se il voto è compreso tra 90 (incluso) e 100 (incluso).
			-	Errore se il voto è negativo o superiore a 100.

ESEMPIO :
$ go run voto.go
Inserisci il voto: 75
Buono

$ go run voto.go
Inserisci il voto: 90
Ottimo

$ go run voto.go
Inserisci il voto: 110
Errore

*/

package main

import (
	"fmt"
)

func main() {
	var voto int // dichiarazione della variabile in ingresso "voto"

	fmt.Print("Inserisci il voto : ") // messaggio in output a terminale
	fmt.Scan(&voto)                   // richiesta del dato in igresso dal terminale

	/*

		NOT-SO-GOOD : ----------------------------------------------------------------------------------------------------------------------------------------

			if voto < 60 { // 0 < voto < 60
				fmt.Println("Insufficiente")
			} else if voto >= 60 && voto < 70 { // 60 <= voto < 70
				fmt.Println("Sufficiente")
			} else if voto >= 70 && voto < 80 { // 70 <= voto < 80
				fmt.Println("Buono")
			} else if voto >= 80 && voto < 90 { // 80 <= voto < 90
				fmt.Println("Distinto")
			} else if voto >= 90 && voto <= 100 { // 90 <= voto <= 100
				fmt.Println("Ottimo")
			} else if voto > 100 || voto < 0 { // voto < 0 || voto > 100
				fmt.Println("Errore")
			}

		IMPORTANTE :	il controllo dell'errore non funziona mai, questo perchè la condizione è alla fine e nel caso la variabile sia negativa rientria
						nel primo ramo condizionale, ossia il ramo dell' Insufficiente.

		SOLUZIONE :		Spostare l'ultimo ramo condizionale all'inizio.

		----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	/*

		GOOD : ---------------------------------------------------------------------------------------------------------------------------------------------

			if voto > 100 || voto < 0 { // voto < 0 || voto > 100
				fmt.Println("Errore")
			} else if voto < 60 { // 0 <= voto < 60
				fmt.Println("Insufficiente")
			} else if voto >= 60 && voto < 70 { // 60 <= voto < 70
				fmt.Println("Sufficiente")
			} else if voto >= 70 && voto < 80 { // 70 <= voto < 80
				fmt.Println("Buono")
			} else if voto >= 80 && voto < 90 { // 80 <= voto < 90
				fmt.Println("Distinto")
			} else if voto >= 90 && voto <= 100 { // 90 <= voto <= 100
				fmt.Println("Ottimo")
			}

		IMPORTANTE :	ci sono condizioni ridondanti nei rami interni, dato che dopo il primo controllo i range sono già garantiti.
						Il ragionamento chiave è: in una catena if/else if, se arrivi a un ramo significa che tutti i rami precedenti erano falsi.
						Quindi voto >= 60 nel terzo else if è già implicito — non serve scriverlo.

		SOLUZIONE :		Tenere solo le condizioni che "chiudono" l'intervallo.

		----------------------------------------------------------------------------------------------------------------------------------------------------

	*/

	// BETTER :

	if voto < 0 || voto > 100 { // fuori range
		fmt.Println("Errore")
	} else if voto < 60 { // 0 <= voto < 60  (già sappiamo voto >= 0)
		fmt.Println("Insufficiente")
	} else if voto < 70 { // 60 <= voto < 70 (già sappiamo voto >= 60)
		fmt.Println("Sufficiente")
	} else if voto < 80 { // 70 <= voto < 80 (già sappiamo voto >= 70)
		fmt.Println("Buono")
	} else if voto < 90 { // 80 <= voto < 90 (già sappiamo voto >= 80)
		fmt.Println("Distinto")
	} else { // 90 <= voto <= 100 (tutti gli altri casi)
		fmt.Println("Ottimo")
	}

}
