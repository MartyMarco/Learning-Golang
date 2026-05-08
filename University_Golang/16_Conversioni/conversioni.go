/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO : Scrivere un programma che:
				-	legga da standard input un valore intero che specifica il tipo di conversione da effettuare, gestendo opportunamente
					l’insertimento di un valore di scelta non compreso tra 1 e 8:
							1: secondi (inseriti dall’utente) in ore
							2: secondi inseriti dall’utente in minuti
							3: minuti inseriti dall’utente in ore
							4: minuti inseriti dall’utente in secondi
							5: ore inserite dall’utente in secondi
							6: ore inserite dall’utente in minuti
							7: minuti inseriti dall’utente in giorni e ore
							8: minuti inseriti dall’utente in anni e giorni

				-	legga da standard input un valore reale da convertire;

				-	stampi a video il valore convertito come da esempio di esecuzione.

	ESEMPIO :
		$ go run conversioni.go
		Scegli la conversione:
		1) secondi -> ore
		2) secondi -> minuti
		3) minuti -> ore
		4) minuti -> secondi
		5) ore -> secondi
		6) ore -> minuti
		7) minuti -> giorni e ore
		8) minuti -> anni e giorni
		: 8
		Inserisci il valore da convertire: 7200
		7200 minuti corrispondono a 0 anni e 5 giorni

		$ go run conversioni.go
		Scegli la conversione:
		1) secondi -> ore
		2) secondi -> minuti
		3) minuti -> ore
		4) minuti -> secondi
		5) ore -> secondi
		6) ore -> minuti
		7) minuti -> giorni e ore
		8) minuti -> anni e giorni
		: 1
		Inserisci il valore da convertire: 3618
		3618 secondi corrispondono a 1.005 ore

		$ go run conversioni.go
		Scegli la conversione:
		1) secondi -> ore
		2) secondi -> minuti
		3) minuti -> ore
		4) minuti -> secondi
		5) ore -> secondi
		6) ore -> minuti
		7) minuti -> giorni e ore
		8) minuti -> anni e giorni
		: 9
		Scelta errata
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var scelta, valore int // dichiarazione delle due variabili necessarie

	fmt.Println("Scegli la conversione:") // stampa a livello del terminale del menu delle opzioni
	fmt.Println("1) secondi -> ore")
	fmt.Println("2) secondi -> minuti")
	fmt.Println("3) minuti -> ore")
	fmt.Println("4) minuti -> secondi")
	fmt.Println("5) ore -> secondi")
	fmt.Println("6) ore -> minuti")
	fmt.Println("7) minuti -> giorni e ore")
	fmt.Println("8) minuti -> anni e giorni")
	fmt.Print(": ")
	fmt.Scan(&scelta) // richiediamo a livello del terminale il valore della scelta dell'operazione

	if scelta < 1 || scelta > 8 { // condizionale nel caso la selta dell'utente non ricadesse nelle possibilità
		fmt.Println("Scelta errata")
	} else {
		fmt.Print("Inserisci il valore da convertire: ") // messaggio in output a terminale
		fmt.Scan(&valore)                                // richiesta in ingressa del dato del valore dal terminale

		if scelta == 1 { // OPZIONE 1 : conversione da secondi a ore
			ore := float64(valore) / 3600
			fmt.Printf("%d secondi corrispondono a %g ore\n", valore, ore)

		} else if scelta == 2 { // OPZIONE 2 : conversione da secondi a minuti
			minuti := float64(valore) / 60
			fmt.Printf("%d secondi corrispondono a %g minuti\n", valore, minuti)

		} else if scelta == 3 { // OPZIONE 3 : conversioni da minuti a ore
			ore := float64(valore) / 60
			fmt.Printf("%d minuti corrispondono a %g ore\n", valore, ore)

		} else if scelta == 4 { // OPZIONE 4 : conversione da minuti a secondi
			secondi := valore * 60
			fmt.Printf("%d minuti corrispondono a %d secondi\n", valore, secondi)

		} else if scelta == 5 { // OPZIONE 5 : conversione da ore a secondi
			secondi := valore * 3600
			fmt.Printf("%d ore corrispondono a %d secondi\n", valore, secondi)

		} else if scelta == 6 { // OPZIONE 6 : conversione da ore a minuti
			minuti := valore * 60
			fmt.Printf("%d ore corrispondono a %d minuti\n", valore, minuti)

		} else if scelta == 7 { // OPZIONE 7 : conversione da minuti a giorni e ore
			/*
				❌ SBAGLIATO :
					giorni := (valore / 60) / 24
					ore := valore % 60 // residuo in MINUTI, non in ore!
			*/

			oreTotali := valore / 60
			giorni := oreTotali / 24
			ore := oreTotali % 24

			fmt.Printf("%d minuti corrispondono a %d giorni e %d ore\n", valore, giorni, ore)

		} else if scelta == 8 { // OPZIONE 8 : conversione da minuti a anni e giorni
			/*
				❌ SBAGLIATO :
				giorni := (valore / 60) / 24 // giorni NON è il residuo!
				anni := giorni / 365
			*/

			giorniTotali := valore / 60 / 24
			anni := giorniTotali / 365
			giorni := giorniTotali % 365

			fmt.Printf("%d minuti corrispondono a %d anni e %d giorni\n", valore, anni, giorni)

		}
	}
}
