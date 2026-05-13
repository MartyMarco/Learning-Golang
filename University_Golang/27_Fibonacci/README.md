# 27 - Successione di Fibonacci

## Testo dell'esercizio
Scrivere un programma che legga da standard input un intero positivo n e stampi i primi n elementi della successione di Fibonacci, in cui ciascun numero è la somma dei due precedenti.

## Esecuzione

```bash
go run fibonacci.go
```

## Esempio di output

```
Inserisci un numero intero:
4
1 1 2 3
```

```
Inserisci un numero intero:
12
1 1 2 3 5 8 13 21 34 55 89 144
```

## Concetti utilizzati

- Ciclo `for` con contatore
- Due variabili accumulatore `a` e `b` per tenere traccia dei due termini precedenti
- Swap e aggiornamento simultaneo con `a, b = b, a+b`
- Output formattato con `fmt.Printf()` gestendo lo spazio e il newline finale

## Note personali

> La successione di Fibonacci è un classico dell'informatica — la logica è semplice ma richiede di ragionare su come aggiornare i due termini precedenti ad ogni passo.
>
> Il mio primo approccio (**❌ ERRATO**) cercava di calcolare ogni termine direttamente dall'indice con `i + (i-1)` — sbagliato perché Fibonacci non è una formula sull'indice ma una relazione tra termini consecutivi.
>
> Il secondo approccio (**✅ GIUSTO**) usa due accumulatori `a` e `b` inizializzati a 0 e 1. Ad ogni iterazione `a, b = b, a+b` aggiorna entrambi simultaneamente: `a` diventa il termine corrente da stampare, `b` il prossimo. Lo swap idiomatico di Go in una riga sola evita la variabile temporanea.
>
> Ho gestito la formattazione stampando lo spazio dopo ogni elemento tranne l'ultimo, dove stampo `\n` — stesso pattern già usato negli esercizi precedenti.
