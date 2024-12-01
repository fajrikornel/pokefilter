from typing import List
import pokebase

class PType:
    def __init__(self, id: int, name: str) -> None:
        if not id:
            raise Exception("type id invalid")
        elif not name:
            raise Exception("type name invalid")

        self.id = id
        self.name = name

    @staticmethod
    def get_all_types() -> List['PType']:
        pokebase_types = pokebase.APIResource('type', '')
        return PType.__from_pokebase_type_list(pokebase_types)

    @staticmethod
    def __from_pokebase_type_list(obj: pokebase.APIResource) -> List['PType']:
        types = []
        for type in obj.results:
            types.append(
                PType(
                    PType.parse_id_from_url(type.url),
                    type.name
                )
            )

        return types

    @staticmethod
    def parse_id_from_url(type_url: str) -> int:
        return int(type_url.split("/")[-2])

    def __repr__(self) -> str:
        return f"type[{self.id}]={self.name}"
    
    def __str__(self) -> str:
        return f"type[{self.id}]={self.name}"
