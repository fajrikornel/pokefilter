CREATE TABLE types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE damage_classes (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE abilities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE moves (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type_id INTEGER NOT NULL REFERENCES types(id),
    damage_class_id INTEGER NOT NULL REFERENCES damage_classes(id)
);

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
);

CREATE TABLE pokemon_abilities (
    id SERIAL PRIMARY KEY,
    pokemon_id INTEGER NOT NULL REFERENCES pokemon(id),
    ability_id INTEGER NOT NULL REFERENCES abilities(id)
);

CREATE TABLE pokemon_moves (
    id SERIAL PRIMARY KEY,
    pokemon_id INTEGER NOT NULL REFERENCES pokemon(id),
    move_id INTEGER NOT NULL REFERENCES moves(id)
);
