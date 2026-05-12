/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che, dopo aver richiesto all'utente di inserire da standard input un numero intero n, stampi a video
	 		la corrispondente tabellina (moltiplicando n per i numeri naturali da 1 a 10) come mostrato nell'Esempio d'esecuzione.

	ESEMPIO :
		$ go run tabelline.go
		Inserisci un numero: 9
		1 x 9 = 9
		2 x 9 = 18
		3 x 9 = 27
		4 x 9 = 36
		5 x 9 = 45
		6 x 9 = 54
		7 x 9 = 63
		8 x 9 = 72
		9 x 9 = 81
		10 x 9 = 90
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var numero int // dichiarazione della variabile necessaria

	fmt.Print("Inserisci un numero : ") // messaggio output a livello di terminale
	fmt.Scan(&numero)                   // richiesta del dato necessario

	for i := 1; i <= 10; i++ { // iterazione che parte da 1 e arriva fino a 10 (compreso)
		fmt.Printf("%d x %d = %d\n", i, numero, i*numero) // stampo a terminale la formattazione della tabellina
	}
}
