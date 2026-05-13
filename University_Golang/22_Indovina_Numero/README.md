# 22 - Indovina il Numero

## Testo dell'esercizio
Scrivere un programma che genera casualmente un numero intero tra 1 e 100 (estremi inclusi) e chiede iterativamente all'utente di indovinarlo. Ad ogni tentativo il programma indica se il numero inserito è troppo alto o troppo basso. Il programma termina quando l'utente indovina.

## Esecuzione

```bash
go run indovina_numero.go
```

## Esempio di output

```
Tentativo n° 1: 50
Troppo basso! Riprova!
Tentativo n° 2: 75
Troppo alto! Riprova!
Tentativo n° 3: 63
Troppo basso! Riprova!
Tentativo n° 4: 68
Troppo alto! Riprova!
Tentativo n° 5: 66
Hai indovinato in 5 tentativi!
```

## Concetti utilizzati

- Ciclo `for` infinito con `break` per uscita condizionale
- Generazione di numeri casuali con `math/rand` e `time`
- Inizializzazione moderna del generatore con `rand.NewSource` e `time.Now().UnixNano()`
- Struttura condizionale `if / else if / else` dentro un ciclo
- Contatore di tentativi incrementato ad ogni iterazione
- Output formattato con `fmt.Printf()` e `fmt.Println()`

## Note personali

> Primo esercizio con generazione di numeri casuali — ho dovuto importare due package nuovi: `math/rand` per la generazione e `time` per il seed.
>
> Ho trovato due approcci.
>
> Il primo (**❌ PRIMA**) aveva due problemi: `rand.Seed()` è deprecato da Go 1.20, e `rand.Intn(100)` genera un numero tra 0 e 99, non tra 1 e 100 come richiesto dalla consegna — un bug logico silenzioso che non causa errori ma produce un risultato sbagliato.
>
> Il secondo (**✅ DOPO**) usa `rand.New(rand.NewSource(...))` che è lo stile moderno e idiomatico di Go, e aggiunge `+1` al risultato di `Intn(100)` per ottenere l'intervallo corretto 1–100.
>
> Il contatore viene incrementato **prima** di leggere il numero — così il messaggio `Tentativo n° X` mostra già il numero corretto del tentativo corrente.
