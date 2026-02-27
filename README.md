# Pokedex CLI

**Pokedex CLI** è un piccolo progetto in Go sviluppato per imparare la gestione delle chiamate Backend (BE) e l'interazione con le API REST utilizzando il linguaggio Go.

Il progetto fa parte del percorso formativo di [Boot.dev - Build a Pokedex CLI](https://www.boot.dev/courses/build-pokedex-cli-golang).

## Descrizione
L'applicazione è un'interfaccia a riga di comando (CLI) che interroga le [PokeAPI](https://pokeapi.co/) per recuperare informazioni dal mondo Pokémon. Attualmente permette di esplorare le aree geografiche (locations) del mondo Pokémon in modo paginato.

## Installazione

Assicurati di avere Go installato sul tuo sistema.

1. Clona la repository:
   ```bash
   git clone https://github.com/marintan01/pokedexcli
   cd pokedex
   ```

1. Avvia l'applicazione
   ```bash
   go run .
   ```

   
## Comandi Disponibili

Una volta avviata la CLI, puoi utilizzare i seguenti comandi:

| Comando | Descrizione |
| :--- | :--- |
| `help` | Visualizza il messaggio di aiuto con tutti i comandi disponibili. |
| `map` | Visualizza le successive 20 aree dei Pokémon. |
| `mapb` | Visualizza le precedenti 20 aree dei Pokémon. |
| `explore` | Visualilzza la lista dei pokemon presenti nell'area scelta |
| `catch` | Cattuare il pokemon scelto|
| `exit` | Chiude l'applicazione Pokedex. |
