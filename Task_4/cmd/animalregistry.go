package main


import "fmt"
import "strings"
import "time"


// Описывает реестр домашних животных
type AnimalRegistry struct {
	animals map[string]Animal
}

// Создает новый реестр домашних животных
func NewAnimalRegistry() AnimalRegistry {
	return AnimalRegistry{
		animals: make(map[string]Animal),
	}
}

func (r *AnimalRegistry) Show() {
    fmt.Printf("%-10v: %-7v %-11v %s  [%v]\n", "Name", "Type", "SubType", "BirthDate", "Commands")
    for name, a := range r.animals {
        fmt.Printf("%-10v: %-7v %-11v %s  [%v]\n", name, a.Type, a.SubType, a.BirthDate.Format(("2006-01-02")), strings.Join(a.Commands, ", "))
  }
}

// Получить животное из реестра
func (r *AnimalRegistry) GetForName(name string) *Animal {
    if animal, ok := r.animals[name]; ok {
        return &animal
    }
    return nil
}

// Добавляет новое животное в реестр
func (r *AnimalRegistry) AddAnimal(animal *Animal) {
    // За нулевую дату рождения принял текущее время...
    if animal.Name != "" || animal.BirthDate == time.Now() || animal.Commands != nil {
        animalConster.add()
        fmt.Println(animalConster.value)
    } else {
        panic("Объект Animal создан не правильно")
    }
	r.animals[animal.Name] = *animal
}

// Удаляет животное из реестра
//func (r *AnimalRegistry) DelAnimal(animal Animal) Animal {
//    return ...
//}

// Возвращает тип, к которому принадлежит животное
func (r *AnimalRegistry) GetType(name string) string {
	a := r.GetForName(name)
    if a != nil {
		return a.Type
	}
	return "Unknown"
}

// Выводит список команд, которые выполняет животное
func (r *AnimalRegistry) ShowCommands(name string) {
	a := r.GetForName(name)
    if a != nil {
        a.ShowCommands()
    } else {
        fmt.Println("Animal not found.")
    }
}

// Добавляет новые команды для животного
func (r *AnimalRegistry) TeachAnimal(name string, newCommands []string) {
	animal := r.GetForName(name)
    if animal != nil {
        animal.Commands = append(animal.Commands, newCommands...)
        r.animals[name] = *animal
        fmt.Printf("%s has learned new commands: %v\n", name, newCommands)
    } else {
        fmt.Println("Animal not found.")
    }
}