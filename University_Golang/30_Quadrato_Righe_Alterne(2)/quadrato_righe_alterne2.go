/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un numero intero n e che, come mostrato nell'Esempio di esecuzione,
			stampi a video un quadrato di n righe costituite ciascuna da n simboli intervallati da spazi, alternando fra loro righe
			costituite solo da simboli * (asterisco) intervallati da spazi, righe costituite solo da simboli + (più) intervallati da spazi
			e righe costituite solo da simboli o (lettera o) intervallati da spazi.

	ESEMPIO :
		$ go run quadrato_righe_alterne_2.go
		Inserisci un numero: 5
		* * * * *
		+ + + + +
		o o o o o
		* * * * *
		+ + + + +
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var dimensione int                 // dichiarazione della variabile che definisce la dimensione del quadrato
	fmt.Print("Inserisci un numero: ") // messaggio di output a livello di terminale
	fmt.Scan(&dimensione)              // richiesta di inserimento del dato

	if dimensione <= 0 { // controllo del valore dato in ingresso
		fmt.Println("Errore: dimensione negativa o nulla")
		return // interruzione del programma
	}

	for i := 0; i < dimensione; i++ { // primo ciclo annidato

		/*
			❌ SBAGLIATO :
				simbolo := "*" // prima riga
				if i%2 == 0 {
					simbolo = "+" // seconda riga
				} else if i%3 == 0 {
					simbolo = "o" // terza riga
				}
		*/

		// ✅ CORRETTO :
		simbolo := "*" // i%3 == 0 → riga 0, 3, 6...
		if i%3 == 1 {
			simbolo = "+" // i%3 == 1 → riga 1, 4, 7...
		} else if i%3 == 2 {
			simbolo = "o" // i%3 == 2 → riga 2, 5, 8...
		}

		for j := 0; j < dimensione; j++ { // secondo ciclo annidato
			if j < dimensione-1 {
				fmt.Print(simbolo, " ") // intervallo spazi per creare la riga
			} else {
				fmt.Print(simbolo) // cancello lo spazio in fondo alla riga
			}
		}

		fmt.Println() // aggiungo la newline per iniziare una nuova riga
	}
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
N.B.	-	Chiamare fmt.Print per ogni carattere è costoso — ogni chiamata ha overhead di I/O.
			Meglio costruire la stringa con ( strings.Repeat ) e ( strings.Join )
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
