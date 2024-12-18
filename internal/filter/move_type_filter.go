package filter

import (
	"errors"
	"fmt"
)

type MoveTypeFilter struct {
	moveType        string
	moveDamageClass string
	chain           Filter
}

func NewMoveTypeFilter(moveType, moveDamageClass string) (*MoveTypeFilter, error) {
	if moveType == "" && moveDamageClass == "" {
		return nil, errors.New("both move type and move damage class is empty")
	}

	return &MoveTypeFilter{
		moveType:        moveType,
		moveDamageClass: moveDamageClass,
	}, nil
}

func (f *MoveTypeFilter) BuildQuery() (string, error) {
	var query string
	if f.moveType != "" && f.moveDamageClass != "" {
		query = fmt.Sprintf(`
			SELECT 
				DISTINCT p.id as pokemon_id 
			FROM 
				pokemon p 
				LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
				LEFT JOIN moves m on m.id = pm.move_id 
				LEFT JOIN types t on t.id = m.type_id 
				LEFT JOIN damage_classes dc on dc.id = m.damage_class_id 
			WHERE 
				t.name = '%s' AND dc.name = '%s' 
		`, f.moveType, f.moveDamageClass)
	} else if f.moveType != "" {
		query = fmt.Sprintf(`
			SELECT 
				DISTINCT p.id as pokemon_id 
			FROM 
				pokemon p 
				LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
				LEFT JOIN moves m on m.id = pm.move_id 
				LEFT JOIN types t on t.id = m.type_id 
			WHERE 
				t.name = '%s' 
		`, f.moveType)
	} else {
		query = fmt.Sprintf(`
			SELECT 
				DISTINCT p.id as pokemon_id 
			FROM 
				pokemon p 
				LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
				LEFT JOIN moves m on m.id = pm.move_id 
				LEFT JOIN damage_classes dc on dc.id = m.damage_class_id 
			WHERE 
				dc.name = '%s' 
		`, f.moveDamageClass)
	}

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

func (f *MoveTypeFilter) Chain(other Filter) (Filter, error) {
	if f.chain != nil {
		return nil, errors.New("filter already chained")
	}

	f.chain = other
	return f, nil
}
