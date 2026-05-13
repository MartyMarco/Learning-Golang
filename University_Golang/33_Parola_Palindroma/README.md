# 33 - Parola Palindroma

## Testo dell'esercizio
Scrivere un programma che legga da standard input una stringa senza spazi e stampi "Palindroma" se la stringa si legge uguale da sinistra a destra e da destra a sinistra, "Non palindroma" altrimenti.

## Esecuzione

```bash
go run parola_palindroma.go
```

## Esempio di output

```
Inserisci una parola: anna
Palindroma
```

```
Inserisci una parola: anni
Non palindroma
```

```
Inserisci una parola: èssè
Palindroma
```

## Concetti utilizzati

- Conversione `string` → `[]rune` per gestire correttamente Unicode
- Ciclo `for` che itera solo fino a metà stringa
- Flag booleano `controllo` inizializzato a `true`, falsificato al primo mismatch
- `break` per interrompere il ciclo appena trovata una coppia non corrispondente
- Confronto booleano idiomatico `if controllo` invece di `if controllo == true`

## Note personali

> Il primo approccio (**❌ NOT-SO-GOOD**) aveva un bug di indice: con `parola[len(parola)-i]` quando `i == 0` si accede a `parola[len(parola)]` che è fuori dai limiti dell'array — crash garantito a runtime.
>
> Il secondo approccio (**✅ BETTER**) converte la stringa in `[]rune` prima di tutto — fondamentale per supportare caratteri Unicode multi-byte come `è`. Indicizzare direttamente su una stringa Go restituisce byte, non caratteri, e rompe i caratteri multi-byte.
>
> Il ciclo si ferma a `lunghezza/2` — ogni coppia di caratteri simmetrici viene confrontata una sola volta, e con `break` ci si ferma subito al primo mismatch senza controllare il resto inutilmente.
>
> Ho sostituito `if controllo == true` con `if controllo` — confrontare un booleano con `true` è ridondante e meno idiomatico. In Go (e nella maggior parte dei linguaggi) il booleano si usa direttamente nella condizione.
