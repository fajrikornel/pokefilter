from typing import List
from pokemon_types import PType
from abilities import Ability
from moves import Move
import pokebase

class Pokemon:
    def __init__(self, id: int, name: str) -> None:
        self.id = id
        self.name = name
        self.type_1_id = None
        self.type_2_id = None
        self.hp = None
        self.atk = None
        self.defn = None
        self.sp_atk = None
        self.sp_def = None
        self.spd = None

        self.ability_ids = []
        self.move_ids = []

    @staticmethod
    def get_all_pokemon_without_stats() -> List['Pokemon']:
        pokebase_pokemon = pokebase.APIResource('pokemon', '')
        return Pokemon.__from_pokebase_pokemon_list_without_stats(pokebase_pokemon)

    @staticmethod
    def __from_pokebase_pokemon_list_without_stats(obj: pokebase.APIResource) -> List['Pokemon']:
        pokemon_list = []
        for pokemon in obj.results:
            pokemon_list.append(
                Pokemon(
                    Pokemon.__parse_id_from_url(pokemon.url),
                    pokemon.name
                )
            )

        return pokemon_list
    
    @staticmethod
    def __parse_id_from_url(pokemon_url: str) -> int:
        return int(pokemon_url.split("/")[-2])
    
    def populate_stats_and_abilities(self) -> None:
        pokemon_obj = pokebase.pokemon(self.id)
        for pokemon_type in pokemon_obj.types:
            if pokemon_type.slot == 1:
                self.type_1_id = PType.parse_id_from_url(pokemon_type.type.url)
            elif pokemon_type.slot == 2:
                self.type_2_id =  PType.parse_id_from_url(pokemon_type.type.url)
        
        for pokemon_stat in pokemon_obj.stats:
            if pokemon_stat.stat.name == "hp":
                self.hp = pokemon_stat.base_stat
            elif pokemon_stat.stat.name == "attack":
                self.atk = pokemon_stat.base_stat
            elif pokemon_stat.stat.name == "defense":
                self.defn = pokemon_stat.base_stat
            elif pokemon_stat.stat.name == "special-attack":
                self.sp_atk = pokemon_stat.base_stat
            elif pokemon_stat.stat.name == "special-defense":
                self.sp_def = pokemon_stat.base_stat
            elif pokemon_stat.stat.name == "speed":
                self.spd = pokemon_stat.base_stat

        for pokemon_ability in pokemon_obj.abilities:
            self.ability_ids.append(Ability.parse_id_from_url(pokemon_ability.ability.url))

        for pokemon_move in pokemon_obj.moves:
            self.move_ids.append(Move.parse_id_from_url(pokemon_move.move.url))

    def __repr__(self) -> str:
        return f"pokemon[{self.id}]={self.name}, type_1_id={self.type_1_id} type_2_id={self.type_2_id} hp={self.hp} atk={self.atk} defn={self.defn} sp_atk={self.sp_atk} sp_def={self.sp_def} spd={self.spd}, ability_ids={self.ability_ids}, move_ids={self.move_ids}"
    
    def __str__(self) -> str:
        return f"pokemon[{self.id}]={self.name}, type_1_id={self.type_1_id} type_2_id={self.type_2_id} hp={self.hp} atk={self.atk} defn={self.defn} sp_atk={self.sp_atk} sp_def={self.sp_def} spd={self.spd}, ability_ids={self.ability_ids}, move_ids={self.move_ids}"
