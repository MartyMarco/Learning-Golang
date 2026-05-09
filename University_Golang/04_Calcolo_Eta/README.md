# 04 - Calcolo Età

## Testo dell'esercizio
Scrivere un programma che legga da standard input le età di due persone (espresse in anni) e calcoli:
- la somma delle età inserite
- la media delle età inserite
- la media delle età inserite arrotondata per difetto all'intero inferiore
- la media delle età inserite arrotondata per eccesso all'intero superiore
- la somma e la media delle età che le due persone avranno fra 10 anni

## Esecuzione

```bash
go run calcolo_eta.go
```

## Esempio di output

```
Inserisci l'età di due persone diverse : 15 20
Somma delle due età :  35
Media delle due età :  17.5
Media arrotondata per Difetto :  17
Media arrotondata per Eccesso :  18
Somma delle due età dopo 10 anni :  55
Media delle due età dopo 10 anni :  27.5
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Conversione di tipo con `float64()`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Print()` e `fmt.Println()`
- Importazione e utilizzo del package `math`
- Utilizzo di `math.Floor()` per arrotondamento per difetto
- Utilizzo di `math.Ceil()` per arrotondamento per eccesso

## Note personali

> Per calcolare la media ho dovuto convertire `Somma` (che è `int`) in `float64` con `float64(Somma)` — senza questa conversione la divisione intera avrebbe troncato i decimali e il risultato sarebbe stato sbagliato (es. 35/2 = 17 invece di 17.5).
> Per le età fra 10 anni ho sommato 20 alla somma totale invece di aggiungere 10 a ciascuna età separatamente — il risultato è lo stesso ma con un calcolo solo.
> Da correggere: le variabili `Somma`, `Media`, `Difetto`, `Eccesso`, `Somma10`, `Media10` andrebbero in minuscolo seguendo la convenzione idiomatica di Go.
