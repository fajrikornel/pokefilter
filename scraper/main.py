import os
from tqdm import tqdm
from pokemon_types import PType
from pokemon import Pokemon
from abilities import Ability, PokemonAbility
from moves import DamageClass, Move, PokemonMove
from csv_writer import CsvWriter

if 'OUTPUT_DIRECTORY' not in os.environ:
    raise Exception("Please specify OUTPUT_DIRECTORY in the environment variables")

OUTPUT_DIRECTORY = os.environ['OUTPUT_DIRECTORY']

print("Fetching pokemon types...")
types = PType.get_all_types()

print("Fetching pokemon list...")
pokemon_list = Pokemon.get_all_pokemon_without_stats()

print("Fetching pokemon abilities list...")
abilities = Ability.get_all_abilities()

print("Fetching pokemon moves damage classes...")
damage_classes = DamageClass.get_all_damage_classes()

print("Fetching pokemon moves list...")
moves = Move.get_all_moves()

print("Fetching move type and damage classes for all moves...")
for move in tqdm(moves):
    move.populate_type_and_damage_class()

print("Fetching stats, abilities, and moves data for all pokemon...")
for pokemon in tqdm(pokemon_list):
    pokemon.populate_stats_and_abilities()

print("Gluing pokemon abilities, and moves...")
pokemon_abilities = PokemonAbility.from_pokemon_list(pokemon_list)
pokemon_moves = PokemonMove.from_pokemon_list(pokemon_list)

print("Saving results...")
csv_writer = CsvWriter(
    pokemon_list,
    types,
    abilities,
    pokemon_abilities,
    moves,
    damage_classes,
    pokemon_moves
)
csv_writer.save_data(OUTPUT_DIRECTORY)

print("Finished!")
