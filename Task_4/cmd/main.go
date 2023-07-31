package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//import "net/http"

// Глобальная осласть видимости
var animalRegistry = NewAnimalRegistry()
var animalConster  = NewCounter()


func init() {
    // Вообще чтобы имело смысл нужна асинхронность... или несколько потоков...
    animalConster.lock = true

    // Заводить новых животных можно разнимы способами...
    dog := &Animal{"Dog", "Pet", "Rex", time.Date(2018, time.January, 10, 0, 0, 0, 0, time.UTC), []string{"Sit", "Stay", "Fetch"}}
    cat := &Animal{"Cat", "Pet", "Whiskers", MyTimeParse("2019-04-05"), strings.Split("Meow, Purr", ", ")}
    // Конструктор явно проще...
    horse := NewHorse( "Spirit", "2015-05-20", "Gallop, Jump, Move, Stay" )

    // Добавляем животных в реестр
    animalRegistry.AddAnimal(dog)
    animalRegistry.AddAnimal(cat)
    animalRegistry.AddAnimal(horse)

    // Создаем и сразе Добавляем животных в реестр
    animalRegistry.AddAnimal(NewDog("Mars", "2018-10-04", "Sit, Fas, Plase, To me"))
    animalRegistry.AddAnimal(NewCat("Grey", "2012-08-20", "Meat")) // Мой кот... других команд не знает...
    animalRegistry.AddAnimal(NewDonkey("Sabi", "2021-01-12", "Move, Stay"))
    animalRegistry.AddAnimal(NewDonkey("Sati", "2020-11-21", "Move, Stay"))
    animalRegistry.AddAnimal(NewDonkey("Ban", "2016-04-23", "Move, Stay"))
    animalRegistry.AddAnimal(NewCamel("Bak", "2016-04-23", "Move, Stay"))
    animalRegistry.AddAnimal(NewCamel("Gak", "2017-07-27", "Move, Stay"))
}


func main() {
    //fmt.Println("Animal Class of Rex:", animalRegistry.GetType("Rex"))
    //fmt.Println("Animal Class of Whiskers:", animalRegistry.GetType("Whiskers"))
    //fmt.Println("Animal Class of Spirit:", animalRegistry.GetType("Spirit"))

    // Выводим список команд для животных
    //animalRegistry.ShowCommands("Rex")
    //animalRegistry.ShowCommands("Whiskers")
    //animalRegistry.ShowCommands("Spirit")

    for {
        // Выводим информацию о животных

        fmt.Printf("\nMenu:\n")
        fmt.Println("0. Show animals registry list")
        fmt.Println("1. Add new Animal")
        fmt.Println("2. Show animal commands")
        fmt.Println("3. Teach animal new commands")
        fmt.Println("4. Exit")

        fmt.Printf("\nCommand: ")

        var input int
        fmt.Scan(&input)
        fmt.Println()

        switch input {
        case 0: // Выводим информацию о животных в реестре // +
            animalRegistry.Show()

        case 1: // Заводим новое животное // +
            var name string
            for {
                fmt.Println("Attention:")
                fmt.Println("- Animal name должно быть уникальным, давать одинаковые имена нельзя")
                fmt.Println("- Animal name должно быть не короче 3-х символов и не более 20")
                fmt.Println()
                fmt.Print("Enter animal name: ")
                fmt.Scan(&name)
                fmt.Println()

                if name == "" {
                    fmt.Printf("Имя необходимо ввести в обязательном порядке...\n\n")
                    continue
                } else if len(name) < 3 || len(name) > 20 {
                    fmt.Printf("Имя животного не может быть короче 3-х символов и не более 20...\n\n")
                    continue
                }
                if animalRegistry.GetForName("name") != nil {
                    fmt.Printf("К сожалению животное с таким именем уже усть в реестре...\n\n")
                    continue
                }
                break
            }

            var birthDate string
            for {
                fmt.Print("Enter birthdate (ГГГГ-ММ-ДД): ")
                fmt.Scan(&birthDate)
                fmt.Println()

                if birthDate == "" {
                    fmt.Printf("Дату рождения необходимо ввести в обязательном порядке...\n\n")
                    continue
                }
                _, err := time.Parse("2006-01-02", birthDate)
                if err != nil {
                    fmt.Printf("Не верный формат даты рождения, попробуйте снова...\n\n")
                    continue
                }
                break
            }

            var commands string
            for {
                fmt.Print("Enter commands, через запятую: ")
                scanner := bufio.NewScanner(os.Stdin)
                scanner.Scan()
                commands = scanner.Text()

                fmt.Println()
                break
            }

            var Type int
            for flag := false; flag == false; {
                fmt.Println("Enter animal type code(от 1 до 6): ")
                fmt.Println("1. Cat")
                fmt.Println("2. Dog")
                fmt.Println("3. Homster")
                fmt.Println("4. Horse")
                fmt.Println("5. Donkey")
                fmt.Println("6. Camal")
                fmt.Print("Animal type code: ")
                fmt.Scan(&Type)
                fmt.Println()
                switch Type {
                    case 1:
                        animalRegistry.AddAnimal(NewCat(name, birthDate, commands))
                        flag = !flag // break - Выходит из switch а не из for
                    case 2:
                        animalRegistry.AddAnimal(NewDog(name, birthDate, commands))
                        flag = !flag
                    case 3:
                        animalRegistry.AddAnimal(NewHamster(name, birthDate, commands))
                        flag = !flag
                    case 4:
                        animalRegistry.AddAnimal(NewHorse(name, birthDate, commands))
                        flag = !flag
                    case 5:
                        animalRegistry.AddAnimal(NewDonkey(name, birthDate, commands))
                        flag = !flag
                    case 6:
                        animalRegistry.AddAnimal(NewCamel(name, birthDate, commands))
                        flag = !flag
                    default:
                        fmt.Printf("%v\n", Type)
                        fmt.Printf("Вы ввели не известный код животного, попробуйте снова\n\n")
                        continue
                }
            }
            fmt.Println("Animal added")

        case 2: // Выводим список команд животного
            var name string
            fmt.Print("Enter animal name: ")
            fmt.Scan(&name)
            fmt.Println()

            animalRegistry.ShowCommands(name)

        case 3: // обучение команде
            var name string
            fmt.Print("Enter animal name: ")
            fmt.Scan(&name)
            fmt.Println()

            animal := animalRegistry.GetForName(name)
            if animal != nil {
                var commands string
                animal.ShowCommands()

                fmt.Print("Enter new commands: ")
                scanner := bufio.NewScanner(os.Stdin)
                scanner.Scan()
                commands = scanner.Text()
                fmt.Println()

                animal.AddCommand(commands)
                fmt.Printf("New commands for %v added!\n", name)
                animalRegistry.animals[name] = *animal
                animal.ShowCommands()
            } else {
                fmt.Println("Такого животного нет")
            }

        case 4: // выход
            fmt.Println(animalConster.value)
            return
        }
    }
}


// Вызывает time.Parse внутри, обрабытавает ошибку и отдает результат
func MyTimeParse(value string) time.Time {
    t, err := time.Parse("2006-01-02", value)
    if err != nil { panic(err) }
	return t
}
