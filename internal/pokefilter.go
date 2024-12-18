package pokefilter

import (
	"github.com/fajrikornel/pokefilter/internal/filter"
	"github.com/fajrikornel/pokefilter/internal/model"
)

type PokeFilter struct {
	filters      []filter.Filter
	sqlStore     *model.InMemorySqlStore
	pokemonStore *model.PokemonStore
}

func NewPokeFilter(sqlStore *model.InMemorySqlStore, pokemonStore *model.PokemonStore) *PokeFilter {
	return &PokeFilter{
		filters:      []filter.Filter{},
		sqlStore:     sqlStore,
		pokemonStore: pokemonStore,
	}
}

func (p *PokeFilter) AddFilter(filter filter.Filter) {
	p.filters = append(p.filters, filter)
}

func (p *PokeFilter) ExecuteFilter() ([]model.Pokemon, error) {
	if len(p.filters) == 0 {
		return []model.Pokemon{}, nil
	}

	if len(p.filters) > 1 {
		for i, _ := range p.filters {
			if i == 0 {
				continue
			}
			
			_, err := p.filters[i-1].Chain(p.filters[i])
			if err != nil {
				return nil, err
			}
		}
	}

	query, err := p.filters[0].BuildQuery()
	if err != nil {
		return nil, err
	}

	res, err := p.sqlStore.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	pokemonIds := make([]int, 0)
	for res.Next() {
		var pokemonId int
		err := res.Scan(&pokemonId)
		if err != nil {
			return nil, err
		}

		pokemonIds = append(pokemonIds, pokemonId)
	}

	pokemon, err := p.pokemonStore.GetPokemonByIds(pokemonIds)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
