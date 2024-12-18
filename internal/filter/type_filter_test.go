package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewTypeFilter(t *testing.T) {
	type args struct {
		type1        string
		type2        string
		excludeTypes []string
	}
	tests := []struct {
		name    string
		args    args
		want    *TypeFilter
		wantErr bool
	}{
		{
			name:    "should return error if type1 and type2 both empty",
			wantErr: true,
		},
		{
			name:    "should return error if type2 and excludeTypes both specified",
			args:    args{type1: "normal", type2: "fairy", excludeTypes: []string{"electric"}},
			wantErr: true,
		},
		{
			name:    "should return error if type1 is included and one of excluded types is type1",
			args:    args{type1: "normal", type2: "", excludeTypes: []string{"normal"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTypeFilter(tt.args.type1, tt.args.type2, tt.args.excludeTypes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTypeFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTypeFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeFilter_BuildQuery(t *testing.T) {
	terminalFilterQuery, _ := NewTerminalFilter().BuildQuery()

	type fields struct {
		type1        string
		type2        string
		excludeTypes []string
		chain        Filter
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should build correct query if type1 and type2 specified",
			fields: fields{
				type1: "normal",
				type2: "flying",
				chain: NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN types t1 on p.type_1_id = t1.id 
						LEFT JOIN types t2 on p.type_2_id = t2.id 
					WHERE 
						((t1.name = 'normal' AND t2.name = 'flying') OR 
						(t2.name = 'normal' AND t1.name = 'flying')) AND  
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query if only type1 is specified",
			fields: fields{
				type1: "normal",
				chain: NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN types t1 on p.type_1_id = t1.id 
						LEFT JOIN types t2 on p.type_2_id = t2.id 
					WHERE 
						(t1.name = 'normal' OR t2.name = 'normal') AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query if type1 and excludeTypes specified",
			fields: fields{
				type1:        "fire",
				excludeTypes: []string{"fighting", "psychic", "flying", "dark"},
				chain:        NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN types t1 on p.type_1_id = t1.id 
						LEFT JOIN types t2 on p.type_2_id = t2.id 
					WHERE 
						((t1.name = 'fire' AND (t2.name NOT IN ('fighting', 'psychic', 'flying', 'dark') OR t2.name IS NULL)) OR 
						(t2.name = 'fire' AND (t1.name NOT IN ('fighting', 'psychic', 'flying', 'dark') OR t1.name IS NULL))) AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query if only excludeTypes specified",
			fields: fields{
				excludeTypes: []string{"fighting", "psychic", "flying", "dark"},
				chain:        NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN types t1 on p.type_1_id = t1.id 
						LEFT JOIN types t2 on p.type_2_id = t2.id 
					WHERE 
						(t1.name NOT IN ('fighting', 'psychic', 'flying', 'dark') AND (t2.name IS NULL OR t2.name NOT IN ('fighting', 'psychic', 'flying', 'dark'))) AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query and chain it with terminal filter when filter is not chained yet",
			fields: fields{
				type1: "normal",
				type2: "flying",
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN types t1 on p.type_1_id = t1.id 
						LEFT JOIN types t2 on p.type_2_id = t2.id 
					WHERE 
						((t1.name = 'normal' AND t2.name = 'flying') OR 
						(t2.name = 'normal' AND t1.name = 'flying')) AND  
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &TypeFilter{
				type1:        tt.fields.type1,
				type2:        tt.fields.type2,
				excludeTypes: tt.fields.excludeTypes,
				chain:        tt.fields.chain,
			}
			got, err := f.BuildQuery()
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			newLineTabRegex, _ := regexp.Compile("[\n\t]")
			got = newLineTabRegex.ReplaceAllString(got, "")
			tt.want = newLineTabRegex.ReplaceAllString(tt.want, "")

			spaceRegex, _ := regexp.Compile("\\s+")
			got = spaceRegex.ReplaceAllString(got, " ")
			tt.want = spaceRegex.ReplaceAllString(tt.want, " ")

			if got != tt.want {
				t.Errorf("BuildQuery() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeFilter_Chain(t *testing.T) {
	terminalFilter := NewTerminalFilter()

	type fields struct {
		type1        string
		type2        string
		excludeTypes []string
		chain        Filter
	}
	type args struct {
		other Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Filter
		wantErr bool
	}{
		{
			name: "should add chain to filter if filter not chained yet",
			fields: fields{
				type1: "normal",
				type2: "fairy",
			},
			args: args{
				other: terminalFilter,
			},
			want: &TypeFilter{
				type1: "normal",
				type2: "fairy",
				chain: terminalFilter,
			},
			wantErr: false,
		},
		{
			name: "should return error if filter already chained",
			fields: fields{
				type1: "normal",
				type2: "fairy",
				chain: terminalFilter,
			},
			args: args{
				other: terminalFilter,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &TypeFilter{
				type1:        tt.fields.type1,
				type2:        tt.fields.type2,
				excludeTypes: tt.fields.excludeTypes,
				chain:        tt.fields.chain,
			}
			got, err := f.Chain(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain() got = %v, want %v", got, tt.want)
			}
		})
	}
}
