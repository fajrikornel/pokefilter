package main

import (
	"fmt"
	pokefilter "github.com/fajrikornel/pokefilter/internal"
	"github.com/fajrikornel/pokefilter/internal/model"
	"github.com/fajrikornel/pokefilter/internal/prompt"
)

func main() {

	sqlStore, err := model.NewInMemorySqlStore()
	if err != nil {
		panic(err)
	}

	pokemonStore := model.NewPokemonStore(sqlStore)
	typeStore := model.NewTypeStore(sqlStore)
	abilityStore := model.NewAbilityStore(sqlStore)
	moveStore := model.NewMoveStore(sqlStore)
	damageClassStore := model.NewDamageClassStore(sqlStore)

	pf := pokefilter.NewPokeFilter(sqlStore, pokemonStore)

	statsFilterPrompter := prompt.NewStatFilterPrompter()
	typeFilterPrompter := prompt.NewTypeFilterPrompter(typeStore)
	abilityFilterPrompter := prompt.NewAbilityFilterPrompter(abilityStore)
	moveFilterPrompter := prompt.NewMoveFilterPrompter(moveStore)
	moveTypeFilterPrompter := prompt.NewMoveTypeFilterPrompter(typeStore, damageClassStore)

	cliPokefilterPrompter, err := prompt.NewCLIPokefilterPrompter(
		pf,
		statsFilterPrompter,
		typeFilterPrompter,
		abilityFilterPrompter,
		moveFilterPrompter,
		moveTypeFilterPrompter,
	)
	if err != nil {
		panic(err)
	}

	filteredPokemon, err := cliPokefilterPrompter.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("PokeFilter results: ")
	for i, pokemon := range filteredPokemon {
		fmt.Printf("%d. %s\n", i+1, pokemon.Name)
	}

	fmt.Println("Made by F. Kornel / 2024")

	var input string
	fmt.Println("Press the Enter key to close the program.")
	_, _ = fmt.Scanln(&input)
}
