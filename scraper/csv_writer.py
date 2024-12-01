from typing import List
from pokemon import Pokemon
from pokemon_types import PType
from abilities import Ability, PokemonAbility
from moves import Move, DamageClass, PokemonMove
import csv


class CsvWriter:

    def __init__(
            self,
            pokemon_list: List[Pokemon],
            type_list: List[PType],
            ability_list: List[Ability],
            pokemon_abilities: List[PokemonAbility],
            move_list: List[Move],
            damage_classes: List[DamageClass],
            pokemon_moves: List[PokemonMove] 
    ):
        self.pokemon_list = pokemon_list
        self.type_list = type_list
        self.ability_list = ability_list
        self.pokemon_abilities = pokemon_abilities
        self.move_list = move_list
        self.damage_classes = damage_classes
        self.pokemon_moves = pokemon_moves

    def save_data(self, output_directory: str) -> None:
        self.__save_pokemon_csv(output_directory)
        self.__save_types_csv(output_directory)
        self.__save_abilities_csv(output_directory)
        self.__save_pokemon_abilities_csv(output_directory)
        self.__save_moves_csv(output_directory)
        self.__save_damage_classes_csv(output_directory)
        self.__save_pokemon_moves_csv(output_directory)

    def __save_pokemon_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'name',
            'type_1_id',
            'type_2_id',
            'hp',
            'atk',
            'def',
            'sp_atk',
            'sp_def',
            'spd',
        ]

        transformed_objs = []
        for pokemon in self.pokemon_list:
            transformed_objs.append({
                'id': pokemon.id,
                'name': pokemon.name,
                'type_1_id': pokemon.type_1_id,
                'type_2_id': pokemon.type_2_id,
                'hp': pokemon.hp,
                'atk': pokemon.atk,
                'def': pokemon.defn,
                'sp_atk': pokemon.sp_atk,
                'sp_def': pokemon.sp_def,
                'spd': pokemon.spd
            })

        with open(output_directory + "/pokemon.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_types_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'name'
        ]

        transformed_objs = []
        for type in self.type_list:
            transformed_objs.append({
                'id': type.id,
                'name': type.name
            })

        with open(output_directory + "/types.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_abilities_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'name'
        ]

        transformed_objs = []
        for ability in self.ability_list:
            transformed_objs.append({
                'id': ability.id,
                'name': ability.name
            })

        with open(output_directory + "/abilities.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_pokemon_abilities_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'pokemon_id',
            'ability_id'
        ]

        transformed_objs = []
        for i, pokemon_ability in enumerate(self.pokemon_abilities):
            transformed_objs.append({
                'id': i,
                'pokemon_id': pokemon_ability.pokemon_id,
                'ability_id': pokemon_ability.ability_id
            })

        with open(output_directory + "/pokemon_abilities.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_moves_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'name',
            'type_id',
            'damage_class_id'
        ]

        transformed_objs = []
        for move in self.move_list:
            transformed_objs.append({
                'id': move.id,
                'name': move.name,
                'type_id': move.type_id,
                'damage_class_id': move.damage_class_id
            })

        with open(output_directory + "/moves.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_damage_classes_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'name'
        ]

        transformed_objs = []
        for damage_class in self.damage_classes:
            transformed_objs.append({
                'id': damage_class.id,
                'name': damage_class.name
            })

        with open(output_directory + "/damage_classes.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)

    def __save_pokemon_moves_csv(self, output_directory: str) -> None:
        field_list = [
            'id',
            'pokemon_id',
            'move_id'
        ]

        transformed_objs = []
        for i, pokemon_move in enumerate(self.pokemon_moves):
            transformed_objs.append({
                'id': i,
                'pokemon_id': pokemon_move.pokemon_id,
                'move_id': pokemon_move.move_id
            })

        with open(output_directory + "/pokemon_moves.csv", 'w', newline='') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=field_list)
            writer.writeheader()
            writer.writerows(transformed_objs)
