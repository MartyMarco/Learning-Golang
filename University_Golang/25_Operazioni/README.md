# 25 - Operazioni tra Due Numeri

## Testo dell'esercizio
Scrivere un programma che legga da standard input due numeri interi x e y e stampi:
- il maggiore e il minore tra i due
- la somma, la differenza (maggiore - minore), il prodotto e la divisione
- il valore medio
- x elevato a y calcolato sia con un ciclo `for` sia con `math.Pow`

## Esecuzione

```bash
go run operazioni.go
```

## Esempio di output

```
Inserisci due numeri interi:
2 4
Maggiore: 4
Minore: 2
Somma: 6
Differenza: 2
Prodotto: 8
Divisione: 0.5
Valore medio: 3
Potenza (ciclo for): 16
Potenza (math.Pow): 16
```

## Concetti utilizzati

- Dichiarazione di variabili con `var` e `:=`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e i verbi `%d`, `%g`, `\n`
- Struttura condizionale `if / else` per determinare maggiore e minore
- Conversione di tipo `float64()` per la divisione decimale
- Ciclo `for` per il calcolo della potenza senza librerie
- `math.Pow` per il calcolo della potenza con la libreria standard
- Gestione del caso limite divisione per zero

## Note personali

> Esercizio riepilogativo che mette insieme molti concetti già visti — operazioni aritmetiche, conversioni di tipo, condizionali e cicli.
>
> Ho calcolato la potenza in due modi distinti come richiesto: con un ciclo `for` che moltiplica la base y volte, e con `math.Pow`. È utile conoscere entrambi — il ciclo per capire la logica, `math.Pow` per il codice reale.
>
> Ho corretto il ciclo della potenza partendo da `i := 0` ed eseguendo esattamente y iterazioni invece di partire da 1 e fermarsi a y-1 — il risultato è lo stesso ma il ragionamento è più diretto.
>
> Ho aggiunto il controllo della divisione per zero prima di eseguirla — dividere per zero in Go causa un crash a runtime per gli interi o produce `+Inf` per i float, entrambi comportamenti indesiderati.
>
> **Note generali documentate nel codice:**
> - Separare calcoli e stampe rende il codice più leggibile e facile da modificare
> - Dichiarare le variabili con `:=` il più vicino possibile al loro primo utilizzo è lo stile idiomatico di Go
