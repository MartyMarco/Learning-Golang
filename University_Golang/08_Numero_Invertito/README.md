# 08 - Numero Invertito

## Testo dell'esercizio
Scrivi un programma che inverta le cifre di un numero intero a tre cifre inserito da standard input.

Esempio: dato n=123 il programma dovrebbe restituire 321.

## Esecuzione

```bash
go run numero_invertito.go
```

## Esempio di output

```
Inserisci un numero da 3 cifre : 123
Risultato :  321
```

```
Inserisci un numero da 3 cifre : 453
Risultato :  354
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Println()`
- Operatore di divisione intera `/` per estrarre cifre
- Operatore modulo `%` per isolare cifre
- Ricomposizione di un numero tramite moltiplicazione per potenze di 10

## Note personali

> Ho trovato due approcci anche qui e li ho entrambi nel codice.
> Il primo (**GOOD**) usa nomi generici come `primo`, `mezzo`, `ultimo` — funziona ma i nomi non descrivono bene cosa rappresentano le variabili.
> Il secondo (**BETTER**) usa nomi semantici come `cifraUnita`, `cifraDecine`, `cifraCentinaia` — molto più leggibile e chiaro, anche a distanza di tempo.
> La logica per estrarre le cifre riprende l'operatore `%` dell'esercizio precedente — si inizia a vedere come gli stessi strumenti tornano in contesti diversi.
> Da correggere: `Invertito` dovrebbe essere in minuscolo seguendo la convenzione idiomatica di Go.
