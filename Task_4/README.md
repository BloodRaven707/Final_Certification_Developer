## Инициализация проекта из папки Task_4
cd .\Task_4
go mod init Task_4
- Создает файл go.mod в папке \Task_4

<br>

## Пробный запуск
cd .\Task_4
go run .\cmd

<br>

## Остановить программу
CTRL + C

<br>

## Сборка исполняемого файла
cd .\Task_4
go build -o bin/zoo.exe .\cmd
- Создает файл animalregistry.exe в папке bin

<br>

## Запуск исполняемого файла
cd .\Task_4
.\bin\animalregistry.exe
