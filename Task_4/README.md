## Инициализация проекта из папки Task_4
go mod init Task_4
- Создает файл go.mod

## Пробный запуск
cd .\Task_4
go run .\cmd

## Остановить программу
CTRL + C

## Сборка исполняемого файла
cd .\Task_4
go build -o bin/zoo.exe .\cmd
- Создает файл animalregistry.exe в папке bin

## Запуск исполняемого файла
cd .\Task_4
.\bin\animalregistry.exe
