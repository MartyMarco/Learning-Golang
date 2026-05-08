/*

TESTO : Scrivere un programma che legga da standard input le età di due persone (espresse in anni) e calcoli:
			-)	la somma delle età inserite;
			-)	la media delle età inserite;
			-)	la media delle età inserite arrotondata per difetto all'intero inferiore;
			-)	la media delle età inserite arrotondata per eccesso all'intero superiore;
			-)	la somma e la media delle età che le due persone avranno fra 10 anni.

SUGGERIMENTO : la media arrotondata per difetto può essere calcolata usando la funzione math.Floor del package math nel seguente modo:
					var mediaArrotondataDifetto float64 = math.Floor(media)
			Similarmente, la media arrotondata per eccesso può essere calcolata usando la funzione math.Ceil nel seguente modo:
					var mediaArrotondataEccesso float64 = math.Ceil(media)

ESEMPIO :
$ go run calcoloeta.go
Età persona 1 = 15
Età persona 2 = 20
Somma età = 35
Media età = 17.5
Media età arrotondata per difetto all'intero inferiore = 17
Media età arrotondata per eccesso all'intero superiore = 18
Somma età a 10 anni = 55
Media età a 10 anni = 27.5

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var eta1, eta2 int // Dichiarazione delle variabili che indicano l'età delle due persone

	fmt.Print("Inserisci l'età di due persone diverse : ") // messaggio in output al terminale
	fmt.Scan(&eta1, &eta2)                                 // richiesta dei dati delle due età al terminale

	Somma := eta1 + eta2            // calcolo della somma delle due età (intero)
	Media := float64(Somma) / 2     // calcolo della media delle due età (float)
	Difetto := math.Floor(Media)    // calcolo arrotondato della media delle due età per difetto (float64)
	Eccesso := math.Ceil(Media)     // calcolo arrotondato della media delle due età per eccesso (float64)
	Somma10 := Somma + 20           // calcolo della somma delle due età dopo 10 anni (intero)
	Media10 := float64(Somma10) / 2 // calcolo della media delle due età dopo 10 anni (float)

	fmt.Println("Somma delle due età : ", Somma)                // restituzione del valore Somma a terminale
	fmt.Println("Media delle due età : ", Media)                // restituzione del valore Media a terminale
	fmt.Println("Media arrotondata per Difetto : ", Difetto)    // restituzione del valore Difetto a terminale
	fmt.Println("Media arrotondata per Eccesso : ", Eccesso)    // restituzione del valore Eccesso a terminale
	fmt.Println("Somma delle due età dopo 10 anni : ", Somma10) // restituzione del valore Somma10 a terminale
	fmt.Println("Media delle due età dopo 10 anni : ", Media10) // restituzione del valore Media10 a terminale
}
