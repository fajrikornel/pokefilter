package filter

import (
	"errors"
	"fmt"
	"math"
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

func (f *StatFilter) Display() string {

	var maxHp string
	if f.MaxHp == math.MaxInt {
		maxHp = "\u221E"
	} else {
		maxHp = fmt.Sprintf("%d", f.MaxHp)
	}

	var maxAtk string
	if f.MaxAtk == math.MaxInt {
		maxAtk = "\u221E"
	} else {
		maxAtk = fmt.Sprintf("%d", f.MaxAtk)
	}

	var maxDef string
	if f.MaxDef == math.MaxInt {
		maxDef = "\u221E"
	} else {
		maxDef = fmt.Sprintf("%d", f.MaxDef)
	}

	var maxSpAtk string
	if f.MaxSpAtk == math.MaxInt {
		maxSpAtk = "\u221E"
	} else {
		maxSpAtk = fmt.Sprintf("%d", f.MaxSpAtk)
	}

	var maxSpDef string
	if f.MaxSpDef == math.MaxInt {
		maxSpDef = "\u221E"
	} else {
		maxSpDef = fmt.Sprintf("%d", f.MaxSpDef)
	}

	var maxSpd string
	if f.MaxSpd == math.MaxInt {
		maxSpd = "\u221E"
	} else {
		maxSpd = fmt.Sprintf("%d", f.MaxSpd)
	}

	return fmt.Sprintf(`Pokemon stats filter:
HP => %d - %s
Attack => %d - %s
Defense => %d - %s
Sp. Attack => %d - %s
Sp. Defense => %d - %s
Speed => %d - %s`,
		f.MinHp, maxHp,
		f.MinAtk, maxAtk,
		f.MinDef, maxDef,
		f.MinSpAtk, maxSpAtk,
		f.MinSpDef, maxSpDef,
		f.MinSpd, maxSpd,
	)
}
