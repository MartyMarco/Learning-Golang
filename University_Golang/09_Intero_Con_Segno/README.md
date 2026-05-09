# 09 - Intero con Segno

## Testo dell'esercizio
Scrivere un programma che legge da standard input un numero intero n (specificato senza segno se maggiore o uguale a 0) e stampi a video il numero con segno.

## Esecuzione

```bash
go run intero_segno.go
```

## Esempio di output

```
Inserisci numero : 5
+5
```

```
Inserisci numero : 0
0
```

```
Inserisci numero : -5
-5
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e il verbo `%+d`
- Prima struttura condizionale `if / else`

## Note personali

> Prima volta che uso una struttura condizionale `if / else` — lo strumento fondamentale per gestire casi diversi.
> Ho trovato due approcci anche qui.
> Il primo (**GOOD**) gestisce i tre casi manualmente con `if / else if / else` — funziona ma contiene una condizione ridondante: se il numero non è negativo e non è zero, è necessariamente positivo, quindi `else if numero > 0` può essere semplicemente `else`.
> Il secondo (**BETTER**) sfrutta il verbo `%+d` di `fmt.Printf()` che aggiunge automaticamente il segno `+` ai positivi e `-` ai negativi — è necessario gestire separatamente solo lo zero perché `%+d` lo stamperebbe come `+0`.
> Lezione importante: prima di implementare logica manuale vale sempre la pena chiedersi se la libreria standard offre già quello che serve.
