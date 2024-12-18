package filter

import (
	"errors"
	"fmt"
)

type AbilityFilter struct {
	ability string
	chain   Filter
}

func NewAbilityFilter(ability string) (*AbilityFilter, error) {
	if ability == "" {
		return nil, errors.New("empty ability")
	}

	return &AbilityFilter{
		ability: ability,
	}, nil
}

func (f *AbilityFilter) BuildQuery() (string, error) {
	query := fmt.Sprintf(`
		SELECT 
		    DISTINCT p.id as pokemon_id 
		FROM 
			pokemon p 
			LEFT JOIN pokemon_abilities pa on p.id = pa.pokemon_id 
			LEFT JOIN abilities a on a.id = pa.ability_id 
		WHERE 
		    a.name = '%s' 
	`, f.ability)

	var chainQuery string
	var err error
	if f.chain != nil {
		chainQuery, err = f.chain.BuildQuery()
	} else {
		chainQuery, err = NewTerminalFilter().BuildQuery()
	}

	if err != nil {
		return "", err
	}

	query += fmt.Sprintf(" AND p.id IN (%s)", chainQuery)
	return query, nil
}

func (f *AbilityFilter) Chain(other Filter) (Filter, error) {
	if f.chain != nil {
		return nil, errors.New("filter already chained")
	}

	f.chain = other
	return f, nil
}
