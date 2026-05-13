# 24 - Divisori Propri

## Testo dell'esercizio
Scrivere un programma che legga da standard input un numero intero n e stampi a video i divisori propri del numero, ovvero tutti i suoi divisori escluso il numero stesso.

Esempio: i divisori propri di 12 sono 1, 2, 3, 4, 6.

## Esecuzione

```bash
go run divisori.go
```

## Esempio di output

```
Inserisci numero: 6
Divisori di 6: 1 2 3
```

```
Inserisci numero: 10
Divisori di 10: 1 2 5
```

## Concetti utilizzati

- Ciclo `for` da 1 a n-1 (escluso il numero stesso)
- Operatore modulo `%` per verificare la divisibilità
- `fmt.Printf()` per unificare stampa di valore e spazio in una chiamata sola
- Validazione dell'input con `return` anticipato

## Note personali

> La logica è semplice: itero da 1 a n-1 e stampo ogni numero per cui la divisione non produce resto. Il ciclo parte da 1 (non da 0, che causerebbe una divisione per zero) e si ferma prima di n (che per definizione non è un divisore proprio).
> Ho migliorato due `fmt.Print()` separati in un unico `fmt.Printf("%d ", i)` — ogni chiamata a una funzione ha un piccolo costo, unirle quando possibile è buona abitudine anche se qui la differenza è minima.
> Ho aggiunto la validazione di `numero <= 0` — con 0 il ciclo non eseguirebbe nulla silenziosamente, con un negativo il comportamento sarebbe indefinito.
