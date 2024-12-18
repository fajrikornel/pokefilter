package filter

import (
	"errors"
	"fmt"
	"strings"
)

type TypeFilter struct {
	type1        string
	type2        string
	excludeTypes []string
	chain        Filter
}

func NewTypeFilter(type1, type2 string, excludeTypes []string) (*TypeFilter, error) {
	if type1 == "" && type2 == "" {
		return nil, errors.New("type1 and type2 cannot be empty")
	}

	if type2 != "" && len(excludeTypes) > 0 {
		return nil, errors.New("cannot specify both types and specify excludeTypes too")
	}

	for _, excludeType := range excludeTypes {
		if type1 == excludeType {
			return nil, errors.New("excluded type is filtered")
		}
	}

	return &TypeFilter{
		type1:        type1,
		type2:        type2,
		excludeTypes: excludeTypes,
	}, nil
}

func (f *TypeFilter) BuildQuery() (string, error) {
	var query string
	if f.type2 != "" {
		query = fmt.Sprintf(`
			SELECT 
				DISTINCT p.id as pokemon_id 
			FROM 
				pokemon p 
				LEFT JOIN types t1 on p.type_1_id = t1.id 
				LEFT JOIN types t2 on p.type_2_id = t2.id 
			WHERE 
				((t1.name = '%s' AND t2.name = '%s') OR 
				(t2.name = '%s' AND t1.name = '%s')) 
			`, f.type1, f.type2, f.type1, f.type2)
	} else {
		if f.excludeTypes != nil && len(f.excludeTypes) > 0 {
			excludeTypesString := "'" + strings.Join(f.excludeTypes, "', '") + "'"
			query = fmt.Sprintf(`
				SELECT 
					DISTINCT p.id as pokemon_id 
				FROM 
					pokemon p 
					LEFT JOIN types t1 on p.type_1_id = t1.id 
					LEFT JOIN types t2 on p.type_2_id = t2.id 
				WHERE 
					((t1.name = '%s' AND t2.name NOT IN (%s)) OR 
					(t2.name = '%s' AND t1.name NOT IN (%s))) 
				`, f.type1, excludeTypesString, f.type1, excludeTypesString)
		} else {
			query = fmt.Sprintf(`
				SELECT 
					DISTINCT p.id as pokemon_id 
				FROM 
					pokemon p 
					LEFT JOIN types t1 on p.type_1_id = t1.id 
					LEFT JOIN types t2 on p.type_2_id = t2.id 
				WHERE 
					(t1.name = '%s' OR t2.name = '%s') 
				`, f.type1, f.type1)
		}
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

func (f *TypeFilter) Chain(other Filter) (Filter, error) {
	if f.chain != nil {
		return nil, errors.New("filter already chained")
	}

	f.chain = other
	return f, nil
}
