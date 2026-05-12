# 17 - Punto e Retta

## Testo dell'esercizio
Scrivere un programma che legga da standard input il coefficiente angolare m e il termine noto q di una retta `y = m*x + q`, e le coordinate px e py di un punto P(px, py). Il programma deve determinare se il punto sta sopra, sotto o appartiene alla retta.

Un punto appartiene alla retta se sostituendo la sua ascissa nell'equazione si ottiene esattamente la sua ordinata. Sta sopra se py > y calcolata, sotto se py < y calcolata.

## Esecuzione

```bash
go run retta.go
```

## Esempio di output

```
Inserisci m e q : 1 0
Inserisci x e y : 5 5
Il punto appartiene alla retta
```

```
Inserisci m e q : 1 1
Inserisci x e y : 5 5
Il punto sta sotto alla retta
```

## Concetti utilizzati

- Dichiarazione di variabili `float64` con `var`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Println()`
- Struttura condizionale `if / else if / else`
- Operatori di confronto `==`, `>`, `<`
- Applicazione di una formula matematica (equazione della retta)

## Note personali

> Primo esercizio con un contesto geometrico — ho dovuto ragionare sull'equazione della retta e su cosa significa che un punto "appartiene" ad essa.
> Ho aggiunto un ramo `else` finale come caso di errore anche se matematicamente irraggiungibile (se y non è uguale, maggiore o minore di py non esiste un quarto caso) — rende esplicito che tutti i casi sono stati considerati.
> **Attenzione importante:** confrontare due `float64` con `==` è generalmente pericoloso — operazioni aritmetiche su numeri decimali possono introdurre piccoli errori di precisione che rendono due valori "uguali" mai esattamente identici. Il modo corretto sarebbe confrontare la differenza con una tolleranza (es. `math.Abs(y - py) < 0.0001`). In questo esercizio il confronto diretto funziona perché i valori sono semplici, ma è un'abitudine da evitare in contesti reali.
