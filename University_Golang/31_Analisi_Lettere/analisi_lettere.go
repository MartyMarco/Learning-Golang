/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input una stringa senza spazi e, considerando solamente l’insieme delle lettere
			dell'alfabeto inglese, stampi:
					-	il numero di lettere maiuscole;
					-	il numero di lettere minuscole;
					-	il numero di consonanti;
					-	il numero di vocali.

	NOTA / SUGGERIMENTO :	Le lettere minuscole e maiuscole dell'alfabeto inglese (o latino) fanno parte dello standard US-ASCII, un
							sottoinsieme dello standard Unicode. I 128 valori numerici (o codici) dei caratteri US-ASCII sono compresi tra
							0 e 127. I codici delle lettere minuscole e quelli delle lettere maiuscole (così come quelli delle cifre) occupano
							posizioni contigue nella tabella US-ASCII (Unicode). Il codice di un qualsiasi carattere si ottiene racchiudendo
							il simbolo corrispondente tra singoli apici (ad esempio ‘A’ corrisponde a 65 e ‘a’ a 97). Quindi per manipolare i
							caratteri non serve ricordarne i valori numerici.

	ESEMPIO :
		$ go run analisi.go
		Ciaoà
		Maiuscole: 1
		Minuscole: 3
		Vocali: 3
		Consonanti: 1

		$ go run analisi.go
		Certo!Sto,bene!ìììììì
		Maiuscole: 2
		Minuscole: 10
		Vocali: 5
		Consonanti: 7

		$ go run analisi.go
		aaAA
		Maiuscole: 2
		Minuscole: 2
		Vocali: 4
		Consonanti: 0
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var parola string                                 // dichiarazione di una variabile di caratteri
	fmt.Print("Inserisci una stringa di caratteri: ") // messaggio output a livello di terminale
	fmt.Scan(&parola)                                 // richiesta di inserimento di una stringa di caratteri senza spazi

	var maiuscole, minuscole, vocali, consonanti int // dichiarazione delle variabili responsabili dei contatori

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ NOT-SO-GOOD :
		for i := 0; i < len(parola); i++ { // ciclo condizionale sulla lunghezza della parola
			carattere := rune(parola[i]) // conversione da byte delle stringa a rune

			if carattere >= 65 && carattere <= 90 { // condizione in cui il carattere è maiuscolo
				maiuscole++

				if carattere == 65 || carattere == 69 || carattere == 73 || carattere == 79 || carattere == 85 { // oltre anche vocale
					vocali++
				} else { // oltre che a essere maiuscolo è una consonante
					consonanti++
				}

			} else if carattere >= 97 && carattere <= 122 { // condizione in cui il carattere è minuscolo
				minuscole++

				if carattere == 97 || carattere == 101 || carattere == 105 || carattere == 111 || carattere == 117 { // oltre anche vocale
					vocali++
				} else { // oltre che a essere minuscolo è una consonante
					consonanti++
				}

			}
		}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	/*
		✅ BETTER :
		Problema con il codice precedente : parola[i] restituisce un byte, non una runa. Questo va bene per ASCII puro, ma è idiomaticamente scorretto in
		Go, dove le stringhe sono UTF-8. Il cast rune(parola[i]) non gestisce correttamente caratteri multi-byte.

		Meglio usare le costanti carattere invece dei magic number, chiunque legga il codice deve sapere a memoria la tabella ASCII. Il suggerimento
		del testo lo dice esplicitamente.
	*/
	for _, carattere := range parola { // Idiomatico Go - range decodifica UTF-8 automaticamente

		if carattere >= 'A' && carattere <= 'Z' { // lettere maiuscole
			maiuscole++
			if carattere == 'A' || carattere == 'E' || carattere == 'I' || carattere == 'O' || carattere == 'U' {
				vocali++
			} else {
				consonanti++
			}

		} else if carattere >= 'a' && carattere <= 'z' { // lettere minuscole
			minuscole++
			if carattere == 'a' || carattere == 'e' || carattere == 'i' || carattere == 'o' || carattere == 'u' {
				vocali++
			} else {
				consonanti++
			}
		}
	}

	fmt.Printf("Maiuscole: %d\n", maiuscole)   // messaggio output MAIUSCOLE
	fmt.Printf("Minuscole: %d\n", minuscole)   // messaggio output MINUSCOLE
	fmt.Printf("Vocali: %d\n", vocali)         // messaggio output VOCALI
	fmt.Printf("Consonanti: %d\n", consonanti) // messaggio output CONSONANTI
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
N.B.	Il principio chiave è: il codice si legge molte più volte di quante si scrive — ogni miglioramento riduce il
		carico cognitivo per chi legge (incluso il tuo futuro te).
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
