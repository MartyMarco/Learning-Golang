# 12 - Fizz Buzz

## Testo dell'esercizio
Scrivere un programma che legge da standard input un numero intero e stampa:
- **Fizz** se il numero è multiplo di 3
- **Buzz** se il numero è multiplo di 5
- **Fizz Buzz** se è multiplo sia di 3 sia di 5
- niente altrimenti

## Esecuzione

```bash
go run fizz_buzz.go
```

## Esempio di output

```
Inserisci un numero : 5
Buzz
```

```
Inserisci un numero : 4

```

```
Inserisci un numero : 15
Fizz Buzz
```

```
Inserisci un numero : 6
Fizz
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Println()`
- Struttura condizionale `if / else if / else`
- Operatore modulo `%` per verificare la divisibilità
- Operatore logico `&&`

## Note personali

> FizzBuzz è un esercizio classico.
> Ho trovato due approcci.
> Il primo (**GOOD**) usa `numero%3 == 0 && numero%5 == 0` per il caso combinato — corretto, ma più verboso del necessario.
> Il secondo (**BETTER**) sostituisce la doppia condizione con `numero%15 == 0` — un numero è multiplo sia di 3 che di 5 se e solo se è multiplo del loro minimo comune multiplo, che è 15. Più conciso e idiomatico.
> Ho aggiunto un `else` esplicito vuoto per il caso in cui il numero non sia multiplo né di 3 né di 5 — anche se non fa nulla, rende chiaro che il caso è stato considerato e non dimenticato.
