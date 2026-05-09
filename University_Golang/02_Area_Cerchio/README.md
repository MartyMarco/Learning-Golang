# 02 - Cerchio

## Testo dell'esercizio
Scrivere un programma che legga da standard input il raggio di un cerchio e ne calcoli circonferenza e area.

## Esecuzione

```bash
go run cerchio.go
```

## Esempio di output

```
Inserisci il raggio del cerchio : 2.5
Circonferenza :  15.707963267948966
Area :  19.634954084936208
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Uso di `float64` per numeri decimali
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Print()` e `fmt.Println()`
- Importazione e utilizzo del package `math`
- Utilizzo della costante `math.Pi`

## Note personali

> Rispetto all'esercizio precedente ho usato `float64` invece di `int`, necessario per gestire valori decimali come il raggio e il pi greco.
> Ho usato `fmt.Print()` per il prompt di input (senza andare a capo) e `fmt.Println()` per i risultati — così l'inserimento del valore avviene sulla stessa riga del messaggio, più pulito visivamente.
> Ho corretto rispetto all'esercizio precedente: le variabili `circonferenza` e `area` sono in minuscolo, seguendo la convenzione idiomatica di Go.
