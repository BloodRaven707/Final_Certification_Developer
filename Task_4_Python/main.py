import sys


from animal import Dog, Cat, Hamster, Horse, Donkey, Camel
from counter import Counter


animals = [] # Тут должна быть субд


def show_register():

    print(f"\n{'Name':9} {'Birth date':11s} {'Commands'}")
    for animal in animals:
        print(f"{animal.name:8} {animal.birth_date:11s} {animal.commands}")


def register_animal():
    name = input("Enter animal name: ")
    animal_type = input("Enter animal type (dog/cat/hamster/horse/donkey/camel): ")

    if animal_type == "dog":
        animal = Dog(name)
    elif animal_type == "cat":
        animal = Cat(name)
    elif animal_type == "hamster":
        animal = Hamster(name)
    elif animal_type == "horse":
        animal = Horse(name)
    elif animal_type == "donkey":
        animal = Donkey(name)
    elif animal_type == "camel":
        animal = Camel(name)
    else:
        print("Invalid animal type")
        return

    birth_date = input("Enter animal birth_date (2006-01-02): ")
    commands = [c.strip() for c in input("Enter animal commands through comma: ").split(",")]

    with Counter() as counter:
        animals.append(animal)
        counter.add()
        print(f"\n{animal.name} was registered successfully")


def get_commands(animal):
    print(f"\n{animal.name}: {animal.commands}")


def teach_command(animal):
    command = [c.strip() for c in input("Enter animal commands through comma: ").split(",")]
    animal.add_commands(command)
    print(f"\nCommands added to {animal.name} successfully")

    get_commands(animal)


def print_menu():
    print()
    print("Animal Registry")
    print("0. Show AnimalRegister")
    print("1. Register new animal")
    print("2. Get commands for animal")
    print("3. Teach command to animal")
    print("4. Quit")


def main():
    while True:
        print_menu()
        choice = input("Enter choice: ")

        # В Windows 7 последняя поддерживаемая версия 3.8.*
        # Поэтому match case недоступен

        if choice == "0":
            show_register()
        elif choice == "1":
            register_animal()
        elif choice == "2":
            name = input("Enter animal name to get commands: ")
            animal = next((a for a in animals if a.name == name), None)
            if animal:
                get_commands(animal)
            else:
                print("Animal not found")
        elif choice == "3":
            name = input("Enter animal name to teach command: ")
            animal = next((a for a in animals if a.name == name), None)
            if animal:
                teach_command(animal)
            else:
                print("Animal not found")
        elif choice == "4":
            print("Goodbye!")
            sys.exit()
        else:
            print("Invalid choice")


if __name__ == "__main__":
    # Данные для тестирования
    animals.append(Dog("Rex", "2019-04-05", ["Sit", "Stay", "Fetch"]))
    animals.append(Cat("Whiskers", "2018-09-08", ["Meow, Purr"]))
    animals.append(Horse("Spirit", "2015-05-20", ["Gallop, Jump, Move, Stay"]))

    main()
