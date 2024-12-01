-- in general, modify the WHERE clause and the SELECT & GROUP BY expression

-- searching for a Normal type pokemon with base HP greater than 150
SELECT p.name FROM
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
    p.hp > 150
    AND (t1.name = 'normal' OR t2.name = 'normal')
GROUP BY p.name;

-- searching for a physically strong Ground type pokemon that learns a physical Electric type move(s)
SELECT p.name, m.name FROM
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
    p.atk > 140
    AND (t1.name = 'ground' OR t2.name = 'ground')
    AND (mt.name = 'electric' AND dc.name = 'physical')
GROUP BY p.name, m.name;

-- searching for a pokemon that has the ability Levitate and also learns Skill Swap
SELECT p.name, a.name, m.name FROM
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
    a.name = 'levitate'
    AND m.name = 'skill-swap'
GROUP BY p.name, a.name, m.name;

-- searching for a Fire/Ghost type pokemon
SELECT p.name, t1.name, t2.name FROM
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
    (t1.name = 'fire' AND t2.name = 'ghost') OR
    (t2.name = 'fire' AND t1.name = 'ghost')
GROUP BY p.name, t1.name, t2.name;

-- more complex, searching for a pokemon that learns both Sunny Day and Solar Beam
SELECT p_name FROM (
    SELECT p_name, count(*) FROM (
        SELECT p.name as p_name, m.name as m_name FROM
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
            (m.name = 'sunny-day' OR m.name = 'solar-beam')
        GROUP BY p.name, m.name
    ) q1
GROUP BY p_name
) q2 WHERE count = 2;
