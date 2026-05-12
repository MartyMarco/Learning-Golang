/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un intero n. Il programma deve stampare a video la sequenza di numeri che
			vanno da 1 a n come mostrato nell'Esempio d'esecuzione. In particolare:

				-	ogni numero divisibile per 3 deve essere rimpiazzato dalla parola "Fizz";
				-	ogni numero divisibile per 5 deve essere rimpiazzato dalla parola "Buzz";
				-	ogni numero divisibile sia per 3 sia per 5 deve essere sostituito da "FizzBuzz".

	ESEMPIO :
		$ go run fizzbuzz.go
		6
		1 2 Fizz 4 Buzz Fizz

		$ go run fizzbuzz.go
		20
		1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var numero int                            // dichiarazione della variabile necessaria per l'inserimento
	fmt.Print("Inserisci un numero intero: ") // messaggio output a livello di terminale
	fmt.Scan(&numero)                         // richiesta di inserimento di un dato a livello di terminale

	// ✅ DOPO : Validazione dell'input
	// Se l'utente inserisce 0 o un numero negativo, il ciclo non esegue nulla e il programma termina silenziosamente.
	// Meglio comunicarlo esplicitamente.
	if numero <= 0 {
		fmt.Println("Errore: inserisci un numero positivo")
		return
	}

	for i := 1; i <= numero; i++ { // ciclo condizionato
		if i%15 == 0 { // numero divisibile sia per 3 che per 5
			fmt.Print("FizzBuzz")
		} else if i%3 == 0 { // numero divisibile per 3
			fmt.Print("Fizz")
		} else if i%5 == 0 { // numero divisibile per 5
			fmt.Print("Buzz")
		} else { // stampa di tutti gli altri numeri della serie
			fmt.Print(i)
		}

		// stampa dello spazio per la formattazione del messaggio a livello di terminale
		if i < numero { // questo per evitare di stampare lo spazio dopo l'ultimo elemento
			fmt.Print(" ")
		}
	}
}
