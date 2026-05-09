# 01 - Rettangolo

## Testo dell'esercizio
Scrivere un programma che legga da standard input le misure dell'altezza e della base di un rettangolo e ne calcoli il perimetro e l'area.

## Esecuzione

```bash
go run rettangolo.go
```

## Esempio di output

```
Inserisci due valori interi, uno per l'altezza e uno per la base del rettangolo :
20 10
Il valore del perimetro del rettangolo è  60
Il valore dell'area del rettangolo è  200
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Println()`
- Operazioni aritmetiche di base

## Note personali

> Ho scelto di richiedere i due valori con un unico `fmt.Scan()` invece di due prompt separati.
> Le variabili `Area` e `Perimetro` sono dichiarate con la prima lettera maiuscola: in Go questa è la convenzione per i nomi pubblici (esportati da un package). All'interno di `main` è più idiomatico usarle in minuscolo — da tenere a mente nei prossimi esercizi.
