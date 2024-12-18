package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewAbilityFilter(t *testing.T) {
	type args struct {
		ability string
	}
	tests := []struct {
		name    string
		args    args
		want    *AbilityFilter
		wantErr bool
	}{
		{
			name:    "should return error if ability is empty",
			wantErr: true,
		},
		{
			name: "should return ability filter if ability is not empty",
			args: args{ability: "no-guard"},
			want: &AbilityFilter{
				Ability: "no-guard",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAbilityFilter(tt.args.ability)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAbilityFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAbilityFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbilityFilter_BuildQuery(t *testing.T) {
	terminalFilterQuery, _ := NewTerminalFilter().BuildQuery()

	type fields struct {
		ability string
		chain   Filter
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should build correct query and chain it with TerminalFilter when filter is not chained yet",
			fields: fields{
				ability: "no-guard",
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_abilities pa on p.id = pa.pokemon_id 
						LEFT JOIN abilities a on a.id = pa.ability_id 
					WHERE 
						a.name = 'no-guard' AND 
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query and chain it with chained filter when filter is already chained",
			fields: fields{
				ability: "no-guard",
				chain:   NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_abilities pa on p.id = pa.pokemon_id 
						LEFT JOIN abilities a on a.id = pa.ability_id 
					WHERE 
						a.name = 'no-guard' AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &AbilityFilter{
				Ability: tt.fields.ability,
				chain:   tt.fields.chain,
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

func TestAbilityFilter_Chain(t *testing.T) {
	terminalFilter := NewTerminalFilter()

	type fields struct {
		ability string
		chain   Filter
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
				ability: "no-guard",
			},
			args: args{
				other: terminalFilter,
			},
			want: &AbilityFilter{
				Ability: "no-guard",
				chain:   terminalFilter,
			},
			wantErr: false,
		},
		{
			name: "should return error if filter already chained",
			fields: fields{
				ability: "no-guard",
				chain:   terminalFilter,
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
			f := &AbilityFilter{
				Ability: tt.fields.ability,
				chain:   tt.fields.chain,
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
