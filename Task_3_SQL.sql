-- 7. В MySQL создать базу данных "Друзья человека"
-- Удаление базы данных, для повтороного использования скрипта
DROP DATABASE IF EXISTS human_friends;

CREATE DATABASE human_friends;
-- CREATE DATABASE IF NOT EXISTS human_friends;

-- Подключение к базе данных
USE human_friends;

-- 8. Создать таблицы с иерархией из диаграммы в БД
CREATE TABLE animals (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	animal_type VARCHAR(20)
);
INSERT INTO animals (animal_type)
VALUES
	('Pack_animal'),
	('Pet');

-- Вывод таблицы animals, содержащей типы животных
-- SELECT * FROM animals;


CREATE TABLE pack_animals (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	animal_subtype VARCHAR(20)
);
INSERT INTO pack_animals (animal_subtype)
VALUES
	('Horse'),
	('Donkey'),
	('Camel');

-- Вывод таблицы pack_animals, содержащей подтипы выючных животных
-- SELECT * FROM pack_animals;

CREATE TABLE pets (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	animal_subtype VARCHAR(20)
);
INSERT INTO pets (animal_subtype)
VALUES
	('Cat'),
	('Dog'),
	('Hamster');

-- Вывод таблицы pets, содержащей подтипы домашних животных
-- SELECT * FROM pets;

CREATE TABLE cats (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);

CREATE TABLE dogs (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);

CREATE TABLE hamsters (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);

CREATE TABLE horses (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);

CREATE TABLE camels (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);

CREATE TABLE donkeys (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	name VARCHAR(20),
	birthday DATE,
	commands JSON
);


-- 9. Заполнить низкоуровневые таблицы именами(животных), командами которые они выполняют и датами рождения
INSERT INTO cats (name, birthday, commands)
VALUES
   ('Мурзик', '2020-03-15', '["сидеть", "лежать", "дай лапу"]'),
   ('Пушок', '2021-05-01', '["мяукать", "играть"]'),
   ('Барсик', '2022-01-03', '["царапать", "прыгать"]');

-- Вывод таблицы cats
-- SELECT * FROM cats;

INSERT INTO dogs (name, birthday, commands)
VALUES
   ('Шарик', '2019-11-12', '["сидеть", "лежать", "голос"]'),
   ('Бобик', '2020-07-08', '["апорт", "фас"]'),
   ('Рекс', '2021-02-15', '["ко мне", "рядом"]');

-- Вывод таблицы dogs
-- SELECT * FROM dogs;

INSERT INTO hamsters (name, birthday, commands)
VALUES
   ('Хомячок', '2020-12-01', '[]'),
   ('Зёрнышко', '2021-06-11', '[]'),
   ('Пухляк', '2022-04-20', '[]');

-- Вывод таблицы hamsters
-- SELECT * FROM hamsters;

INSERT INTO horses (name, birthday, commands)
VALUES
   ('Буцефал', '2018-05-16', '["направо", "налево"]'),
   ('Балкан', '2019-11-29', '["стоп"]'),
   ('Дикарь', '2020-02-11', '["шагом"]');

-- Вывод таблицы horses
-- SELECT * FROM horses;

INSERT INTO camels (name, birthday, commands)
VALUES
   ('Копыта', '2017-09-12', '["иди", "пить"]'),
   ('Горбун', '2018-06-18', '["стой"]'),
   ('Дюна', '2020-08-15', '["сидеть", "вставать"]');

-- Вывод таблицы camels
-- SELECT * FROM camels;

INSERT INTO donkeys (name, birthday, commands)
VALUES
   ('Иа', '2016-01-05', '["назад", "стоп"]'),
   ('Рамзес', '2017-11-15', '["иди"]'),
   ('Ослик', '2019-03-19', '["назад"]');

-- Вывод таблицы donkeys
-- SELECT * FROM donkeys;

-- 10. Удалите таблицу верблюды.
DROP TABLE camels;

-- 10.1. Объединить таблицы лошади, и ослы в одну таблицу
-- Добавил поле с подтипом животного
CREATE TABLE equines AS
    SELECT h.id, pan.animal_subtype, h.name, h.birthday, h.commands
    FROM horses AS h JOIN pack_animals AS pan ON pan.animal_subtype = 'Horse'

    UNION ALL
    SELECT d.id, pan.animal_subtype, d.name, d.birthday, d.commands
    FROM donkeys AS d JOIN pack_animals AS pan ON pan.animal_subtype = 'Donkey';

-- Вывод объединенной таблицы лощади и ослы
-- SELECT * FROM equines;


-- 11. Создать новую таблицу "молодые животные" в которую попадут все животные старше 1 года, но младше 3 лет
-- В отдельном столбце с точностью до месяца подсчитать возраст животных в новой таблице
-- Как и в прошлой таблице добавил поле с подтипом животного (другим способом)
CREATE TABLE young_animals AS
    SELECT a.id, a.animal_subtype, a.name, a.commands, TIMESTAMPDIFF(MONTH, a.birthday, CURDATE()) AS age_months
    FROM (
        SELECT id, 'cat' AS animal_subtype, name, birthday, commands FROM cats
        UNION ALL
        SELECT id, 'dog' AS animal_subtype , name, birthday, commands FROM dogs
        UNION ALL
        SELECT id, 'hamster' AS animal_subtype, name, birthday, commands FROM hamsters
        UNION ALL
        SELECT id, 'horse' AS animal_subtype, name, birthday, commands FROM horses
        UNION ALL
        SELECT id, 'donkey' AS animal_subtype, name, birthday, commands FROM donkeys
        -- UNION ALL
        -- SELECT id, 'camel' AS animal_subtype, name, birthday, commands FROM camels
    ) AS a
    WHERE a.birthday BETWEEN DATE_SUB(CURDATE(), INTERVAL 3 YEAR) AND DATE_SUB(CURDATE(), INTERVAL 1 YEAR);

-- Вывод объединенной таблицы "Молодые животные"
-- SELECT * FROM young_animals;


-- 12. Объединить все таблицы в одну, при этом сохраняя поля, указывающие на прошлую принадлежность к старым таблицам.
CREATE TABLE all_animals AS
    SELECT c.id, 'cat' AS animal_subtype, c.name, c.birthday, c.commands, a.animal_type
    FROM cats AS c JOIN animals AS a ON a.animal_type = 'Pet'
    UNION ALL
    SELECT d.id, 'dog' AS animal_subtype, d.name, d.birthday, d.commands, a.animal_type
    FROM dogs AS d JOIN animals AS a ON a.animal_type = 'Pet'
    UNION ALL
    SELECT hr.id, 'hamsters' AS animal_subtype, hr.name, hr.birthday, hr.commands, a.animal_type
    FROM hamsters AS hr JOIN animals AS a ON a.animal_type = 'Pet'
    UNION ALL
    SELECT h.id, 'horse' AS animal_subtype, h.name, h.birthday, h.commands, a.animal_type
    FROM horses AS h JOIN animals AS a ON a.animal_type = 'Pack_animal'
    UNION ALL
    SELECT d.id, 'donkey' AS animal_subtype, d.name, d.birthday, d.commands, a.animal_type
    FROM donkeys AS d JOIN animals AS a ON a.animal_type = 'Pack_animal'
    -- SELECT d.id, 'camel' AS animal_subtype, d.name, d.birthday, d.commands, a.animal_type
    -- FROM camels AS d JOIN animals AS a ON a.animal_type = 'Pack_animal'
;

-- Вывод объединенной таблицы "Все животные"
SELECT * FROM all_animals;
