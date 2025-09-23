DELETE FROM heroes.races;

INSERT INTO heroes.races (
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
    ('Hobbit', 'Rápidos e ligeiros, hobbits se destacam em furtividade e agilidade.',0,2,0,2),
    ('Orc', 'Ferozes e brutais, orcs são temidos por sua força e ferocidade.',2,2,0,0),
    ('Troll', 'Grandes e poderosos, trolls possuem força bruta incomparável.',4,0,0,0),
    ('Morto Vivo', 'Criaturas reanimadas, mortos vivos possuem resistência sobrenatural.',1,0,1,2),
    ('Aranha', 'Ágeis e astutas, aranhas são mestres em emboscadas e ataques surpresa.',0,3,1,0),
    ('Dragão', 'Seres lendários, dragões possuem força, inteligência e magia excepcionais.',3,1,3,1),
    ('Lobo', 'Ferozes e leais, lobos são caçadores habilidosos e companheiros fiéis.',2,2,0,1),
    ('Golem', 'Construtos mágicos, golems são incrivelmente resistentes e fortes.',4,0,0,1),
    ('Fada', 'Pequenas e encantadoras, fadas possuem habilidades mágicas únicas.',0,1,3,2),
    ('Centauro', 'Metade humano, metade cavalo, centauros são rápidos e fortes.',2,2,1,0),
    ('Minotauro', 'Criaturas poderosas com corpo de touro, minotauros são guerreiros formidáveis.',3,1,0,1),
    ('Kobold', 'Pequenos e astutos, kobolds são mestres em armadilhas e emboscadas.',0,2,1,1)
;