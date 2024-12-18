# PokeFilter

PokeFilter is a lightweight standalone app that allows us to search a pokemon that matches specific filters, namely:
1. Stats filter; search a pokemon that has stats in the range we specify
2. Type filter; search a pokemon that has certain types, or pokemon that doesn't have certain types
2. Abilities filter; search a pokemon that has certain abilities
3. Moves filter; search a pokemon that can know a specific move, or in general a move of certain type and damage class

## How to use

Check the `Releases` section on the GitHub page, and download the version according to your platform. Just run the download in your terminal.

## Building from source

Run tests: 

```bash
go test -v -p 1 ./...
```

Build:

```bash
go build -o build/pokefilter cmd/pokefilter.go
```
