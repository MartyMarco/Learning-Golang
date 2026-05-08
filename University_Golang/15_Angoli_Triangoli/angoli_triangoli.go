/*
	-----------------------------------------------------------------------------------------------------------------------------------------------------
		TESTO : Scrivere un programma che legga da standard input le ampiezze di due angoli di un triangolo e stampi, se possibile, l'ampiezza del
				terzo angolo.

		SUGGERIMENTO : ricordatevi che in un triangolo la somma delle ampiezze degli angoli interni è sempre 180°.

		ESEMPIO :
			$ go angolitriangolo.go
			Inserire le ampiezze dei due angoli: 50 60
			Ampiezza terzo angolo = 70°

			$ go angolitriangolo.go
			Inserire le ampiezze dei due angoli: 150 70
			I due angoli non appartengono ad un triangolo
	-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var angolo1, angolo2 int // dichiarazione delle variabili necessarie dei due angoli

	fmt.Print("Inserire le ampiezze dei due angoli : ") // messaggio di output a livello di terminale
	fmt.Scan(&angolo1, &angolo2)                        // richiesta dei dati degli angoli a livello di terminale

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		GOOD :

			risultato := 180 - (angolo1 + angolo2) // calcolo del terzo angolo del triangolo

			if risultato > 0 { // condizionale se il terzo angolo esiste
				fmt.Printf("Ampiezza terzo angolo = %d°\n", risultato)
			} else { // terzo angolo negativo o nullo, inesistente
				fmt.Printf("I due angoli non appartengono ad un triangolo\n")
			}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	// BETTER :

	if angolo1 <= 0 || angolo2 <= 0 {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	} else if (angolo1 + angolo2) >= 180 {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	} else {
		risultato := 180 - (angolo1 + angolo2)
		fmt.Printf("Ampiezza terzo angolo = %d°\n", risultato)
	}

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
			-	Attualmente il codice non controlla se i singoli angoli sono validi (es. negativi o zero), questo è un uso classico della
				catena if-else-if.

			-	Nel tuo codice originale risultato viene calcolato sempre, anche quando non serve. È buona abitudine calcolarlo solo quando i
				dati sono validi.
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/
}
