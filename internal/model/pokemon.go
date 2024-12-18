package model

import (
	"fmt"
	"strings"
)

type PokemonStore struct {
	sqlStore *InMemorySqlStore
}

type Pokemon struct {
	id      int
	name    string
	type1Id int
	type2Id int
	hp      int
	atk     int
	def     int
	spAtk   int
	spDef   int
	spd     int
}

func NewPokemonStore(sqlStore *InMemorySqlStore) *PokemonStore {
	return &PokemonStore{
		sqlStore: sqlStore,
	}
}

func (m *PokemonStore) GetPokemonByIds(pokemonIds []int) ([]Pokemon, error) {
	transformedIds := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(pokemonIds)), ","), "[]")
	res, err := m.sqlStore.ExecuteQuery(fmt.Sprintf("SELECT id, name, type_1_id, type_2_id, hp, atk, def, sp_atk, sp_def, spd FROM pokemon WHERE id IN (%s)", transformedIds))
	if err != nil {
		return nil, err
	}
	defer res.Close()

	pokemon := make([]Pokemon, 0)
	for res.Next() {
		var r Pokemon
		err = res.Scan(&r.id, &r.name, &r.type1Id, &r.type2Id, &r.hp, &r.atk, &r.def, &r.spAtk, &r.spDef, &r.spd)
		if err != nil {
			return nil, err
		}

		pokemon = append(pokemon, r)
	}

	return pokemon, nil
}

func (p Pokemon) GetName() string {
	return p.name
}
