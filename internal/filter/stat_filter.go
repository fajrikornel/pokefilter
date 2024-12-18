package filter

import (
	"errors"
	"fmt"
)

type StatFilter struct {
	MinHp    int
	MaxHp    int
	MinAtk   int
	MaxAtk   int
	MinDef   int
	MaxDef   int
	MinSpAtk int
	MaxSpAtk int
	MinSpDef int
	MaxSpDef int
	MinSpd   int
	MaxSpd   int
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
		MinHp:    minHp,
		MaxHp:    maxHp,
		MinAtk:   minAtk,
		MaxAtk:   maxAtk,
		MinDef:   minDef,
		MaxDef:   maxDef,
		MinSpAtk: minSpAtk,
		MaxSpAtk: maxSpAtk,
		MinSpDef: minSpDef,
		MaxSpDef: maxSpDef,
		MinSpd:   minSpd,
		MaxSpd:   maxSpd,
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
		f.MinHp, f.MaxHp,
		f.MinAtk, f.MaxAtk,
		f.MinDef, f.MaxDef,
		f.MinSpAtk, f.MaxSpAtk,
		f.MinSpDef, f.MaxSpDef,
		f.MinSpd, f.MaxSpd,
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
