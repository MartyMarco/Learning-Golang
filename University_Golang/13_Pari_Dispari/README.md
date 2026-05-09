# 13 - Pari o Dispari

## Testo dell'esercizio
Scrivere un programma che legge da standard input un intero n e stampa a video se il numero è pari o dispari.

## Esecuzione

```bash
go run pari_dispari.go
```

## Esempio di output

```
Inserisci un numero : 10
10 è pari
```

```
Inserisci un numero : 11
11 è dispari
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e i verbi `%d` e `\n`
- Struttura condizionale `if / else`
- Operatore modulo `%` per verificare la parità

## Note personali

> Il pattern `numero % 2 == 0` per verificare la parità è uno dei più usati in programmazione — ormai è diventato automatico dopo gli esercizi precedenti.
> Non servono condizioni ridondanti: se il numero non è pari è necessariamente dispari, quindi un semplice `else` è sufficiente senza aggiungere un'altra condizione esplicita.
