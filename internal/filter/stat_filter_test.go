package filter

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestNewStatFilter(t *testing.T) {
	type args struct {
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
	}
	tests := []struct {
		name string
		args args
		want *StatFilter
	}{
		{
			name: "should place correct arguments",
			args: args{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
			},
			want: &StatFilter{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := NewStatFilter(tt.args.minHp, tt.args.maxHp, tt.args.minAtk, tt.args.maxAtk, tt.args.minDef, tt.args.maxDef, tt.args.minSpAtk, tt.args.maxSpAtk, tt.args.minSpDef, tt.args.maxSpDef, tt.args.minSpd, tt.args.maxSpd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatFilter_BuildQuery(t *testing.T) {
	terminalFilterQuery, _ := NewTerminalFilter().BuildQuery()

	type fields struct {
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
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should build correct query and chain it with TerminalFilter when filter is not chained yet",
			fields: fields{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
			},
			want: fmt.Sprintf(`
					SELECT 
						id as pokemon_id 
					FROM pokemon p WHERE 
						p.hp >= 1 AND p.hp <= 2 AND 
						p.atk >= 3 AND p.atk <= 4 AND 
						p.def >= 5 AND p.def <= 6 AND 
						p.sp_atk >= 7 AND p.sp_atk <= 8 AND 
						p.sp_def >= 9 AND p.sp_def <= 10 AND 
						p.spd >= 11 AND p.spd <= 12 AND 
						p.id IN (%s)`, terminalFilterQuery),
			wantErr: false,
		},
		{
			name: "should build correct query and chain it with chained filter when filter is already chained",
			fields: fields{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
				chain: NewMockFilter(),
			},
			want: fmt.Sprintf(`
					SELECT 
						id as pokemon_id 
					FROM pokemon p WHERE 
						p.hp >= 1 AND p.hp <= 2 AND 
						p.atk >= 3 AND p.atk <= 4 AND 
						p.def >= 5 AND p.def <= 6 AND 
						p.sp_atk >= 7 AND p.sp_atk <= 8 AND 
						p.sp_def >= 9 AND p.sp_def <= 10 AND 
						p.spd >= 11 AND p.spd <= 12 AND 
						p.id IN (%s)`, mockFilterQuery),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &StatFilter{
				minHp:    tt.fields.minHp,
				maxHp:    tt.fields.maxHp,
				minAtk:   tt.fields.minAtk,
				maxAtk:   tt.fields.maxAtk,
				minDef:   tt.fields.minDef,
				maxDef:   tt.fields.maxDef,
				minSpAtk: tt.fields.minSpAtk,
				maxSpAtk: tt.fields.maxSpAtk,
				minSpDef: tt.fields.minSpDef,
				maxSpDef: tt.fields.maxSpDef,
				minSpd:   tt.fields.minSpd,
				maxSpd:   tt.fields.maxSpd,
				chain:    tt.fields.chain,
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

func TestStatFilter_Chain(t *testing.T) {
	terminalFilter := NewTerminalFilter()

	type fields struct {
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
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
			},
			args: args{
				other: terminalFilter,
			},
			want: &StatFilter{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
				chain: terminalFilter,
			},
			wantErr: false,
		},
		{
			name: "should return error if filter already chained",
			fields: fields{
				minHp: 1, maxHp: 2,
				minAtk: 3, maxAtk: 4,
				minDef: 5, maxDef: 6,
				minSpAtk: 7, maxSpAtk: 8,
				minSpDef: 9, maxSpDef: 10,
				minSpd: 11, maxSpd: 12,
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
			f := &StatFilter{
				minHp:    tt.fields.minHp,
				maxHp:    tt.fields.maxHp,
				minAtk:   tt.fields.minAtk,
				maxAtk:   tt.fields.maxAtk,
				minDef:   tt.fields.minDef,
				maxDef:   tt.fields.maxDef,
				minSpAtk: tt.fields.minSpAtk,
				maxSpAtk: tt.fields.maxSpAtk,
				minSpDef: tt.fields.minSpDef,
				maxSpDef: tt.fields.maxSpDef,
				minSpd:   tt.fields.minSpd,
				maxSpd:   tt.fields.maxSpd,
				chain:    tt.fields.chain,
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
