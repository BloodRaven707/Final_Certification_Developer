package main


import "fmt"
import "strings"
import "time"


// Описывает базовую структуру для всех животных
type Animal struct {
    Type      string
    SubType   string
    Name      string
    BirthDate time.Time
    Commands  []string
}

// Создает животное типа Animal
//
// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewAnimal(name string, birthDate string, commands string) *Animal {
    t, err := time.Parse("2006-01-02", birthDate)
    if err != nil { panic(err) }
    return &Animal{
	  Type:	     "Animal",
	  SubType:   "Unknown",
	  Name:      name,
	  BirthDate: t,
	  Commands:  strings.Split(commands, ", "),
	}
}

/*func (a *Animal) GetName() string {
    return a.Name
}*/
/*func (a *Animal) GetBirthDate() time.Time {
    return a.BirthDate
}*/
/*func (a *Animal) GetCommands() []string {
    return a.Commands
}*/

// Добавляет команду, существующему в реестре Животному
func (a *Animal) AddCommand(newCommands string) {
    a.Commands = append(a.Commands, newCommands)
}

func (a *Animal) ShowCommands() {
    fmt.Printf("Commands for %s: [%s]\n", a.Name, strings.Join(a.Commands, ", "))
}

// RenameAnimal - переименовать
// AddAnimalCommand - добавить команду
// AddAnimalCommands - добавить команды

// Описывает домашних животных, которые являются потомками Animal
type Pet struct {
    Animal
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewPet(name string, birthDate string, commands string) *Animal {
    a := NewAnimal(name, birthDate, commands)
    a.Type, a.SubType = "Pet", "Pet"
    return a
}

// Описывает вьючных животных, которые являются потомками Animal
type PackAnimal struct {
    Animal
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewPackAnimal(name string, birthDate string, commands string) *Animal {
    a := NewAnimal(name, birthDate, commands)
    a.Type, a.SubType = "PackAnimal", "PackAnimal"
    return a
}

// Описывает собаку, которая является потомком Pet
type Dog struct {
    Pet
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewDog(name string, birthDate string, commands string) *Animal {
    a := NewPet(name, birthDate, commands)
    a.Type = "Dog"
    return a
}
// func NewDog(name string, birthDate time.Time, commands []string) *Dog {
//     return &Dog{
//         Animal: Animal{name, birthDate, commands},
//     }
// }

// Описывает кошку, которая является потомком Pet
type Cat struct {
    Pet
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewCat(name string, birthDate string, commands string) *Animal {
    a := NewPet(name, birthDate, commands)
    a.Type = "Cat"
    return a
}

// Описывает хомяка, который является потомком Pet
type Hamster struct {
    Pet
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewHamster(name string, birthDate string, commands string) *Animal {
    a := NewPet(name, birthDate, commands)
    a.Type = "Hamster"
    return a
}

// Описывает лошадь, которая является потомком PackAnimal
type Horse struct {
    PackAnimal
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewHorse(name string, birthDate string, commands string) *Animal {
    a := NewPackAnimal(name, birthDate, commands)
    a.Type = "Horse"
    return a
}

// Описывает верблюда, который является потомком PackAnimal
type Camel struct {
    PackAnimal
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewCamel(name string, birthDate string, commands string) *Animal {
    a := NewPackAnimal(name, birthDate, commands)
    a.Type = "Camel"
    return a
}

// Описывает осла, который является потомком PackAnimal
type Donkey struct {
    PackAnimal
}

// Parameters:
//     birthDate: string, format is "2006-01-02"
//     commands : string, format is "command, command, ..."
func NewDonkey(name string, birthDate string, commands string) *Animal {
    a := NewPackAnimal(name, birthDate, commands)
    a.Type = "Donkey"
    return a
}
