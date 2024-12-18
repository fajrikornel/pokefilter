package filter

import (
	"errors"
	"fmt"
	"strings"
)

type TypeFilter struct {
	Type1        string
	Type2        string
	ExcludeTypes []string
	chain        Filter
}

func NewTypeFilter(type1, type2 string, excludeTypes []string) (*TypeFilter, error) {
	if type1 == "" && type2 == "" && (excludeTypes == nil || len(excludeTypes) == 0) {
		return nil, errors.New("type1, type2, excludeTypes cannot be empty")
	}

	if type2 != "" && len(excludeTypes) > 0 {
		return nil, errors.New("cannot specify both types and specify excludeTypes too")
	}

	if type1 != "" && excludeTypes != nil && len(excludeTypes) > 0 {
		for _, excludeType := range excludeTypes {
			if type1 == excludeType {
				return nil, errors.New("excluded type is filtered")
			}
		}
	}

	return &TypeFilter{
		Type1:        type1,
		Type2:        type2,
		ExcludeTypes: excludeTypes,
	}, nil
}

func (f *TypeFilter) BuildQuery() (string, error) {
	var query string
	if f.Type2 != "" {
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
			`, f.Type1, f.Type2, f.Type1, f.Type2)
	} else if f.Type1 != "" {
		if f.ExcludeTypes != nil && len(f.ExcludeTypes) > 0 {
			excludeTypesString := "'" + strings.Join(f.ExcludeTypes, "', '") + "'"
			query = fmt.Sprintf(`
				SELECT 
					DISTINCT p.id as pokemon_id 
				FROM 
					pokemon p 
					LEFT JOIN types t1 on p.type_1_id = t1.id 
					LEFT JOIN types t2 on p.type_2_id = t2.id 
				WHERE 
					((t1.name = '%s' AND (t2.name NOT IN (%s) OR t2.name IS NULL)) OR 
					(t2.name = '%s' AND (t1.name NOT IN (%s) OR t1.name IS NULL))) 
				`, f.Type1, excludeTypesString, f.Type1, excludeTypesString)
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
				`, f.Type1, f.Type1)
		}
	} else {
		excludeTypesString := "'" + strings.Join(f.ExcludeTypes, "', '") + "'"
		query = fmt.Sprintf(`
			SELECT 
				DISTINCT p.id as pokemon_id 
			FROM 
				pokemon p 
				LEFT JOIN types t1 on p.type_1_id = t1.id 
				LEFT JOIN types t2 on p.type_2_id = t2.id 
			WHERE 
				(t1.name NOT IN (%s) AND (t2.name IS NULL OR t2.name NOT IN (%s))) 
			`, excludeTypesString, excludeTypesString)
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

func (f *TypeFilter) Display() string {
	if f.Type2 != "" {
		return fmt.Sprintf("Pokemon has the type %s and %s", f.Type1, f.Type2)
	} else if f.Type1 != "" {
		if f.ExcludeTypes != nil && len(f.ExcludeTypes) > 0 {
			excludeTypesString := "'" + strings.Join(f.ExcludeTypes, "', '") + "'"
			return fmt.Sprintf("Pokemon has the type %s but does not have type %s", f.Type1, excludeTypesString)
		} else {
			return fmt.Sprintf("Pokemon has the type %s", f.Type1)
		}
	} else {
		excludeTypesString := "'" + strings.Join(f.ExcludeTypes, "', '") + "'"
		return fmt.Sprintf("Pokemon does not have type %s", excludeTypesString)
	}
}
