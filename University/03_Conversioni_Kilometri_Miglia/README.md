# 03 - Convertitore Km → Miglia

## Testo dell'esercizio
Scrivere un programma che legga da standard input una distanza in Km ed effettui la conversione di tale distanza in miglia (1 Km = 0.62137 mi).

## Esecuzione

```bash
go run convertitore.go
```

## Esempio di output

```
Inserisci il numero di kilometri da convertire : 12
Miglia :  7.45644
```

## Concetti utilizzati

- Dichiarazione di una costante con `const` a livello di package
- Dichiarazione di variabili con `var`
- Short variable declaration con `:=`
- Uso di `float64` per numeri decimali
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Print()` e `fmt.Println()`

## Note personali

> Prima volta che uso `const` fuori da `main()`, a livello di package — la costante è quindi visibile da tutto il file.
> Rispetto a usare un numero magico direttamente nel calcolo (`kilometri * 0.62137`), dichiarare la costante con un nome descrittivo rende il codice più leggibile e facile da modificare in futuro.
