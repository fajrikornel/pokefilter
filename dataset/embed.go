package dataset

import _ "embed"

//go:embed types.csv
var TypesCsv string

//go:embed damage_classes.csv
var DamageClassesCsv string

//go:embed abilities.csv
var AbilitiesCsv string

//go:embed moves.csv
var MovesCsv string

//go:embed pokemon.csv
var PokemonCsv string

//go:embed pokemon_abilities.csv
var PokemonAbilitiesCsv string

//go:embed pokemon_moves.csv
var PokemonMovesCsv string
