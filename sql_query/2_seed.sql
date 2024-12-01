COPY types FROM '/mnt/types.csv' DELIMITER ',' CSV HEADER;
COPY damage_classes FROM '/mnt/damage_classes.csv' DELIMITER ',' CSV HEADER;
COPY abilities FROM '/mnt/abilities.csv' DELIMITER ',' CSV HEADER;
COPY moves FROM '/mnt/moves.csv' DELIMITER ',' CSV HEADER;
COPY pokemon FROM '/mnt/pokemon.csv' DELIMITER ',' CSV HEADER;
COPY pokemon_abilities FROM '/mnt/pokemon_abilities.csv' DELIMITER ',' CSV HEADER;
COPY pokemon_moves FROM '/mnt/pokemon_moves.csv' DELIMITER ',' CSV HEADER;
