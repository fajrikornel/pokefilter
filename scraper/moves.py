from typing import List
from pokemon_types import PType
import pokebase


class DamageClass:
    def __init__(self, id: int, name: str) -> None:
        if not id:
            raise Exception("damage class id invalid")
        elif not name:
            raise Exception("damage class name invalid")

        self.id = id
        self.name = name

    @staticmethod
    def get_all_damage_classes() -> List['DamageClass']:
        pokebase_damage_classes = pokebase.APIResource('move-damage-class', '')
        return DamageClass.__from_pokebase_damage_class_list(pokebase_damage_classes)

    @staticmethod
    def __from_pokebase_damage_class_list(obj: pokebase.APIResource) -> List['DamageClass']:
        damage_classes = []
        for damage_class in obj.results:
            damage_classes.append(
                DamageClass(
                    DamageClass.parse_id_from_url(damage_class.url),
                    damage_class.name
                )
            )

        return damage_classes

    @staticmethod
    def parse_id_from_url(damage_class_url: str) -> int:
        return int(damage_class_url.split("/")[-2])

    def __repr__(self) -> str:
        return f"damage_class[{self.id}]={self.name}"
    
    def __str__(self) -> str:
        return f"damage_class[{self.id}]={self.name}"


class Move:
    def __init__(self, id: int, name: str) -> None:
        if not id:
            raise Exception("move id invalid")
        elif not name:
            raise Exception("move name invalid")
        self.id = id
        self.name = name
        self.type_id = None
        self.damage_class_id = None

    @staticmethod
    def get_all_moves() -> List['Move']:
        pokebase_moves = pokebase.APIResource('move', '')
        return Move.__from_pokebase_move_list(pokebase_moves)

    @staticmethod
    def __from_pokebase_move_list(obj: pokebase.APIResource) -> List['Move']:
        moves = []
        for move in obj.results:
            moves.append(
                Move(
                    Move.parse_id_from_url(move.url),
                    move.name
                )
            )

        return moves
    
    def populate_type_and_damage_class(self) -> None:
        pokebase_move = pokebase.move(self.id)
        self.type_id = PType.parse_id_from_url(pokebase_move.type.url)
        self.damage_class_id = DamageClass.parse_id_from_url(pokebase_move.damage_class.url)

    @staticmethod
    def parse_id_from_url(move_url: str) -> int:
        return int(move_url.split("/")[-2])

    def __repr__(self) -> str:
        return f"move[{self.id}]={self.name}, type_id=[{self.type_id}], damage_class_id=[{self.damage_class_id}]"
    
    def __str__(self) -> str:
        return f"move[{self.id}]={self.name}, type_id=[{self.type_id}], damage_class_id=[{self.damage_class_id}]"


class PokemonMove:
    def __init__(self, pokemon_id: int, move_id: int) -> None:
        self.pokemon_id = pokemon_id
        self.move_id = move_id

    @staticmethod
    def from_pokemon_list(pokemon_list: List['Pokemon']) -> List['PokemonMove']:
        pokemon_move_list = []
        for pokemon in pokemon_list:
            for move_id in pokemon.move_ids:
                pokemon_move_list.append(
                    PokemonMove(
                        pokemon.id,
                        move_id
                    )
                )
        
        return pokemon_move_list

    def __repr__(self) -> str:
        return f"pokemon_move[{self.pokemon_id}]={self.move_id}"
    
    def __str__(self) -> str:
        return f"pokemon_move[{self.pokemon_id}]={self.move_id}"
