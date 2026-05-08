/*
TESTO : Scrivere un programma che legga da standard input le misure dell’altezza e della base di un rettangolo e ne calcoli il perimetro e l’area.
ESEMPIO :
$ go run rettangolo.go
Inserisci la base: 20
Inserisci l'altezza: 10
Perimetro = 60
Area = 200
*/

package main

import (
	"fmt"
)

func main() {
	var altezza, base int // Dichiarazione delle variabili

	fmt.Println("Inserisci due valori interi, uno per l'altezza e uno per la base del rettangolo :") // messaggio in output
	fmt.Scan(&altezza, &base)                                                                        // richiesta di dati in input

	Area := base * altezza            // Calcolo dell'area del rettangolo
	Perimetro := 2 * (base + altezza) // Calcolo del perimetro del rettangolo

	fmt.Println("Il valore del perimetro del rettangolo è ", Perimetro) // Restituzione del valore del perimetro al terminale
	fmt.Println("Il valore dell'area del rettangolo è ", Area)          // Restituzione del valore dell'area al terminale
}
