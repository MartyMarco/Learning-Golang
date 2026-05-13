# 28 - Quadrato di Asterischi

## Testo dell'esercizio
Scrivere un programma che legga da standard input un numero intero n e stampi a video un quadrato di n asterischi intervallati da spazi.

## Esecuzione

```bash
go run quadrato_asterischi.go
```

## Esempio di output

```
Inserisci un numero: 3
* * *
* * *
* * *
```

## Concetti utilizzati

- Due cicli `for` annidati (uno per le righe, uno per le colonne)
- Gestione dello spazio tra elementi senza spazio finale di riga
- `fmt.Println()` senza argomenti per andare a capo — forma idiomatica in Go
- Validazione dell'input con `return` anticipato

## Note personali

> Primo esercizio con **cicli annidati** — il ciclo esterno gestisce le righe, quello interno le colonne. Il numero totale di iterazioni è n².
>
> Ho corretto lo spazio finale di riga: `fmt.Print("* ")` applicato a ogni elemento incluso l'ultimo lasciava uno spazio invisibile ma tecnicamente scorretto. Ho usato `if j < dimensione-1` per stampare lo spazio solo tra gli elementi, non dopo l'ultimo — stesso pattern già visto negli esercizi precedenti.
>
> Ho sostituito `fmt.Print("\n")` con `fmt.Println()` senza argomenti — è la forma idiomatica in Go per stampare solo un newline.
>
> **Nota per il futuro:** il ciclo interno potrebbe essere sostituito da `strings.Repeat("* ", n)` che costruisce l'intera riga in memoria e la stampa in una chiamata sola invece di n chiamate separate — meno I/O e più efficiente. Da esplorare quando si imparerà il package `strings`.
