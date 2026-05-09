# 07 - Conversione Secondi in Ore, Minuti e Secondi

## Testo dell'esercizio
Scrivere un programma che prenda in ingresso un valore intero rappresentante una quantità di tempo in secondi e restituisca il numero di ore, minuti e secondi corrispondenti.

## Esecuzione

```bash
go run secondi_minuti_ore.go
```

## Esempio di output

```
Inserisci il numero di secondi : 15634
h:m:sec - 4:20:34 
```

```
Inserisci il numero di secondi : 134
h:m:sec - 0:2:14 
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e il verbo `%d`
- Operatore di divisione intera `/`
- Operatore modulo `%` per isolare le parti del tempo

## Note personali

> Ho trovato due approcci per risolvere il problema e li ho entrambi commentati nel codice.
> Il primo (**GOOD**) usa sottrazioni manuali con variabili intermedie — funziona ma è più verboso.
> Il secondo (**BETTER**) usa l'operatore `%` per isolare direttamente ogni componente del tempo senza variabili intermedie inutili — è più idiomatico e leggibile.
> Prima volta che uso `fmt.Printf()` con i verbi di formattazione: `%d` per interi e `\n` per andare a capo — più controllo sull'output rispetto a `fmt.Println()`.
