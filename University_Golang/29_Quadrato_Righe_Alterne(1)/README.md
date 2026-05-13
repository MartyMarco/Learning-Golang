# 29 - Quadrato con Righe Alterne (1)

## Testo dell'esercizio
Scrivere un programma che legga da standard input un numero intero n e stampi un quadrato di n righe alternando righe di `*` e righe di `+`, intervallati da spazi.

## Esecuzione

```bash
go run quadrato_righe_alterne1.go
```

## Esempio di output

```
Inserisci un numero: 5
* * * * *
+ + + + +
* * * * *
+ + + + +
* * * * *
```

## Concetti utilizzati

- Due cicli `for` annidati
- Operatore modulo `%` per distinguere righe pari e dispari
- Variabile `simbolo` per separare la scelta del carattere dalla logica di stampa
- `fmt.Println()` senza argomenti per il newline finale di riga
- Validazione dell'input con `return` anticipato

## Note personali

> Variante dell'esercizio precedente — stessa struttura a cicli annidati, con la logica aggiuntiva di alternare il simbolo tra righe pari e dispari.
>
> Il primo approccio (**❌ PRIMA**) controllava `i%2` due volte per ogni elemento del ciclo interno — quattro rami condizionali per quello che è essenzialmente una decisione sola per riga.
>
> Il secondo approccio (**✅ DOPO**) separa le due responsabilità: prima si decide il simbolo della riga intera con una variabile `simbolo`, poi il ciclo interno usa quella variabile senza sapere nulla di righe pari o dispari. Il risultato è più leggibile e il controllo `i%2` viene eseguito una sola volta per riga invece di n volte.
