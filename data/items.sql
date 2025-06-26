DELETE FROM items;

INSERT INTO items (
    id, name, description, category, type, weight, price, damage, mod_strength, mod_dexterity, mod_intelligence, mod_wisdom
) VALUES
('1', 'Espada Longa', 'Uma espada longa de aço.', 'espada', 'arma', 5, 1500, 8, 2, 0, 0, 0),
('2', 'Espada Curta', 'Espada leve e rápida.', 'espada', 'arma', 3, 900, 5, 1, 1, 0, 0),
('3', 'Armadura de Couro', 'Armadura leve feita de couro.', 'armadura', 'armadura', 7, 1200, 0, 0, 1, 0, 0),
('4', 'Armadura de Placas', 'Armadura pesada de placas metálicas.', 'armadura', 'armadura', 15, 3500, 0, 2, 0, 0, 0),
('5', 'Botas de Agilidade', 'Botas que aumentam a destreza.', 'botas','botas', 1, 800, 0, 0, 2, 0, 0),
('6', 'Botas de Força', 'Botas reforçadas para guerreiros.', 'botas','botas', 2, 1000, 0, 1, 0, 0, 0),
('7', 'Cajado de Fogo', 'Cajado mágico que lança fogo.', 'cajado','arma', 4, 2000, 10, 0, 0, 2, 0),
('8', 'Cajado de Gelo', 'Cajado mágico que lança gelo.', 'cajado','arma', 4, 2000, 10, 0, 0, 2, 0),
('9', 'Chapéu do Mago', 'Chapéu pontudo que aumenta inteligência.', 'chapeu','chapeu', 1, 600, 0, 0, 0, 1, 0),
('10', 'Chapéu do Sábio', 'Chapéu que aumenta sabedoria.', 'chapeu','chapeu', 1, 700, 0, 0, 0, 0, 1),
('11', 'Espada Bastarda', 'Espada de duas mãos.', 'espada', 'arma', 6, 1800, 10, 3, 0, 0, 0),
('12', 'Armadura de Escamas', 'Armadura feita de escamas de dragão.', 'armadura', 'armadura', 10, 5000, 0, 3, 0, 0, 1),
('13', 'Botas Silenciosas', 'Botas que aumentam furtividade.', 'botas','botas', 1, 1100, 0, 0, 2, 0, 0),
('14', 'Cajado da Vida', 'Cajado que aumenta sabedoria.', 'cajado','arma', 3, 2500, 7, 0, 0, 0, 2),
('15', 'Chapéu do Ladino', 'Chapéu que aumenta destreza.', 'chapeu','chapeu', 1, 550, 0, 0, 1, 0, 0),
('16', 'Espada Curvada', 'Espada exótica e afiada.', 'espada', 'arma', 4, 1300, 7, 1, 1, 0, 0),
('17', 'Armadura Mística', 'Armadura leve que aumenta inteligência.', 'armadura', 'armadura', 5, 2200, 0, 0, 0, 2, 0),
('18', 'Botas do Vento', 'Botas que aumentam velocidade.', 'botas','botas', 1, 1500, 0, 0, 2, 0, 0),
('19', 'Cajado do Trovão', 'Cajado que invoca raios.', 'cajado','arma', 4, 3000, 12, 0, 0, 2, 0),
('20', 'Chapéu do Guerreiro', 'Chapéu que aumenta força.', 'chapeu','chapeu', 1, 650, 0, 1, 0, 0, 0);