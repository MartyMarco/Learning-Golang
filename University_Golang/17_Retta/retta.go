/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input 4 numeri all'interno di due variavili di tipo float64:
				-	i primi due valori sono il coefficiente angolare m e il termine noto q di una retta r: y = m*x + q
				-	il terzo e il quarto valore sono le coordinate px e py di un punto P(px,py)
			Il programma deve determinare se il punto P sta sopra o sotto la retta od appartiene ad essa, e stampare a video il relativo
			messaggio.

	SUGGERIMENTO : un punto appartiene ad una retta se sostituendo le sue coordinate nell'equazione della retta l'uguaglianza è verificata. Un
	punto sta sopra una retta se sostituendo il valore dell'ascissa nell'equazione della retta si ottiene y < py.

	ESEMPIO :
		$ go run retta.go
		Inserisci m e q: 1 0
		Inserisci x e y: 5 5
		Il punto appartiene alla retta

		$ go run retta.go
		Inserisci m e q: 1 1
		Inserisci x e y: 5 5
		Il punto sta sotto la retta
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var m, q, px, py float64 // dichiarazioni delle variabili necessarie

	fmt.Print("Inserisci m e q : ") // messaggio output a terminale
	fmt.Scan(&m, &q)                // richiesta dei due dati necessari a livello del terminale
	fmt.Print("Inserisci x e y : ") // messaggio output a terminale
	fmt.Scan(&px, &py)              // richiesta dei due dati necessari a livello del terminale

	y := (m * px) + q // calcolo del valore della funzione y

	if y == py { // il punto appartiene alla retta
		fmt.Println("Il punto appartiene alla retta")
	} else if py > y { // il punto sta sopra alla retta
		fmt.Println("Il punto sta sopra alla retta")
	} else if py < y { // il punto sta sotto alla retta
		fmt.Println("Il punto sta sotto alla retta")
	} else { // controllo degli errori
		fmt.Println("Errore")
	}
}

/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	ATTENZIONE	:	normalmente sarebbe pericoloso comparare ( float64 == float64 ) in questo modo perchè raramente sono esattamente uguali
					dopo i calcoli aritmetici, di conseguenza normalmente sabbe meglio confronta la differenza con una tolleranza.

					In questo caso, nel piano cartesiano un punto appartiene alla retta se la coordinata coincide perfettamente.
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/
