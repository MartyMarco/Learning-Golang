/*

TESTO : Scrivere un programma che legga da standard input due numeri interi che chiameremo N e L e calcoli l’area di un poligono regolare con N lati di lunghezza L.

SUGGERIMENTO : l'area di un poligono regolare può essere calcolata utilizzando le funzioni math.Pow (per il calcolo della potenza) e math.Tan (per il calcolo della
tangente di un angolo) del package math applicando la seguente formula:
									area = (n*l^2)/4*tan(π/n)

ESEMPIO :
$ go run areapoligono.go
Inserisci il numero di lati del poligono: 6
Inserisci la lunghezza dei lati del poligono: 3
Area calcolata: 23.382685902179844

$ go run areapoligono.go
Inserisci il numero di lati del poligono: 4
Inserisci la lunghezza dei lati del poligono: 3
Area calcolata: 9

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var numero, lunghezza int // dichiarazione delle due variabili necessarie

	fmt.Print("Inserisci il numero di lati del poligono : ")     // messaggio output a terminale
	fmt.Scan(&numero)                                            // richiesta del primo dato dal terminale
	fmt.Print("Inserisci la lunghezza dei lati del poligono : ") // messaggio output a terminale
	fmt.Scan(&lunghezza)                                         // richiesta del secondo dato dal terminale

	area := float64(numero) * math.Pow(float64(lunghezza), 2) / (4 * math.Tan(math.Pi/float64(numero))) // calcolo dell'area del poligono tramite la formula
	fmt.Println("Area : ", area)                                                                        // stampa a terminale del valore dell'area

}

/*
IMPORTANTE : In Go la divisione ha la stessa precedenza della moltiplicazione e viene eseguita da sinistra a destra, quindi è meglio rendere l'intenzione esplicita
con le parentesi.
*/
