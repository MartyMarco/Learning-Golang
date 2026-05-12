# 18 - Tabellina

## Testo dell'esercizio
Scrivere un programma che, dopo aver richiesto all'utente di inserire un numero intero n, stampi a video la corrispondente tabellina moltiplicando n per i numeri naturali da 1 a 10.

## Esecuzione

```bash
go run tabellina.go
```

## Esempio di output

```
Inserisci un numero : 9
1 x 9 = 9
2 x 9 = 18
3 x 9 = 27
4 x 9 = 36
5 x 9 = 45
6 x 9 = 54
7 x 9 = 63
8 x 9 = 72
9 x 9 = 81
10 x 9 = 90
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e i verbi `%d` e `\n`
- Primo ciclo `for` con variabile contatore, condizione e incremento

## Note personali

> Prima volta che uso il ciclo `for` — la struttura è `for i := 1; i <= 10; i++` dove:
> - `i := 1` è l'inizializzazione del contatore
> - `i <= 10` è la condizione di continuazione
> - `i++` è l'incremento ad ogni iterazione
> Quello che prima avrebbe richiesto 10 righe di `fmt.Printf()` separate si risolve in 3 righe — il ciclo è lo strumento giusto ogni volta che si ripete la stessa operazione con un valore che cambia.
