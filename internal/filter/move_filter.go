package filter

import (
	"errors"
	"fmt"
)

type MoveFilter struct {
	move  string
	chain Filter
}

func NewMoveFilter(move string) (*MoveFilter, error) {
	if move == "" {
		return nil, errors.New("empty move")
	}

	return &MoveFilter{
		move: move,
	}, nil
}

func (f *MoveFilter) BuildQuery() (string, error) {
	query := fmt.Sprintf(`
		SELECT 
		    DISTINCT p.id as pokemon_id 
		FROM 
			pokemon p 
			LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
			LEFT JOIN moves m on m.id = pm.move_id 
		WHERE 
		    m.name = '%s' 
	`, f.move)

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

func (f *MoveFilter) Chain(other Filter) (Filter, error) {
	if f.chain != nil {
		return nil, errors.New("filter already chained")
	}

	f.chain = other
	return f, nil
}
