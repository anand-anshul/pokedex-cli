# pokedex-cli

A simple Go-based CLI to explore, catch, and manage Pokémon using the PokeAPI. Runs as an interactive REPL with built-in caching for faster responses.

---

## Features

* Interactive CLI (REPL)
* Fetch Pokémon data from API
* Catch and track Pokémon
* In-memory caching with expiration
* Modular command-based design

---

## Installation

```bash
git clone https://github.com/anand-anshul/pokedex-cli.git
cd pokedex-cli
go build -o pokedex-cli
./pokedex-cli
```

---

## Commands

| Command          | Description                 |
| ---------------- | --------------------------- |
| `help`           | Show available commands     |
| `exit`           | Exit CLI                    |
| `map` / `mapb`   | Navigate locations          |
| `explore`        | List Pokémon in area        |
| `catch <name>`   | Catch a Pokémon             |
| `inspect <name>` | View caught Pokémon details |
| `pokedex`        | List caught Pokémon         |

---

## Architecture

* `main.go` → app entry point
* `repl.go` → command loop
* `internal/pokeapi/` → API client
* `internal/pokecache/` → cache layer
* command files → CLI actions

---

## Testing

```bash
go test ./...
```

---

## Notes

* Requires internet (uses PokeAPI)
* Config is currently hardcoded
* Some commands may be partially implemented

---

## Future Improvements

* Better command abstraction
* Input validation & error handling
* Retry + rate limiting
* Persistent cache
* Improved test coverage

