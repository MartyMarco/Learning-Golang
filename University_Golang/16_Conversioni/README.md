# 16 - Convertitore di Tempo

## Testo dell'esercizio
Scrivere un programma che legga da standard input un valore intero che specifica il tipo di conversione da effettuare tra le seguenti opzioni:

1. secondi → ore
2. secondi → minuti
3. minuti → ore
4. minuti → secondi
5. ore → secondi
6. ore → minuti
7. minuti → giorni e ore
8. minuti → anni e giorni

Il programma legge poi il valore da convertire e stampa il risultato. Gestisce opportunamente una scelta non compresa tra 1 e 8.

## Esecuzione

```bash
go run conversioni.go
```

## Esempio di output

```
Scegli la conversione:
1) secondi -> ore
...
: 8
Inserisci il valore da convertire: 7200
7200 minuti corrispondono a 0 anni e 5 giorni
```

```
Scegli la conversione:
...
: 1
Inserisci il valore da convertire: 3618
3618 secondi corrispondono a 1.005 ore
```

```
Scegli la conversione:
...
: 9
Scelta errata
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e i verbi `%d`, `%g` e `\n`
- Menu interattivo a terminale con `fmt.Println()`
- Struttura condizionale `if / else if` con casi multipli
- Conversione di tipo con `float64()` per risultati decimali
- Operatore di divisione intera `/` e operatore modulo `%` per scomporre il tempo
- Validazione dell'input prima di procedere

## Note personali

> Esercizio più complesso dei precedenti — primo caso in cui la struttura condizionale gestisce 8 casi distinti più un caso di errore.
> Ho validato la scelta **prima** di chiedere il valore da convertire, seguendo il principio "valida prima, calcola dopo" già visto nell'esercizio sugli angoli del triangolo.
> Per le opzioni 7 e 8 ho documentato nel codice l'approccio sbagliato con ❌ — è un errore sottile: bisogna prima convertire tutto nell'unità intermedia (ore totali, giorni totali) e poi scomporre con `/` e `%`. Saltare questo passaggio produce risultati errati silenziosamente.
> Le opzioni che producono risultati interi (4, 5, 6) non richiedono `float64()`, mentre quelle che producono decimali (1, 2, 3) sì — ho applicato la conversione solo dove serve.
