# 34 - Triangoli Rettangoli con Asterischi

## Testo dell'esercizio
Scrivere un programma che legga da standard input un intero n > 1 e stampi il perimetro di due triangoli rettangoli con base e altezza di lunghezza n, posizionati a specchio verticale.

## Esecuzione

```bash
go run triangoli.go
```

## Esempio di output

```
Inserisci dimensione: 4
****
 * *
  **
   *
   *
   **
   * *
   ****
```

```
Inserisci dimensione: 6
******
 *   *
  *  *
   * *
    **
     *
     *
     **
     * *
     *  *
     *   *
     ******
```

## Concetti utilizzati

- Due cicli `for` separati per i due triangoli
- `strings.Repeat()` per costruire righe di spazi e asterischi in una chiamata sola
- Concatenazione di stringhe con `+`
- Gestione separata della prima riga (base), delle righe intermedie e del vertice
- Validazione dell'input con `return` anticipato
- Primo uso del package `strings`

## Note personali

> Esercizio più complesso finora sulla geometria con caratteri — richiede di ragionare su ogni riga come una combinazione di spazi a sinistra, asterisco sinistro, spazi interni e asterisco destro.
>
> Il primo approccio (**❌ BAD**) tentava di gestire entrambi i triangoli in un unico ciclo su una griglia `n*2 × n*2` — la logica diventava ingestibile con condizioni che si sovrapponevano e casi non coperti.
>
> Il secondo approccio (**✅ BETTER**) separa chiaramente le tre parti: triangolo superiore, vertice centrale (due righe singole) e triangolo inferiore (specchio del superiore). Ogni parte ha la propria logica semplice e leggibile.
>
> Primo esercizio in cui uso `strings.Repeat()` in modo sostanziale — costruire la stringa intera prima di stamparla in una chiamata sola è più efficiente di chiamare `fmt.Print()` per ogni carattere, come notato negli esercizi precedenti.
