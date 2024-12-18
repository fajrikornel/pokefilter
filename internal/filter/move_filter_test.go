package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewMoveFilter(t *testing.T) {
	type args struct {
		move string
	}
	tests := []struct {
		name    string
		args    args
		want    *MoveFilter
		wantErr bool
	}{
		{
			name:    "should return error if move is empty",
			wantErr: true,
		},
		{
			name: "should return move filter if move is not empty",
			args: args{move: "hypnosis"},
			want: &MoveFilter{
				move: "hypnosis",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMoveFilter(tt.args.move)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoveFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMoveFilter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveFilter_BuildQuery(t *testing.T) {
	terminalFilterQuery, _ := NewTerminalFilter().BuildQuery()

	type fields struct {
		move  string
		chain Filter
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
				move: "hypnosis",
			},
			want: fmt.Sprintf(`
					SELECT 
						p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
					WHERE 
						m.name = 'hypnosis' AND 
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query and chain it with chained filter when filter is already chained",
			fields: fields{
				move:  "hypnosis",
				chain: NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						p.id as pokemon_id 
					FROM 
					    pokemon p 
						LEFT JOIN pokemon_moves pm on p.id = pm.pokemon_id 
						LEFT JOIN moves m on m.id = pm.move_id 
					WHERE 
						m.name = 'hypnosis' AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &MoveFilter{
				move:  tt.fields.move,
				chain: tt.fields.chain,
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

func TestMoveFilter_Chain(t *testing.T) {
	terminalFilter := NewTerminalFilter()

	type fields struct {
		move  string
		chain Filter
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
				move: "hypnosis",
			},
			args: args{
				other: terminalFilter,
			},
			want: &MoveFilter{
				move:  "hypnosis",
				chain: terminalFilter,
			},
			wantErr: false,
		},
		{
			name: "should return error if filter already chained",
			fields: fields{
				move:  "hypnosis",
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
			f := &MoveFilter{
				move:  tt.fields.move,
				chain: tt.fields.chain,
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
