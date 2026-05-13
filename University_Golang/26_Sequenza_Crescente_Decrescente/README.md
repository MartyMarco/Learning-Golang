# 26 - Sequenza Crescente o Decrescente

## Testo dell'esercizio
Il programma legge da tastiera una serie di numeri maggiori di -1 e dopo ogni lettura stampa "crescente" se il nuovo valore è maggiore o uguale al precedente e "decrescente" altrimenti. Si assume che il primo numero inserito sia sempre preceduto da uno 0. Il programma si ferma quando il numero in input è <= -1.

## Esecuzione

```bash
go run sequenza_crescente_decrescente.go
```

## Esempio di output

```
Inserisci una sequenza di numeri:
1 3 2 5 6 3 1 -1
crescente
crescente
decrescente
crescente
crescente
decrescente
decrescente
```

## Concetti utilizzati

- Ciclo `for` infinito con `return` per uscita anticipata
- Variabile `primo` inizializzata a 0 come valore sentinella iniziale
- Confronto tra il valore corrente e il precedente
- `numero` dichiarato dentro il ciclo perché vive solo lì
- `return` invece di `break` per terminare direttamente il programma

## Note personali

> La chiave di questo esercizio è tenere traccia del numero precedente — uso `primo` inizializzato a 0 come da consegna, e lo aggiorno ad ogni iterazione con il valore appena letto.
> Ho usato `return` invece di `break` per interrompere il programma quando arriva il valore di stop — `break` avrebbe interrotto solo il ciclo continuando l'esecuzione dopo, `return` esce direttamente da `main`.
> La condizione di uscita è `numero < 0` invece di `numero <= -1` — sono equivalenti per gli interi, ma `< 0` è più diretto e idiomatico.
> Ho fatto attenzione alla condizione `primo <= numero` per "crescente" — include il caso di uguaglianza come specificato nel testo, dettaglio facile da dimenticare.
