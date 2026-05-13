# 21 - Media

## Testo dell'esercizio
Scrivere un programma che legga da standard input una sequenza di numeri reali strettamente positivi. La lettura termina quando viene letto un numero minore o uguale a 0. Il programma deve stampare la media aritmetica dei valori inseriti.

## Esecuzione

```bash
go run medie.go
```

## Esempio di output

```
Inserisci una sequenza di numeri (interrompi con numero<=0): 4 6 8 0
Media aritmetica: 6
```

```
Inserisci una sequenza di numeri (interrompi con numero<=0): 3 5 2 6 1 -1
Media aritmetica: 3.4
```

## Concetti utilizzati

- Dichiarazione di variabili con scope ragionato (fuori e dentro il ciclo)
- Input da terminale con `fmt.Scan()` dentro un ciclo
- Output formattato con `fmt.Printf()` e il verbo `%g`
- Ciclo `for` infinito senza condizione
- `break` per interrompere il ciclo quando la condizione di uscita è soddisfatta
- Variabile accumulatore `somma` e contatore `contatore`
- Validazione con `return` anticipato se non è stato inserito nessun valore valido
- Conversione di tipo `float64(contatore)` per la divisione decimale

## Note personali

> Primo ciclo `for` senza condizione — il ciclo gira all'infinito e si interrompe solo con `break` quando l'utente inserisce un numero <= 0. È il pattern giusto quando non si sa in anticipo quante iterazioni fare.
> Ho dichiarato `numero` dentro il ciclo perché esiste solo lì e si azzera ad ogni iterazione, mentre `somma` e `contatore` vivono fuori perché devono accumularsi tra un'iterazione e l'altra.
> Ho aggiunto il controllo `contatore == 0` prima di calcolare la media — senza di esso, se l'utente inserisce subito 0 come primo valore, il programma farebbe una divisione per zero.
