# 06 - Divisione con Resto

## Testo dell'esercizio
Scrivere un programma che prenda in ingresso due numeri e stampi il quoziente ed il resto della divisione del primo numero per il secondo.

## Esecuzione

```bash
go run divisione_resto.go
```

## Esempio di output

```
Inserisci due numeri : 5 3
Quoziente :  1
Resto :  2
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Print()` e `fmt.Println()`
- Operatore di divisione intera `/`
- Operatore modulo `%` per il calcolo del resto

## Note personali

> Esercizio semplice ma introduce l'operatore `%` (modulo) — uno degli operatori più utili in programmazione, usato spessissimo per verificare se un numero è pari/dispari, per i cicli, per le operazioni su cifre di un numero e molto altro.
> La divisione tra due `int` in Go restituisce automaticamente la parte intera senza bisogno di conversioni — al contrario degli esercizi precedenti dove serviva `float64()`.
