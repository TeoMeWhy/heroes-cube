DELETE FROM classes;

INSERT INTO classes (
    name,
    description,
    initial_strength,
    initial_dexterity,
    initial_intelligence,
    initial_wisdom
) VALUES
('Warrior', 'A strong and resilient fighter, skilled in melee combat.', 5, 3, 2, 2),
('Rogue', 'A stealthy and agile character, adept at sneaking and dealing critical hits.', 2, 6, 2, 2),
('Mage', 'A master of arcane arts, capable of casting powerful spells.', 1, 3, 6, 2),
('Cleric', 'A holy warrior with healing abilities and divine magic.', 1, 2, 3, 6);

