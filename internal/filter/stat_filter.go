package filter

import (
	"errors"
	"fmt"
)

type StatFilter struct {
	minHp    int
	maxHp    int
	minAtk   int
	maxAtk   int
	minDef   int
	maxDef   int
	minSpAtk int
	maxSpAtk int
	minSpDef int
	maxSpDef int
	minSpd   int
	maxSpd   int
	chain    Filter
}

func NewStatFilter(
	minHp int, maxHp int,
	minAtk int, maxAtk int,
	minDef int, maxDef int,
	minSpAtk int, maxSpAtk int,
	minSpDef int, maxSpDef int,
	minSpd int, maxSpd int,
) (*StatFilter, error) {
	return &StatFilter{
		minHp:    minHp,
		maxHp:    maxHp,
		minAtk:   minAtk,
		maxAtk:   maxAtk,
		minDef:   minDef,
		maxDef:   maxDef,
		minSpAtk: minSpAtk,
		maxSpAtk: maxSpAtk,
		minSpDef: minSpDef,
		maxSpDef: maxSpDef,
		minSpd:   minSpd,
		maxSpd:   maxSpd,
	}, nil
}

func (f *StatFilter) BuildQuery() (string, error) {
	query := fmt.Sprintf(`
		SELECT 
		    id as pokemon_id 
		FROM 
			pokemon p  
		WHERE 
		    p.hp >= %d AND p.hp <= %d AND 
		    p.atk >= %d AND p.atk <= %d AND 
		    p.def >= %d AND p.def <= %d AND 
		    p.sp_atk >= %d AND p.sp_atk <= %d AND 
		    p.sp_def >= %d AND p.sp_def <= %d AND 
			p.spd >= %d AND p.spd <= %d 
	`,
		f.minHp, f.maxHp,
		f.minAtk, f.maxAtk,
		f.minDef, f.maxDef,
		f.minSpAtk, f.maxSpAtk,
		f.minSpDef, f.maxSpDef,
		f.minSpd, f.maxSpd,
	)

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

func (f *StatFilter) Chain(other Filter) (Filter, error) {
	if f.chain != nil {
		return nil, errors.New("filter already chained")
	}

	f.chain = other
	return f, nil
}
