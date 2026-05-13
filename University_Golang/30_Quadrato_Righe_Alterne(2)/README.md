# 30 - Quadrato con Righe Alterne (2)

## Testo dell'esercizio
Scrivere un programma che legga da standard input un numero intero n e stampi un quadrato di n righe alternando ciclicamente righe di `*`, `+` e `o`, intervallati da spazi.

## Esecuzione

```bash
go run quadrato_righe_alterne2.go
```

## Esempio di output

```
Inserisci un numero: 5
* * * * *
+ + + + +
o o o o o
* * * * *
+ + + + +
```

## Concetti utilizzati

- Due cicli `for` annidati
- Operatore modulo `%3` per alternare ciclicamente tra tre simboli
- Variabile `simbolo` per separare la scelta del carattere dalla logica di stampa
- `fmt.Println()` senza argomenti per il newline finale di riga
- Validazione dell'input con `return` anticipato

## Note personali

> Evoluzione dell'esercizio precedente — da due simboli alternati con `%2` a tre simboli in ciclo con `%3`. La struttura rimane identica, cambia solo la logica di selezione del simbolo.
>
> Il primo approccio (**❌ SBAGLIATO**) usava `%2` e `%3` insieme — sbagliato perché le condizioni si sovrappongono: ad esempio `i=6` è divisibile sia per 2 che per 3, producendo risultati inconsistenti.
>
> Il secondo approccio (**✅ CORRETTO**) usa solo `%3` con i tre possibili resti (0, 1, 2) che mappano esattamente ai tre simboli — ogni riga ha un solo simbolo possibile e il ciclo è garantito.
>
> **Nota per il futuro:** costruire la riga con `strings.Repeat` e `strings.Join` invece di chiamare `fmt.Print` per ogni carattere riduce le chiamate I/O da n a 1 per riga — da esplorare quando si imparerà il package `strings`.
