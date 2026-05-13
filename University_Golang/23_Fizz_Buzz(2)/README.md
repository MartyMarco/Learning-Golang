# 23 - FizzBuzz (2) Sequenza

## Testo dell'esercizio
Scrivere un programma che legga da standard input un intero n e stampi la sequenza da 1 a n sostituendo:
- ogni multiplo di 3 con **Fizz**
- ogni multiplo di 5 con **Buzz**
- ogni multiplo sia di 3 che di 5 con **FizzBuzz**

## Esecuzione

```bash
go run fizz_buzz_2.go
```

## Esempio di output

```
Inserisci un numero intero: 6
1 2 Fizz 4 Buzz Fizz
```

```
Inserisci un numero intero: 20
1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz
```

## Concetti utilizzati

- Ciclo `for` con contatore da 1 a n
- Struttura condizionale `if / else if / else` dentro un ciclo
- Operatore modulo `%` per la divisibilità
- `fmt.Print()` senza `\n` per stampare tutto sulla stessa riga
- Gestione della formattazione: spazio tra elementi ma non dopo l'ultimo
- Validazione dell'input con `return` anticipato

## Note personali

> Variante del FizzBuzz dell'esercizio 12 — lì si testava un singolo numero, qui si stampa una sequenza intera da 1 a n.
> Ho riapplicato l'ottimizzazione già vista: `i%15 == 0` invece del doppio `&&` per il caso combinato, e il controllo del caso multiplo va sempre per primo.
> Ho gestito la formattazione dell'output con un `if i < numero` per stampare lo spazio separatore solo tra gli elementi e non dopo l'ultimo — dettaglio piccolo ma che fa la differenza visivamente.
> Ho aggiunto la validazione di `numero <= 0` con `return` anticipato — senza di essa il ciclo non eseguirebbe nulla e il programma terminerebbe silenziosamente senza spiegazioni.
