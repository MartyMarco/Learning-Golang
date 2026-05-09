# 14 - Divisione con Controllo dello Zero

## Testo dell'esercizio
Scrivere un programma che legga da standard input due numeri a e b di tipo `int` e calcoli il risultato della divisione a/b. Se b è uguale a 0, il programma stampa "Impossibile".

## Esecuzione

```bash
go run divisione.go
```

## Esempio di output

```
Inserisci due numeri : 
5 2
Quoziente = 2.5
```

```
Inserisci due numeri : 
5 0
Impossibile
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e il verbo `%g`
- Struttura condizionale `if / else`
- Conversione di tipo con `float64()` per divisione decimale
- Gestione di un caso di errore (divisione per zero)

## Note personali

> La divisione per zero è un errore classico che in Go causa un panic a runtime se non gestita — controllarla prima con `if divisore != 0` è la soluzione corretta.
> Ho usato `float64(dividendo) / float64(divisore)` per ottenere un risultato decimale — senza la conversione la divisione intera troncherebbe i decimali (es. 5/2 = 2 invece di 2.5).
> Ho scelto `fmt.Printf()` con il verbo `%g` invece di `fmt.Println()` per due motivi: evita lo spazio doppio che `fmt.Println()` inserisce automaticamente tra gli argomenti, e `%g` rimuove gli zeri decimali inutili (es. stampa `2.5` invece di `2.500000`).
