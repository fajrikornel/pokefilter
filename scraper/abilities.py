from typing import List
import pokebase

class Ability:
    def __init__(self, id: int, name: str) -> None:
        if not id:
            raise Exception("ability id invalid")
        elif not name:
            raise Exception("ability name invalid")

        self.id = id
        self.name = name

    @staticmethod
    def get_all_abilities() -> List['Ability']:
        pokebase_abilities = pokebase.APIResource('ability', '')
        return Ability.__from_pokebase_ability_list(pokebase_abilities)

    @staticmethod
    def __from_pokebase_ability_list(obj: pokebase.APIResource) -> List['Ability']:
        abilities = []
        for ability in obj.results:
            abilities.append(
                Ability(
                    Ability.parse_id_from_url(ability.url),
                    ability.name
                )
            )

        return abilities

    @staticmethod
    def parse_id_from_url(ability_url: str) -> int:
        return int(ability_url.split("/")[-2])

    def __repr__(self) -> str:
        return f"ability[{self.id}]={self.name}"
    
    def __str__(self) -> str:
        return f"ability[{self.id}]={self.name}"

class PokemonAbility:
    def __init__(self, pokemon_id: int, ability_id: int) -> None:
        self.pokemon_id = pokemon_id
        self.ability_id = ability_id

    @staticmethod
    def from_pokemon_list(pokemon_list: List['Pokemon']) -> List['PokemonAbility']:
        pokemon_ability_list = []
        for pokemon in pokemon_list:
            for ability_id in pokemon.ability_ids:
                pokemon_ability_list.append(
                    PokemonAbility(
                        pokemon.id,
                        ability_id
                    )
                )
        
        return pokemon_ability_list

    def __repr__(self) -> str:
        return f"pokemon_ability[{self.pokemon_id}]={self.ability_id}"
    
    def __str__(self) -> str:
        return f"pokemon_ability[{self.pokemon_id}]={self.ability_id}"
