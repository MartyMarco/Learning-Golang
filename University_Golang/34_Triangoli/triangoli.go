/*
-----------------------------------------------------------------------------------------------------------------------------------------------------
	TESTO :	Scrivere un programma che legga da standard input un intero n > 1 e stampi, utilizzando il carattere *, il perimetro di
			due triangoli rettangoli con base e altezza di lunghezza n, posizionati come mostrato nell'Esempio d'esecuzione.

	ESEMPIO :
		$ go run triangoli.go
		2
		**
		 *
		 *
		 **

		$ go run triangoli.go
		3
		***
		 **
		  *
		  *
		  **
		  ***

		$ go run triangoli.go
		4
		****
		 * *
		  **
		   *
		   *
		   **
		   * *
		   ****

		$ go run triangoli.go 6
		******
		 *   *
		  *  *
		   * *
		    **
		     *
		     *
		     **
		     * *
		     *  *
		     *   *
		     ******
-----------------------------------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var numero int                      // dichiarazione della variabile necessaria per la dimensione dei triangoli
	fmt.Print("Inserisci dimensione: ") // messaggio di output a livello di terminale
	fmt.Scan(&numero)                   // richiesta di inserimento di un numero intero

	if numero <= 1 { // controllo che l'inserimento sia avvenuto correttamente
		fmt.Println("Errore: il numero deve essere > 1")
		return
	}

	/*
		-----------------------------------------------------------------------------------------------------------------------------------------------------
		❌ BAD :
		dimensione := numero * 2 // consideriamo i due trinagoli racchiusi in un quadrato di lato (n*2)x(n*2)

		for i := 0; i < dimensione; i++ {

			for j := 0; j < dimensione; j++ {
				if i == 0 && j < numero {
					fmt.Print("*")
				} else if i == dimensione-1 && j < numero {
					fmt.Print(" ")
				} else if i == dimensione-1 && j > numero {
					fmt.Print("*")
				} else if i >= 1 && i <= numero/2+1 {
					fmt.Print(strings.Repeat(" ", i), "*", strings.Repeat(" ", numero-2-i), "*")
				}

				if i == dimensione/2 || i == dimensione/2+1 {
					fmt.Print(strings.Repeat(" ", numero-1), "*")
				}
			}

			fmt.Println()

		}
		-----------------------------------------------------------------------------------------------------------------------------------------------------
	*/

	// ✅ BETTER :

	// TRIANGOLO SUPERIORE: n righe, vertice in basso a destra
	for i := 0; i < numero; i++ {
		if i == 0 { // Prima riga: base intera
			fmt.Println(strings.Repeat("*", numero))
		} else if numero-2-i >= 0 { // Righe intermedie e ultima: spazi + "*" + spazi interni + "*"
			spazioSx := strings.Repeat(" ", i)
			spazioInterno := strings.Repeat(" ", numero-2-i)
			fmt.Println(spazioSx + "*" + spazioInterno + "*")
		}
	}

	if numero-2 >= 0 {
		fmt.Println(strings.Repeat(" ", numero-2), "*")
		fmt.Println(strings.Repeat(" ", numero-2), "*")
	}

	// TRIANGOLO INFERIORE: n righe, specchio verticale del superiore
	for i := numero - 1; i >= 0; i-- {
		if i == 0 { // Ultima riga: base intera
			fmt.Println(strings.Repeat(" ", numero-1) + strings.Repeat("*", numero))
		} else if numero-2-i >= 0 { // Righe intermedie e prima: spazi + "*" + spazi interni + "*"
			spazioSx := strings.Repeat(" ", numero-1)
			spazioInterno := strings.Repeat(" ", numero-2-i)
			fmt.Println(spazioSx + "*" + spazioInterno + "*")
		}
	}
}
