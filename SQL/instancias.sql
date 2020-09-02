
INSERT INTO usuario VALUES (1, 'matheusbampi@gmail.com', 'Matheus D Bampi', '7751a23fa55170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (2, 'erickpuls@gmail.com', 'Erick Puls', '7751a23s5170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (3, 'jorge.silva@hotmail.com', 'Jorge Da Silva', '7751a23s5170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (4, 'contato@gamestore.com.br', 'Game Store', '7751a23s5170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (5, 'comercial@pontodopapel.com.br', 'Ponto do Papel', '7751a23s5170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (6, 'livrariamagica@gmail.com', 'Livraria Magica', '7751a23s5170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (7, 'brasil@amazon.com', 'Amazon', '7751a23sa3s23ddf13a3ab78e123fe0e99');
INSERT INTO usuario VALUES (8, 'livros@online.com', 'Livros Online', '7751a23fa55170a57e90374df13a3ab78efe0e99');
INSERT INTO usuario VALUES (9, 'kids@email.com', 'Kids Place', '7751a23sa3s23ddf13a3ab78e123fe0e99');

INSERT INTO comprador VALUES (1, '01312340099', TRUE, 1);
INSERT INTO comprador VALUES (2, '90876041232', TRUE, 2);
INSERT INTO comprador VALUES (3, '98800112342', FALSE, 3);

INSERT INTO endereco VALUES (1, 'Marlise', '(54)123882828', 'Brasil', 'RS', 'Farroupilha', 901123222, 'Rua Paim Filho', 223, 'apto 301', 1);
INSERT INTO endereco VALUES (2, 'Matheus Bampi', '(51)234244432', 'Brasil', 'RS', 'Porto Alegre', 912929303, 'Rua Marques do Herval', 512, 'apto 403', 1);
INSERT INTO endereco VALUES (3, 'Jorge Silva', '(51)93414321', 'Brasil', 'RS', 'Porto Alegre', 712839923, 'Rua Dr Timoteo', 11, 'casa A', 3);

INSERT INTO cartao VALUES (1, 'Matheus D Bampi', 1192837384928374673, '2024-01-01 00:00:00', 123, 1);
INSERT INTO cartao VALUES (2, 'Jorge A Silva', 2828373809083746444, '2023-01-01 00:00:00', 099, 3);
INSERT INTO cartao VALUES (3, 'Marlise Bampi', 1919283783746829238, '2021-01-01 00:00:00', 878, 1);

INSERT INTO anunciante VALUES (1, '29384602937465', '(54)87761233', 4);
INSERT INTO anunciante VALUES (2, '91832749274747', '(11)32610099', 5);
INSERT INTO anunciante VALUES (3, '71736389626255', '(51)99887725', 6);
INSERT INTO anunciante VALUES (4, '10910928475744', '(20)33229090', 7);
INSERT INTO anunciante VALUES (5, '92837493830484', '(16)98872626', 8);
INSERT INTO anunciante VALUES (6, '74638263862900', '(54)32664077', 9);

INSERT INTO produto VALUES (1, 'O Morro dos Ventos Uivantes', 'O morro dos ventos uivantes retrata uma trágica historia de amor e obsessão em que os personagens principais são a obstinada e geniosa Catherine Earnshaw e seu irmão adotivo, Heathcliff.', 'https://images-na.ssl-images-amazon.com/images/I/519dh83uTeL._SX345_BO1,204,203,200_.jpg', 'Emily Bronte');
INSERT INTO produto VALUES (2, 'As Crônicas de Nárnia', 'Viagens ao fim do mundo, criaturas fantásticas e batalhas épicas entre o bem e o mal - o que mais um leitor poderia querer de um livro?', 'https://images-na.ssl-images-amazon.com/images/I/51+2QAB7I+L._SX329_BO1,204,203,200_.jpg', 'C. S. Lewis');
INSERT INTO produto VALUES (3, 'O homem mais rico da Babilônia', 'Com mais de dois milhões de exemplares vendidos no mundo todo, O homem mais rico da Babilônia é um clássico sobre como multiplicar riqueza e solucionar problemas financeiros.', 'https://images-na.ssl-images-amazon.com/images/I/41+WJy1N1dL._SX333_BO1,204,203,200_.jpg', 'George S. Clason');
INSERT INTO produto VALUES (4, 'Fifa 20 - PS4', 'Fifa 20 para Playstation 4.', 'https://images-na.ssl-images-amazon.com/images/I/810m7X9943L._AC_SX679_.jpg', 'EA Sports');
INSERT INTO produto VALUES (5, 'Dualshock 4', 'Controle de Playstation 4.', 'https://images-na.ssl-images-amazon.com/images/I/51Ieq0twp%2BL._AC_SL1000_.jpg', 'Sony');
INSERT INTO produto VALUES (6, 'Boneco', 'Boneco do Batman Articulado', 'https://images-na.ssl-images-amazon.com/images/I/brinquedo.jpg', 'DC');
INSERT INTO produto VALUES (7, 'One Piece', 'One Piece vol 5', 'https://images-na.ssl-images-amazon.com/images/I/manga.jpg', 'Shonen Jump');
INSERT INTO produto VALUES (8, 'Como é bom compartilhar', 'Compartilhar a refeição e os brinquedos com os colegas ajuda a construir uma convivência cheia de harmonia. Aprenda com os animais da selva sobre a importância de dividir!', 'https://images-na.ssl-images-amazon.com/images/I/51j-gE7hSIL._SX435_BO1,204,203,200_.jpg', 'Fisher Price');

INSERT INTO categoria VALUES (1, 'Livros', 'Livros dos mais variados temas.', NULL);
INSERT INTO categoria VALUES (5, 'Negócios', 'Livros para incentivar seu lado empreendedor.', 1);
INSERT INTO categoria VALUES (6, 'Romance', 'Histórias que satisfazem a mente.', 1);
INSERT INTO categoria VALUES (2, 'Games', 'Tudo que você precisa para uma diversão completa.', NULL);
INSERT INTO categoria VALUES (3, 'Acessórios', 'Controles e suportes para seu videogame.', 2);
INSERT INTO categoria VALUES (4, 'Jogos', 'Os melhores jogos.', 2);
INSERT INTO categoria VALUES (7, 'Playstation 4', 'Tudo para o seu PS4.', 2);
INSERT INTO categoria VALUES (8, 'Livro Infantil', 'Conto de fadas', 2);
INSERT INTO categoria VALUES (9, 'Mangas', 'Literatura Japonesa', 2);
INSERT INTO categoria VALUES (10, 'Brinquedos', 'Brinquedo infantil', 2);

INSERT INTO produto_categoria VALUES (1, 6);
INSERT INTO produto_categoria VALUES (2, 6);
INSERT INTO produto_categoria VALUES (3, 5);
INSERT INTO produto_categoria VALUES (4, 4);
INSERT INTO produto_categoria VALUES (4, 7);
INSERT INTO produto_categoria VALUES (5, 3);
INSERT INTO produto_categoria VALUES (5, 7);
INSERT INTO produto_categoria VALUES (2, 8);
INSERT INTO produto_categoria VALUES (7, 9);
INSERT INTO produto_categoria VALUES (6, 10);
INSERT INTO produto_categoria VALUES (8, 8);

INSERT INTO cupom VALUES (1, 5, '2020-06-22 19:10:00');
INSERT INTO cupom VALUES (2, 10, '2021-12-25 12:30:00');
INSERT INTO cupom VALUES (3, 15, '2020-01-10 08:00:00');
INSERT INTO cupom VALUES (4, 30,'2022-01-01 23:59:30');

INSERT INTO avaliacao VALUES (3, 'Bom livro! Tive diversos insights pro meu restaurante, mas poderia ser mais objetivo..', '2019-01-21 22:14:29', 3, 3);
INSERT INTO avaliacao VALUES (5, 'Livro muito bom. Fácil leitura e com muitos ensinamentos.', '2017-06-13 23:04:55', 3, 1);
INSERT INTO avaliacao VALUES (4, 'Bom jogo, mas faltaram os times brasileiros :(', '2017-06-13 23:04:55', 4, 1);

INSERT INTO produto_cupom VALUES (1, 1);
INSERT INTO produto_cupom VALUES (2, 1);
INSERT INTO produto_cupom VALUES (3, 3);

INSERT INTO listaDesejos VALUES (3, 1);
INSERT INTO listaDesejos VALUES (4, 3);
INSERT INTO listaDesejos VALUES (5, 3);

INSERT INTO produtofornecido VALUES (1, 15.45, 12, 5.00, FALSE, 2, 1);
INSERT INTO produtofornecido VALUES (2, 19.00, 6, 11.90, FALSE, 3, 1);
INSERT INTO produtofornecido VALUES (3, 12.50, 8, 5.10, FALSE, 2, 2);
INSERT INTO produtofornecido VALUES (4, 18.80, 20, 8.00, FALSE, 2, 3);
INSERT INTO produtofornecido VALUES (5, 17.45, 1, 9.50, FALSE, 3, 3);
INSERT INTO produtofornecido VALUES (6, 149.90, 3, 10.00, FALSE, 1, 4);
INSERT INTO produtofornecido VALUES (7, 210.00, 54, 15.00, FALSE, 1, 5);
INSERT INTO produtofornecido VALUES (8, 180.00, 8, 10.00, TRUE, 4, 5);
INSERT INTO produtofornecido VALUES (9, 15.00, 8, 10.00, TRUE, 4, 3);
INSERT INTO produtofornecido VALUES (10, 15.45, 12, 5.00, FALSE, 2, 6);
INSERT INTO produtofornecido VALUES (11, 19.00, 6, 11.90, FALSE, 3, 7);
INSERT INTO produtofornecido VALUES (12, 22.00, 8, 15.90, FALSE, 3, 6);
INSERT INTO produtofornecido VALUES (13, 42.00, 14, 5.90, FALSE, 5, 8);
INSERT INTO produtofornecido VALUES (14, 39.50, 2, 10.50, TRUE, 6, 8);

INSERT INTO carrinho VALUES (2, 15.00, 1, 1);
INSERT INTO carrinho VALUES (1, 180.00, 8, 1);
INSERT INTO carrinho VALUES (1, 17.50, 5, 3);

INSERT INTO pedido VALUES (1, 165.90, 18.00, '2019-01-11 19:17:07', 'EM TRANSPORTE', FALSE, NULL, 1, 1, 1);
INSERT INTO pedido VALUES (2, 12.50, 5.10, '2020-08-10 15:50:33', 'CONFIRMADO', FALSE, 1, 3, 3, 2);
INSERT INTO pedido VALUES (3, 14.50, 0.00, '2018-11-03 08:41:46', 'ENTREGUE', TRUE, 3, 1, 2, 1);

INSERT INTO produtofornecido_pedido VALUES (6, 1, 1, 149.90);
INSERT INTO produtofornecido_pedido VALUES (4, 1, 3, 16.00);
INSERT INTO produtofornecido_pedido VALUES (3, 2, 1, 12.50);
INSERT INTO produtofornecido_pedido VALUES (9, 3, 1, 14.50);
