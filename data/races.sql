DELETE FROM races;

INSERT INTO races (
    name,
    description,
    mod_strength,
    mod_dexterity,
    mod_intelligence,
    mod_wisdom
) VALUES
    ('Human', 'Versatile and adaptable, humans excel in all areas.',1,1,1,1),
    ('Elf', 'Graceful and agile, elves are masters of archery and magic.',0,2,1,1),
    ('Dwarf', 'Sturdy and resilient, dwarves are skilled craftsmen and warriors.',2,1,0,1),
    ('Halfling', 'Quick and clever, halflings excel in stealth and agility.',0,2,0,1)
;