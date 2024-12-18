package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewMoveTypeFilter(t *testing.T) {
	type args struct {
		moveType        string
		moveDamageClass string
	}
	tests := []struct {
		name    string
		args    args
		want    *MoveTypeFilter
		wantErr bool
	}{
		{
			name:    "should return error if both move type and move damage class is empty",
			wantErr: true,
		},
		{
			name: "should return move type filter",
			args: args{moveType: "psychic", moveDamageClass: "special"},
			want: &MoveTypeFilter{
				moveType:        "psychic",
				moveDamageClass: "special",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMoveTypeFilter(tt.args.moveType, tt.args.moveDamageClass)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoveTypeFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoveTypeFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveTypeFilter_BuildQuery(t *testing.T) {
	terminalFilterQuery, _ := NewTerminalFilter().BuildQuery()

	type fields struct {
		moveType        string
		moveDamageClass string
		chain           Filter
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should build correct query if moveType and moveDamageClass is specified",
			fields: fields{
				moveType:        "psychic",
				moveDamageClass: "physical",
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
						LEFT JOIN types t on t.id = m.type_id 
						LEFT JOIN damage_classes dc on dc.id = m.damage_class_id 
					WHERE 
						t.name = 'psychic' AND dc.name = 'physical' AND 
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query if only moveType is specified",
			fields: fields{
				moveType: "psychic",
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
						LEFT JOIN types t on t.id = m.type_id 
					WHERE 
						t.name = 'psychic' AND  
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query if only moveDamageClass is specified",
			fields: fields{
				moveDamageClass: "physical",
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
						LEFT JOIN damage_classes dc on dc.id = m.damage_class_id 
					WHERE 
						dc.name = 'physical' AND  
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query and chain it with chained filter when filter is already chained",
			fields: fields{
				moveType:        "psychic",
				moveDamageClass: "physical",
				chain:           NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						DISTINCT p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
						LEFT JOIN types t on t.id = m.type_id 
						LEFT JOIN damage_classes dc on dc.id = m.damage_class_id 
					WHERE 
						t.name = 'psychic' AND dc.name = 'physical' AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &MoveTypeFilter{
				moveType:        tt.fields.moveType,
				moveDamageClass: tt.fields.moveDamageClass,
				chain:           tt.fields.chain,
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

func TestMoveTypeFilter_Chain(t *testing.T) {
	terminalFilter := NewTerminalFilter()

	type fields struct {
		moveType        string
		moveDamageClass string
		chain           Filter
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
				moveType:        "grass",
				moveDamageClass: "special",
			},
			args: args{
				other: terminalFilter,
			},
			want: &MoveTypeFilter{
				moveType:        "grass",
				moveDamageClass: "special",
				chain:           terminalFilter,
			},
			wantErr: false,
		},
		{
			name: "should return error if filter already chained",
			fields: fields{
				moveType:        "grass",
				moveDamageClass: "special",
				chain:           terminalFilter,
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
			f := &MoveTypeFilter{
				moveType:        tt.fields.moveType,
				moveDamageClass: tt.fields.moveDamageClass,
				chain:           tt.fields.chain,
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
