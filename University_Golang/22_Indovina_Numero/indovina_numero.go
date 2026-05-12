/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che:
				1)	Genera in modo casuale un numero intero compreso nell'intervallo che va da 1 a 100, estremi inclusi (sia numeroGenerato
					la variabile intera in cui viene memorizzato il numero generato, come indicato nella consegna deve valere
					1<= numeroGenerato <= 100).

				2)	Chiede iterativamente all'utente di inserire da standard input un numero intero; ad ogni iterazione il programma controlla
					se il numero inserito è uguale al numero memorizzato in numeroGenerato:

							-	se sono uguali, il programma termina;
							-	se sono diversi, il programma fornisce un suggerimento specificando se il numero inserito è più alto o più
								basso di quello memorizzato in numeroGenerato.

			L'output stampato a video dal programma deve essere quello riportato nell'Esempio d'esecuzione (eccezion fatta per i numeri
			inseriti dall'utente).

	SUGGERIMENTO : per generare in modo casuale un numero intero, potete utilizzare le funzioni dei package math/rand e time come mostrato nel
	seguente frammento di codice:
			inizializzazione del generatore di numeri casuali
			// rand.Seed(int64(time.Now().Nanosecond()))
			generazione di un numero casuale compreso nell'intervallo
   			che va da 0 a 99 (estremi inclusi)
			// var numeroGenerato int = rand.Intn(100)

	ESEMPIO :
		$ go run indovina.go
		Tentativo n° 1: 50
		Troppo basso! Riprova!
		Tentativo n° 2: 75
		Troppo alto! Riprova!
		Tentativo n° 3: 63
		Troppo basso! Riprova!
		Tentativo n° 4: 68
		Troppo alto! Riprova!
		Tentativo n° 5: 66
		Hai indovinato in 5 tentativi!
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ PRIMA	:
				-	Deprecato da Go 1.20, questa inizializzazione del generatore di numeri casuali è superata.
				-	rand.Intn(100) genera 0–99, non 1–100 come richiesto, quindi il numero generato non rispetta la consegna.

			rand.Seed(int64(time.Now().Nanosecond())) // inizializzazione del generatore di numeri casuali
			var numeroGenerato int = rand.Intn(100)   // generazione di un numero casuale compreso nell'intervallo che va da 0 a 99 (estremi inclusi)
			var contatore int                         // dichiarazione del contatore per le risposte all'utente
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	/* ✅ DOPO :
	-	Aggiungere +1 al numero generato è una risoluzione di un bug logico, non solo stilistico.
	-	Precedentemente ho utilizzato uno stile prolisso non necessario quando è possibile usare uno stile più idiomatico molto caratteristico
		di Go.
	*/
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numeroGenerato := r.Intn(100) + 1
	contatore := 0

	for {
		var numero int // dichiarazione della variabile in ingresso
		contatore++    // incremento di uno del contatore dei tentativi

		fmt.Printf("Tentativo n° %d: ", contatore) // messaggio a terminale per l'inserimento del tentativo
		fmt.Scan(&numero)                          // richiesta di inserimento di un numero intero

		if numero == numeroGenerato { // Hai indovinato il numero
			fmt.Printf("Hai indovinato in %d tentativi!\n", contatore) // messaggio a terminale di conferma
			break                                                      // interrompo il ciclo infinito
		} else if numero < numeroGenerato { // Il numero è troppo basso
			fmt.Println("Troppo basso! Riprova!") // messaggio a terminale offre un indizio
		} else { // Il numero è troppo alto
			fmt.Println("Troppo alto! Riprova!") // messaggio a terminale offre un indizio
		}
	}
}
