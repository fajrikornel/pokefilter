# PokeFilter manual query

You can also import the dataset into an SQL database like PostgreSQL and build your own queries.

## Steps

Steps to query by yourself on PostgreSQL:

### Run the migrations

As provided in the `1_migration.sql` file.

### Seed the data

As provided in the `2_seed.sql` file. You need to change the file path for each of the CSVs.

### Run your own query!

Sample queries are listed on `3_sample_queries.sql`. You can query what you need, like this query that searches for a pokemon that has the ability Sturdy, learns Endeavor, and has a Speed base stat greater than 90:

```sql
SELECT p.name as pokemon_name, a.name as ability_name, m.name as move_name FROM
    pokemon p
    JOIN types t1 ON p.type_1_id = t1.id
    LEFT JOIN types t2 ON p.type_2_id = t2.id
    JOIN pokemon_abilities pa ON p.id = pa.pokemon_id
    JOIN abilities a ON a.id = pa.ability_id
    JOIN pokemon_moves pm ON p.id = pm.pokemon_id
    JOIN moves m ON m.id = pm.move_id
    JOIN damage_classes dc ON m.damage_class_id = dc.id
    JOIN types mt ON m.type_id = mt.id
WHERE
    a.name = 'sturdy'
    AND m.name = 'endeavor'
    AND p.spd > 90
GROUP BY p.name, a.name, m.name;
```

Results:
```markdown
   pokemon_name   | ability_name | move_name 
------------------+--------------+-----------
 togedemaru       | sturdy       | endeavor
 togedemaru-totem | sturdy       | endeavor
(2 rows)
```
