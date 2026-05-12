# 20 - Operazioni su N Valori Variabili

## Testo dell'esercizio
Scrivere un programma che, dopo aver letto da standard input un numero intero n, chiede all'utente di inserire n numeri interi e stampa:
- la somma degli n numeri letti
- il minimo tra i numeri letti
- il massimo tra i numeri letti
- il numero di interi strettamente positivi, strettamente negativi e nulli

## Esecuzione

```bash
go run 20_Operazioni_Valori_Variabili.go
```

## Esempio di output

```
9
Inserisci 9 numeri:
1 -2 3 -4 5 -6 7 -8 9
somma = 5
valore minimo = -8
valore massimo = 9
interi > 0 = 5
interi < 0 = 4
interi = 0 = 0
```

## Concetti utilizzati

- Dichiarazione di variabili con `var` con scope ridotto al minimo necessario
- Input da terminale con `fmt.Scan()` dentro e fuori dal ciclo
- Output formattato con `fmt.Printf()`
- Ciclo `for` con contatore
- Struttura condizionale `if / else if / else` dentro un ciclo
- Variabili accumulatore (`somma`) e contatori (`positivi`, `negativi`, `nulli`)
- Validazione dell'input con `return` per uscire anticipatamente
- Inizializzazione di minimo e massimo con il primo valore reale

## Note personali

> Esercizio più complesso — combina ciclo, condizionali, accumulatori e contatori tutti insieme.
>
> Ho trovato due approcci.
>
> Il primo (**GOOD**) funziona ma usa `if i == 0` dentro il ciclo per inizializzare minimo e massimo — è un "trucco" che dipende dall'indice, non dal valore logico, e aggiunge un controllo extra ad ogni iterazione anche quando non serve più.
>
> Il secondo (**BETTER**) legge il primo valore fuori dal ciclo e lo usa direttamente come valore iniziale di minimo e massimo — più pulito e senza controlli inutili dentro il ciclo.
>
> Ho applicato il principio dello **scope ridotto**: `numero` è dichiarato dentro il ciclo perché esiste solo lì. Dichiarare tutto in cima alla funzione anche quando non serve è un'abitudine da evitare.
>
> Ho aggiunto la validazione di `quanto <= 0` con `return` per uscire anticipatamente — primo uso di `return` fuori da una funzione con valore di ritorno, usato qui come interruzione anticipata del programma.
