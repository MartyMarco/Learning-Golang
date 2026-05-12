/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione, stampi
			a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe costituite
			solo da simboli * (asterisco) intervallati da spazi e righe costituite solo da simboli + (più) intervallati da spazi.

	SUGGERIMENTO : potete utilizzare due cicli for annidati ed utilizzare l'operatore % per distinguere le righe pari da quelle dispari.

	ESEMPIO :
		$ go run quadrato_righe_alterne_1.go
		Inserisci un numero: 5
		* * * * *
		+ + + + +
		* * * * *
		+ + + + +
		* * * * *
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var dimensione int                 // dichiarazione della variabile che rappresenta la dimensione del quadrato
	fmt.Print("Inserisci un numero: ") // messaggio output a livello di terminale
	fmt.Scan(&dimensione)              // richiesta di inserimento del dato

	// ✅ Validazione dell'input : il mio codice non gestiva prima input non validi
	if dimensione <= 0 {
		fmt.Println("Errore: inserisci un numero positivo")
		return
	}

	/*
		---------------------------------------------------------------------------------------------------------------------------------------
		❌ PRIMA :
		   		for i := 0; i < dimensione; i++ { // cicli annidati : righe
		   		for j := 0; j < dimensione; j++ { // cicli annidati : colonne

		   			if i%2 == 0 && j < dimensione-1 { // righe pari e non ultimo elemento
		   				fmt.Print("* ")
		   			} else if i%2 == 0 && j == dimensione-1 { // righe pari e ultimo elemento
		   				fmt.Print("*")
		   			} else if i%2 != 0 && j < dimensione-1 { // righe dispari e non ultimo elemento
		   				fmt.Print("+ ")
		   			} else { // righe dispari e ultimo elemento
		   				fmt.Print("+")
		   			}
		---------------------------------------------------------------------------------------------------------------------------------------
	*/

	/*
		✅ DOPO : condizioni ridondanti
		Nel tuo codice, controlli i%2 due volte per ogni elemento: prima per la riga pari, poi per quella dispari.
		Puoi separare le due responsabilità.
	*/
	for i := 0; i < dimensione; i++ { // primo ciclo annidato
		simbolo := "+" // simbolo nel caso delle righe dispari
		if i%2 == 0 {
			simbolo = "*" // cambio il simbolo nel caso delle righe dispari
		}

		for j := 0; j < dimensione; j++ { // secondo ciclo annidato
			if j < dimensione-1 { // intervallo gli spazi nella costruzione della riga
				fmt.Print(simbolo + " ")
			} else { // tolgo lo spazio in fondo all'ultimo elemento della riga
				fmt.Print(simbolo)
			}
		}
		fmt.Println() // inserimento della newline
	}
}
