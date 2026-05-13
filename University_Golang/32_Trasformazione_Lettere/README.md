# 32 - Trasformazione Maiuscolo/Minuscolo

## Testo dell'esercizio
Scrivere un programma che legga da standard input una stringa e la ristampi due volte: la prima in maiuscolo e la seconda in minuscolo, considerando solo le lettere dell'alfabeto inglese.

## Esecuzione

```bash
go run trasformazioni_lettere.go
```

## Esempio di output

```
Inserisci un testo: TestoDiProva!!!
Testo maiuscolo: TESTODIPROVA!!!
Testo minuscolo: testodiprova!!!
```

```
Inserisci un testo: Testo_Di_PrOvA....
Testo maiuscolo: TESTO_DI_PROVA....
Testo minuscolo: testo_di_prova....
```

## Concetti utilizzati

- Iterazione su stringa con `for _, carattere := range`
- Aritmetica sui codici ASCII: `carattere - 32` per maiuscolo, `carattere + 32` per minuscolo
- Conversione `rune` → `string` con `string()`
- `fmt.Println()` senza argomenti per il newline tra i due output

## Note personali

> Nella tabella ASCII le lettere maiuscole e minuscole distano sempre esattamente 32 posizioni (`'a' - 'A' == 32`), quindi sommare o sottrarre 32 a una runa converte direttamente tra i due casi senza bisogno di tabelle o condizioni aggiuntive.
>
> L'aritmetica `±32` è una scelta consapevole e corretta **in questo contesto specifico** — il testo dell'esercizio dice esplicitamente "considerando l'alfabeto inglese", quindi il vincolo ASCII è intenzionale. In un programma reale che deve gestire testi multilingua si userebbe `unicode.ToUpper()` e `unicode.ToLower()` del package `unicode`.
>
> **Lezione importante:** leggere attentamente il testo del problema cambia completamente cosa è un miglioramento e cosa è una scelta progettuale consapevole. Non ogni soluzione "più generale" è necessariamente migliore per il problema dato.
