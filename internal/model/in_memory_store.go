package model

import (
	"database/sql"
	"encoding/csv"
	ds "github.com/fajrikornel/pokefilter/dataset"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type InMemorySqlStore struct {
	db *sql.DB
}

func NewInMemorySqlStore() (*InMemorySqlStore, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	sqlStore := &InMemorySqlStore{db: db}
	err = createTables(sqlStore)
	if err != nil {
		return nil, err
	}

	err = seedTables(sqlStore)
	if err != nil {
		return nil, err
	}

	return sqlStore, nil
}

func createTables(sqlStore *InMemorySqlStore) error {
	err := createTypesTable(sqlStore)
	if err != nil {
		return err
	}

	err = createDamageClassesTable(sqlStore)
	if err != nil {
		return err
	}

	err = createAbilitiesTable(sqlStore)
	if err != nil {
		return err
	}

	err = createMovesTable(sqlStore)
	if err != nil {
		return err
	}

	err = createPokemonTable(sqlStore)
	if err != nil {
		return err
	}

	err = createPokemonAbilitiesTable(sqlStore)
	if err != nil {
		return err
	}

	err = createPokemonMovesTable(sqlStore)
	if err != nil {
		return err
	}

	return nil
}

func seedTables(sqlStore *InMemorySqlStore) error {
	err := seedTypesTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedDamageClassesTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedAbilitiesTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedMovesTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedPokemonTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedPokemonAbilitiesTable(sqlStore)
	if err != nil {
		return err
	}

	err = seedPokemonMovesTable(sqlStore)
	if err != nil {
		return err
	}

	return nil
}

func createTypesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE types (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createDamageClassesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE damage_classes (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createAbilitiesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE abilities (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createMovesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE moves (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL,
					type_id INTEGER NOT NULL REFERENCES types(id),
					damage_class_id INTEGER NOT NULL REFERENCES damage_classes(id)
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createPokemonTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE pokemon (
					id SERIAL PRIMARY KEY,
					name TEXT NOT NULL,
					type_1_id INTEGER NOT NULL REFERENCES types(id),
					type_2_id INTEGER REFERENCES types(id),
					hp INTEGER NOT NULL,
					atk INTEGER NOT NULL,
					def INTEGER NOT NULL,
					sp_atk INTEGER NOT NULL,
					sp_def INTEGER NOT NULL,
					spd INTEGER NOT NULL
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createPokemonAbilitiesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE pokemon_abilities (
					id SERIAL PRIMARY KEY,
					pokemon_id INTEGER NOT NULL REFERENCES pokemon(id),
					ability_id INTEGER NOT NULL REFERENCES abilities(id)
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func createPokemonMovesTable(sqlStore *InMemorySqlStore) error {
	query := `
				CREATE TABLE pokemon_moves (
					id SERIAL PRIMARY KEY,
					pokemon_id INTEGER NOT NULL REFERENCES pokemon(id),
					move_id INTEGER NOT NULL REFERENCES moves(id)
				)`

	_, err := sqlStore.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func seedTypesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.TypesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO types (id, name) VALUES (?, ?)", row[0], row[1])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedDamageClassesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.DamageClassesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO damage_classes (id, name) VALUES (?, ?)", row[0], row[1])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedAbilitiesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.AbilitiesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO abilities (id, name) VALUES (?, ?)", row[0], row[1])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedMovesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.MovesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO moves (id, name, type_id, damage_class_id) VALUES (?, ?, ?, ?)", row[0], row[1], row[2], row[3])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedPokemonTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.PokemonCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO pokemon (id, name, type_1_id, type_2_id, hp, atk, def, sp_atk, sp_def, spd) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedPokemonAbilitiesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.PokemonAbilitiesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO pokemon_abilities (id, pokemon_id, ability_id) VALUES (?, ?, ?)", row[0], row[1], row[2])
		if err != nil {
			return err
		}
	}

	return nil
}

func seedPokemonMovesTable(sqlStore *InMemorySqlStore) error {
	results, err := readCsv(ds.PokemonMovesCsv)
	if err != nil {
		return err
	}

	for _, row := range results[1:] {
		_, err := sqlStore.db.Exec("INSERT INTO pokemon_moves (id, pokemon_id, move_id) VALUES (?, ?, ?)", row[0], row[1], row[2])
		if err != nil {
			return err
		}
	}

	return nil
}

func readCsv(c string) ([][]string, error) {
	csvReader := csv.NewReader(strings.NewReader(c))
	csvReader.Comma = ','
	results, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return results, nil
}
