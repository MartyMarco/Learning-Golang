# 10 - Multiplo di 10

## Testo dell'esercizio
Scrivere un programma che legge da standard input un numero intero n e verifica se il numero è multiplo di 10.

## Esecuzione

```bash
go run multipli.go
```

## Esempio di output

```
Inserisci numero : 15
15 non è multiplo di 10
```

```
Inserisci numero : 20
20 è multiplo di 10
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e i verbi `%d` e `\n`
- Struttura condizionale `if / else`
- Operatore modulo `%` per verificare la divisibilità

## Note personali

> La verifica di divisibilità tramite `numero % 10 == 0` è il pattern classico per controllare se un numero è multiplo di un altro — funziona con qualsiasi divisore, non solo 10.
> Ho ricordato di aggiungere `\n` alla fine di `fmt.Printf()` — senza di esso il cursore del terminale rimane sulla stessa riga dell'output, comportamento indesiderato. Con `fmt.Println()` non serve perché va a capo automaticamente, ma con `fmt.Printf()` bisogna ricordarselo.
