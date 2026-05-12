/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	(Definizione) : 	Una parola è palindroma se può essere letta normalmente, da sinistra verso destra, sia viceversa, cioè da destra verso
						sinistra.

	TESTO :	Scrivere un programma che:
				-	legga da standard input una stringa senza spazi;
				-	stampi a video il messaggio Palindroma nel caso in cui la stringa letta sia palindroma e Non palindroma altrimenti.

	ESEMPIO :
		$ go run palindroma.go
		anna
		Palindroma

		$ go run palindroma.go
		anni
		Non palindroma

		$ go run palindroma.go
		osso
		Palindroma

		$ go run palindroma.go
		èssè
		Palindroma

		$ go run palindroma.go
		èsse
		Non palindroma
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var parola string                   // dichiarazine della variabile necessaria per l'inseriemento della parola
	fmt.Print("Inserisci una parola: ") // messaggio di output a livello di terminale
	fmt.Scan(&parola)                   // richiesta di inserimento di una stringa di caratteri

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ NOT-SO-GOOD :
			controllo := false
			for i, caratteri := range parola {

			// SBAGLIATO — quando i == 0, len(parola)-i == len(parola) → out of bounds!

				if string(caratteri) == string(parola[len(parola)-i]) {
					controllo = true
				} else {
					controllo = false
				}
			}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	/*
		✅ BETTER :
	*/
	controllo := true           // assumiamo subito che sia vera
	caratteri := []rune(parola) // conversione in rune per supportare Unicode
	lunghezza := len(caratteri) // numero di caratteri della stringa inserita dall'utente

	for i := 0; i < lunghezza/2; i++ { // si ferma a metà — ogni coppia viene confrontata una volta sola
		if caratteri[i] != caratteri[lunghezza-1-i] { // gli indici vanno da (0 a len-1)
			controllo = false // controlliamo solo il caso negativo in cui devo escludere la parola
			break             // ottimizzazione: inutile continuare dopo aver trovato un mismatch
		}
	}

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ GOOD :
			if controllo == true {
				fmt.Println("Palindroma")
			} else {
				fmt.Println("Non palindroma")
			}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	/*
		✅ BETTER :
		meno ridendante e più idiomatico per il linguaggio utilizzato
	*/

	if controllo {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}
}
