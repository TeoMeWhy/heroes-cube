DELETE FROM races;

INSERT INTO races (
    name,
    description,
    mod_strength,
    mod_dexterity,
    mod_intelligence,
    mod_wisdom
) VALUES
    ('Humano', 'Versátil e adaptável, humanos de destacam em todas áreas.',1,1,1,1),
    ('Elfo', 'Cheios de graça e ágeis, elfos são mestres de arco e magia.',0,2,1,1),
    ('Anão', 'Resistentes e resilientes, anões são excelentes construtores e guerreiros.',2,1,0,1),
    ('Hobbit', 'Rápidos e ligeiros, hobbits se destacam em furtividade e agilidade.',0,2,0,2)
;