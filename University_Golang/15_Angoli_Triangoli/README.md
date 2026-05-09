# 15 - Angoli del Triangolo

## Testo dell'esercizio
Scrivere un programma che legga da standard input le ampiezze di due angoli di un triangolo e stampi, se possibile, l'ampiezza del terzo angolo.

La somma degli angoli interni di un triangolo è sempre 180°.

## Esecuzione

```bash
go run angoli_triangolo.go
```

## Esempio di output

```
Inserire le ampiezze dei due angoli : 50 60
Ampiezza terzo angolo = 70°
```

```
Inserire le ampiezze dei due angoli : 150 70
I due angoli non appartengono ad un triangolo
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output formattato con `fmt.Printf()` e `fmt.Println()`
- Struttura condizionale `if / else if / else`
- Operatori di confronto e operatore logico `||`
- Validazione dell'input prima di calcolare il risultato

## Note personali

> Ho trovato due approcci.
> Il primo (**GOOD**) calcola subito il terzo angolo e poi controlla se è valido — funziona, ma esegue un calcolo inutile anche quando i dati in ingresso sono già sbagliati.
> Il secondo (**BETTER**) valida prima i dati: controlla che nessuno dei due angoli sia zero o negativo, poi che la loro somma non superi 180°, e solo alla fine calcola il terzo angolo. È buona abitudine non fare calcoli su dati che non sono ancora stati validati.
> Questo pattern — **valida prima, calcola dopo** — è un principio generale che si applica in qualsiasi programma, non solo qui.
