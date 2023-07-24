# Итоговая контрольная работа по специализации "Программист"

## Информация о проекте
- Необходимо организовать систему учета для питомника в котором живут домашние и вьючные животные.

## Как сдавать проект
- Для сдачи проекта необходимо создать отдельный общедоступный репозиторий.
- Разработку вести в этом репозитории, использовать пул реквесты на изменения.
- Программа должна запускаться и работать, ошибок при выполнении программы быть не должно.
- Программа, может использоваться в различных системах, поэтому необходимо разработать класс в виде конструктора.

<br>

## Задание
### Блок по работе с терминалом в OS Linux
- Операции с файлами:
    1. Используя команду cat создать файл "Домашние животные", с данными: собаки, кошки, хомяки.
    2. Создать файл "Вьючные животные", с данными: лошади, верблюды, ослы.
    3. Объединить созданные файлы.
    4. Вывести содержимое созданного файла.
    5. Переименовать файл в "Друзья человека".
    6. Создать директорию, переместить в нее файл.
- Репозотории и работа с пакетами
    1. Подключить дополнительный репозиторий MySQL.
    2. Установить любой пакет из подключенного репозитория.
    3. Установить и удалить deb-пакет с помощью dpkg.
- Выложить историю команд из терминала OS Linux.

<br>

### Блок по работе с Диаграммами
- Нарисовать диаграмму:
    1. Родительский класс.
    2. Подклассы домашние животные и вьючные животные.
    3. Подклассы для домашних животных: собаки, кошки, хомяки.
    4. Подклассы для вьючных животных: лошади, верблюды, ослы.

<br>

### Блок по работе с SQL в СУБД MySQL
- Создать базу данных "Друзья человека" // "human_friends"
    1. Создать таблицы "pets", "dogs", "cats", "hamsters", "pack_animals", "horses", "camels", "donkeys" с иерархией из диаграммы.
    2. Низкоуровневые таблицы должны иметь следующие поля:
        - имена животных.
        - команды которые животные выполняют.
        - даты рождения животных.
    3. Наполните низкоуровневые таблицы данными.
    4. Удалите таблицу верблюды.
    5. Объедините таблицы "лошади" и "ослы".
    6. Создайте таблицу "молодые животные". // "young_animals"
        - В отдельном столбце с точностью до месяца должен храниться возраст животных.
    7. Поместите в нее всех животных старше 1 года и младше 3 лет.
    8. Объединить все таблицы в одну, при этом сохраняя поля, указывающие на прошлую принадлежность к старым таблицам.

<br>

### Блок "Программирование"
- Создайте класс с Инкапсуляцией методов и наследованием по диаграмме.
- Напишите программу, имитирующую работу реестра домашних животных, со следующим функционалом:
    1. Завести новое животное.
    2. Определять животное в правильный класс.
    3. Увидеть список команд, которое выполняет животное.
    4. Обучить животное новым командам.
    5. Реализовать навигацию по меню.
- Создайте класс-счетчик, со следующим функционалом:
    1. Метод add(), увеличивающий значение внутренней int переменной на 1 при нажатии "Завести новое животное".
    2. Объект типа счетчик должен уметь работать в блоке try-with-resources.
    3. Должно срабатывать исключение, если работа с объектом типа счетчик была не в ресурсном try и/или ресурс остался открыт.
    4. Значение считать в ресурсе try, если при заведения животного заполнены все поля.

<br><br>

## Команды для GIT
### Создание локального репозитория
> git init
- В Windows по умолчанию создает ветку master

<br>

### Переименование главной ветки в main для GitHub
> git branch -M main

<br>

### Создание локального репозитория c главной веткой main
> git init --initial-branch=main
- Указывает имя создаваемой главной ветки для репозитория

<br>

### Добавление всех файлов в отслеживаемые
> git add .

<br>

### Создание коммита
> git commit -m "Initial commit"

<br>

### Подключение к удаленному репозиторию
> git remote add origin https://github.com/BloodRaven707/Final_Certification_Developer.git

<br>

### Загрузка в удаленный репозиторий
> git push -u origin main

<br><br>

## Команды Linux
### Открытие файла с возможностью ввода строк
> cat > "Домашние животние"
- Создает файл если его не существует

### Ctrl + C
- Сохранить и выйти

<br>

### Объединить файлы в один
> cat "Домашние животние" "Вьючные животние" > "Human friends"

<br>

### Вывод в терминал
> cat "Human friends"

<br>

### Перемещение (переименование)
> mv "Human friends" "Друзья человека"

<br>

### Создание дирректории
> mkdir animals

<br>

### Перемещение
mv "Друзья человека" ./animals/

<br>

### Удаление папки со всеми вложенныеми объектами
> raven@raven-vm:~$ rm -rf animals/

<br>

### Удаление файла
> raven@raven-vm:~$ rm "Вьючные животние"

<br>

### Скачиваем пакет c удаленного хранилища
> raven@raven-vm:~$ sudo wget https://dev.mysql.com/get/mysql-apt-config_0.8.25-1_all.deb

<br>

### Устанавливаем пакет с конфурациями с помощью dpkg
> sudo dpkg -i mysql-apt-config_0.8.25-1_all.deb

<br>

### Обновляем зависиости
> sudo apt update

<br>

### Устанавливаем пакет
> sudo apt install mysql-client -y

<br>

### Удаляем пакет
> sudo apt remove mysql-client -y

<br>

### Установка одной командой
> raven@raven-vm:~$ sudo dpkg -i mysql-apt-config_0.8.25-1_all.deb && sudo apt update && sudo apt install mysql-client

<br>

### Удаляем пакет через dpkg
> sudo dpkg -r mysql-client

<br><br>

## Диаграмма Task_2_Chart.xml
![Диаграмма Task_2_Chart.xml](./Task_2_Chart.png)

<br><br>

## Команды SQL
### Создание базы данных
```
CREATE DATABASE IF NOT EXISTS human_friends;
```
<br>

### Удаление базы данных
```
DROP DATABASE IF EXISTS human_friends;
```
<br>

### Подключение к базе данных
```
USE human_friends;
```
<br>

### Создание таблицы
```
CREATE TABLE animals (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	animal_type VARCHAR(20)
);
```
<br>

### Удаление таблицы
```
DROP TABLE camels;
```

<br>

### Вставка данных
```
INSERT INTO animals (animal_type)
VALUES
	('Pack_animal'),
	('Pet');
```
<br>

### Вывод всех полей таблицы
```
SELECT * FROM animals;
```
<br>

### Объединение таблицы и использование JOIN
```
CREATE TABLE equines AS
  SELECT h.id, pan.animal_subtype, h.name, h.birthday, h.commands
  FROM horses AS h JOIN pack_animals AS pan ON pan.animal_subtype = 'Horse'

  UNION ALL
  SELECT d.id, pan.animal_subtype, d.name, d.birthday, d.commands
  FROM donkeys AS d JOIN pack_animals AS pan ON pan.animal_subtype = 'Donkey';
```
<br>

### Объединение таблицы, использование JOIN, встроенные функции, работа с временными интервалами и условия для выборки данных
```
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
    --    UNION ALL
    --    SELECT id, 'camel' AS animal_subtype, name, birthday, commands FROM camels
    ) AS a
    WHERE a.birthday BETWEEN DATE_SUB(CURDATE(), INTERVAL 3 YEAR) AND DATE_SUB(CURDATE(), INTERVAL 1 YEAR);
```

<br>

## Результат работы Task_3_SQL.sql
![Результат работы Task_3_SQL](./Task_3_SQL.png)
