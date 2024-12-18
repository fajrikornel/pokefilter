package model

import (
	"fmt"
	"strings"
)

type PokemonStore struct {
	sqlStore *InMemorySqlStore
}

type Pokemon struct {
	Id      int
	Name    string
	Type1Id int
	Type2Id int // -1 means no type2
	Hp      int
	Atk     int
	Def     int
	SpAtk   int
	SpDef   int
	Spd     int
}

func NewPokemonStore(sqlStore *InMemorySqlStore) *PokemonStore {
	return &PokemonStore{
		sqlStore: sqlStore,
	}
}

func (m *PokemonStore) GetPokemonByIds(pokemonIds []int) ([]Pokemon, error) {
	transformedIds := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(pokemonIds)), ","), "[]")
	res, err := m.sqlStore.ExecuteQuery(fmt.Sprintf("SELECT id, name, type_1_id, COALESCE(-1, type_2_id), hp, atk, def, sp_atk, sp_def, spd FROM pokemon WHERE id IN (%s)", transformedIds))
	if err != nil {
		return nil, err
	}
	defer res.Close()

	pokemon := make([]Pokemon, 0)
	for res.Next() {
		var r Pokemon
		err = res.Scan(&r.Id, &r.Name, &r.Type1Id, &r.Type2Id, &r.Hp, &r.Atk, &r.Def, &r.SpAtk, &r.SpDef, &r.Spd)
		if err != nil {
			return nil, err
		}

		pokemon = append(pokemon, r)
	}

	return pokemon, nil
}

func (p Pokemon) GetName() string {
	return p.Name
}
