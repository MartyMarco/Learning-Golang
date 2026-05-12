/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che, dopo aver letto da standard input un numero intero n, chiede all'utente di inserire n numeri
	 		interi (sempre da standard input).

			Dopo aver letto gli n numeri interi, il programma deve stampare:
				-	la somma degli n numeri letti;
				-	il minimo tra i numeri letti;
				-	il massimo tra i numeri letti;
				-	il numero di interi letti strettamente positivi (maggiori di 0), strettamente negativi (minori di 0), e nulli.

	ESEMPIO :
		$ go run nnumeri.go
		9
		Inserisci 9 numeri:
		1 -2 3 -4 5 -6 7 -8 9
		somma = 5
		valore minimo = -8
		valore massimo = 9
		interi > 0 = 5
		interi < 0 = 4
		interi = 0 = 0
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import "fmt"

func main() {

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		GOOD :

			var quanto, numero, somma, minimo, massimo, positivi, negativi, nulli int // dichiarazione di tutte le variabili necessarie

			fmt.Print("Quanti numeri vuoi inserire : ") // messaggio a livello di output a terminale
			fmt.Scan(&quanto)                           // richiedere il dato al terminale

			fmt.Printf("Inserisci %d numeri : \n", quanto) // messaggio a livello di output a terminale

			for i := 0; i < quanto; i++ { // ciclo per l'ingresso dei valori

				fmt.Scan(&numero) // richiedere il dato al terminale
				somma += numero   // calcolo della somma di tutti i numeri inseriti

				if i == 0 { // nel primo giro del ciclo
					minimo = numero  // segno tutti i valori come il primo valore inserito
					massimo = numero // segno tutti i valori come il primo valore inserito
				}
				if numero < minimo { // se il numero è minore di quello inserito nella variabile minore
					minimo = numero // sostituisco nella variabile minore il nuovo valore
				}
				if numero > massimo { // se il numero è maggiore di quello inserito nella variabile maggiore
					massimo = numero // sostituisco nella variabile maggiore il nuovo valore
				}
				if numero > 0 { // se il numero inserito è positivo
					positivi++ // aumento di uno il contatore dei numeri positivi
				} else if numero < 0 { // se il numero inseritgo è negativo
					negativi++ // aumento di uno il contatore dei numeri negativi
				} else { // se il numero inserito è nullo (numero == 0)
					nulli++ // aumento di uno il contatore dei numeri nulli
				}

			}

			// Stampo a terminale tutti i dati calcolati precedenti :
			fmt.Printf("somma = %d\n", somma)
			fmt.Printf("valore minimo = %d\n", minimo)
			fmt.Printf("valore massimo = %d\n", massimo)
			fmt.Printf("interi > 0 = %d\n", positivi)
			fmt.Printf("interi < 0 = %d\n", negativi)
			fmt.Printf("interi = 0 = %d\n", nulli)
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	// BETTER :

	/*	❌ PRIMA — dichiarare tutto insieme, anche ciò che appartiene al ciclo
		✅ DOPO — dichiarare separatamente le variabili per scope logico

		Perché : ogni variabile dovrebbe vivere nel blocco più piccolo in cui è necessaria. È un principio fondamentale di leggibilità
		e correttezza.
	*/
	var quanto int
	fmt.Scan(&quanto)

	// ✅ Aggiungi subito dopo
	if quanto <= 0 {
		fmt.Println("Errore: inserisci un numero maggiore di zero.")
		return
	}

	/*	❌ PRIMA :	usare if i == 0 per inizializzare min/max è un "trucco" che dipende dall'indice del ciclo, non dal valore logico.
		✅ DOPO :	utilizziamo un primo scan al di fuori del ciclo-for per leggere il primo valore e inserirlo come valore iniziale sia minimo
					 massimio.
	*/

	var somma, minimo, positivi, negativi, nulli int // inizializza prima del ciclo usando il primo Scan

	fmt.Scan(&minimo) // leggi il primo numero
	massimo := minimo
	somma += minimo

	for i := 1; i < quanto; i++ { // parti da 1, il primo è già letto
		var numero int // numero vive SOLO dentro il ciclo, dove ha senso
		fmt.Scan(&numero)
		if numero < minimo { // se il numero è minore di quello inserito nella variabile minore
			minimo = numero // sostituisco nella variabile minore il nuovo valore
		}
		if numero > massimo { // se il numero è maggiore di quello inserito nella variabile maggiore
			massimo = numero // sostituisco nella variabile maggiore il nuovo valore
		}
		if numero > 0 { // se il numero inserito è positivo
			positivi++ // aumento di uno il contatore dei numeri positivi
		} else if numero < 0 { // se il numero inseritgo è negativo
			negativi++ // aumento di uno il contatore dei numeri negativi
		} else { // se il numero inserito è nullo (numero == 0)
			nulli++ // aumento di uno il contatore dei numeri nulli
		}
	}

	// Stampo a terminale tutti i dati calcolati precedenti :
	fmt.Printf("somma = %d\n", somma)
	fmt.Printf("valore minimo = %d\n", minimo)
	fmt.Printf("valore massimo = %d\n", massimo)
	fmt.Printf("interi > 0 = %d\n", positivi)
	fmt.Printf("interi < 0 = %d\n", negativi)
	fmt.Printf("interi = 0 = %d\n", nulli)

}
