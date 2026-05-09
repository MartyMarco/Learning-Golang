# 05 - Area Poligono Regolare

## Testo dell'esercizio
Scrivere un programma che legga da standard input due numeri interi che chiameremo N e L e calcoli l'area di un poligono regolare con N lati di lunghezza L.

La formula utilizzata è:
```
area = (n * l²) / (4 * tan(π/n))
```

## Esecuzione

```bash
go run area_poligono_regolare.go
```

## Esempio di output

```
Inserisci il numero di lati del poligono : 6
Inserisci la lunghezza dei lati del poligono : 3
Area :  23.382685902179844
```

```
Inserisci il numero di lati del poligono : 4
Inserisci la lunghezza dei lati del poligono : 3
Area :  9
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Conversione di tipo con `float64()`
- Input da terminale con `fmt.Scan()` separati per ogni variabile
- Output a terminale con `fmt.Print()` e `fmt.Println()`
- Importazione e utilizzo del package `math`
- Utilizzo di `math.Pow()` per il calcolo della potenza
- Utilizzo di `math.Tan()` per il calcolo della tangente
- Utilizzo della costante `math.Pi`

## Note personali

> Ho dovuto convertire più volte `int` in `float64` nella stessa espressione — `float64(numero)` e `float64(lunghezza)` — perché `math.Pow()` e `math.Tan()` accettano solo `float64`.
> **Attenzione alla precedenza degli operatori in Go:** la moltiplicazione e la divisione hanno la stessa precedenza e vengono eseguite da sinistra a destra. Senza le parentesi giuste il calcolo sarebbe sbagliato — ho reso l'intenzione esplicita con le parentesi per sicurezza.
> Ho scelto due `fmt.Scan()` separati invece di uno solo per rendere più chiaro quale valore si sta inserendo in ogni momento.
