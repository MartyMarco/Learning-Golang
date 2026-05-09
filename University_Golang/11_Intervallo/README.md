# 11 - Valutazione Voto

## Testo dell'esercizio
Scrivere un programma che legga da standard input un voto da 0 a 100 e stampi la valutazione corrispondente:
- **Insufficiente** se il voto è inferiore a 60
- **Sufficiente** se il voto è compreso tra 60 (incluso) e 70 (escluso)
- **Buono** se il voto è compreso tra 70 (incluso) e 80 (escluso)
- **Distinto** se il voto è compreso tra 80 (incluso) e 90 (escluso)
- **Ottimo** se il voto è compreso tra 90 (incluso) e 100 (incluso)
- **Errore** se il voto è negativo o superiore a 100

## Esecuzione

```bash
go run intervallo.go
```

## Esempio di output

```
Inserisci il voto : 75
Buono
```

```
Inserisci il voto : 90
Ottimo
```

```
Inserisci il voto : 110
Errore
```

## Concetti utilizzati

- Dichiarazione di variabili con `var`
- Input da terminale con `fmt.Scan()`
- Output a terminale con `fmt.Println()`
- Struttura condizionale `if / else if / else` con più rami
- Operatori di confronto `<`, `>`, `<=`, `>=`
- Operatori logici `||` e `&&`

## Note personali

> Questo esercizio è stato il più articolato finora sul ragionamento condizionale — ho trovato tre approcci con difetti diversi.
>
> Il primo (**NOT-SO-GOOD**) ha un bug silenzioso: il controllo dell'errore è alla fine, ma un voto negativo rientra già nel ramo `Insufficiente` prima di arrivarci. Il programma non crasha ma dà una risposta sbagliata — il tipo di bug più pericoloso.
>
> Il secondo (**GOOD**) risolve spostando il controllo dell'errore all'inizio, ma ha condizioni ridondanti nei rami interni: scrivere `voto >= 60 && voto < 70` nel terzo ramo è inutile, perché se si arriva lì significa che i rami precedenti erano già falsi, quindi `voto >= 60` è già garantito.
>
> Il terzo (**BETTER**) sfrutta proprio questo: in una catena `if / else if`, ogni ramo implica che tutti i precedenti erano falsi. Basta quindi la sola condizione che "chiude" l'intervallo superiore, rendendo il codice più pulito e leggibile.
