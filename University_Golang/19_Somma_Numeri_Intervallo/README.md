# 19 - Somma dei Pari in un Intervallo

## Testo dell'esercizio
Scrivere un programma che legga da standard input due numeri interi a e b e stampi a video la somma dei numeri pari compresi tra a e b (estremi esclusi).

## Esecuzione

```bash
go run sommaintervallo.go
```

## Esempio di output

```
Inserisci due numeri : 1 9
Somma = 20
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()`
- Ciclo `for` con condizione e incremento personalizzato (`i += 2`)
- Struttura condizionale `if` per validare l'ordine degli estremi
- Swap idiomatico di due variabili in Go: `a, b = b, a`
- Operatore modulo `%` per verificare la parità
- Variabile accumulatore `somma`

## Note personali

> Ho trovato tre approcci con miglioramenti progressivi.
>
> Il primo (**NOT-SO-GOOD**) duplica lo stesso codice del ciclo due volte cambiando solo l'ordine delle variabili — funziona ma viola il principio DRY (*Don't Repeat Yourself*): se devo modificare la logica, devo farlo in due posti.
>
> Il secondo (**GOOD**) risolve con lo swap idiomatico di Go `a, b = b, a` — in altri linguaggi serviva una variabile temporanea, in Go si fa in una riga sola. Dopo lo swap il ciclo è scritto una volta sola.
>
> Il terzo (**BETTER**) ottimizza il ciclo stesso: invece di controllare ogni numero con `i%2 == 0`, trova direttamente il primo pari nell'intervallo e poi avanza di 2 in 2 con `i += 2`. Così si dimezzano le iterazioni e si elimina il controllo interno.
