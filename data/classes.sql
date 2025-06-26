DELETE FROM classes;

INSERT INTO classes (
    name,
    description,
    initial_strength,
    initial_dexterity,
    initial_intelligence,
    initial_wisdom
) VALUES
('Guerreiro', 'Um forte e resistente lutador.', 5, 3, 2, 2),
('Ladino', 'Um personagem furtivo e ágil, habilidoso em se esgueirar e desferir golpes críticos.', 2, 6, 2, 2),
('Mago', 'Um mestre das artes arcanas, capaz de lançar poderosos feitiços.', 1, 3, 6, 2),
('Clérigo', 'Um guerreiro sagrado com habilidades de cura e magia divina.', 1, 2, 3, 6);

