insert into users(name, nick, email, password)
values
("Test 2", "test2", "test1@test.com", "$2a$10$uvIz9/DsYs.aJUi1uWi1fuycjxiIeeoTVnXwXHOCM/PbaKbqMJz92"),
("Test 1", "test1", "test2@test.com", "$2a$10$uvIz9/DsYs.aJUi1uWi1fuycjxiIeeoTVnXwXHOCM/PbaKbqMJz92"),
("Test 3", "test3", "test3@test.com", "$2a$10$uvIz9/DsYs.aJUi1uWi1fuycjxiIeeoTVnXwXHOCM/PbaKbqMJz92");

insert into followers(userID, followerID)
values
(1, 2),
(2, 3),
(3, 2);

insert into posts(title, content, author_id)
values
("Titulo 1", "Conteudo 1",1),
("Titulo 2", "Conteudo 2",2),
("Titulo 3", "Conteudo 3",3);