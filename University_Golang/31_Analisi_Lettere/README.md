# 31 - Analisi delle Lettere

## Testo dell'esercizio
Scrivere un programma che legga da standard input una stringa senza spazi e stampi il numero di lettere maiuscole, minuscole, vocali e consonanti (considerando solo l'alfabeto inglese).

## Esecuzione

```bash
go run analisi_lettere.go
```

## Esempio di output

```
Inserisci una stringa di caratteri: Ciaoà
Maiuscole: 1
Minuscole: 3
Vocali: 3
Consonanti: 1
```

```
Inserisci una stringa di caratteri: aaAA
Maiuscole: 2
Minuscole: 2
Vocali: 4
Consonanti: 0
```

## Concetti utilizzati

- Variabile di tipo `string`
- Iterazione su una stringa con `for _, carattere := range`
- Tipo `rune` per la gestione corretta di caratteri UTF-8
- Confronto di caratteri con costanti letterali (`'A'`, `'z'`, ecc.) invece di magic number
- Contatori multipli aggiornati dentro un ciclo con condizionali annidati

## Note personali

> Primo esercizio con stringhe e caratteri — ho dovuto capire come Go gestisce i testi internamente.
>
> Il primo approccio (**❌ NOT-SO-GOOD**) usava `parola[i]` che restituisce un `byte`, non una `rune`. Funziona per ASCII puro ma è idiomaticamente sbagliato in Go, dove le stringhe sono UTF-8: un carattere può occupare più di un byte, e indicizzare con `[i]` rompe i caratteri multi-byte.
>
> Il secondo approccio (**✅ BETTER**) usa `for _, carattere := range parola` che decodifica automaticamente ogni carattere UTF-8 come `rune` — è il modo idiomatico di iterare su una stringa in Go.
>
> Ho sostituito i magic number (`65`, `97`, ecc.) con costanti letterali (`'A'`, `'a'`, ecc.) — chiunque legga il codice capisce subito senza dover consultare la tabella ASCII. Come nota il suggerimento del testo: i caratteri contigui nella tabella ASCII si possono confrontare direttamente con `>=` e `<=`.
>
> **Principio documentato nel codice:** il codice si legge molte più volte di quante si scrive — ogni miglioramento alla leggibilità riduce il carico cognitivo per chi legge, incluso il proprio futuro sé.
