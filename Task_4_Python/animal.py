class Animal:
    def __init__(self, name, birth_date, commands):
        self.name = name
        self.birth_date = birth_date
        self.commands = commands

    def get_commands(self):
        print(f"[{', '.join(self.commands)}]")

    def add_commands(self, new_commands):
        self.commands.extend(new_commands)


class Pet(Animal):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class PackAnimal(Animal):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Dog(Pet):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Cat(Pet):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Hamster(Pet):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Horse(PackAnimal):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Camel(PackAnimal):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)

class Donkey(PackAnimal):
    def __init__(self, name, birth_date, commands):
        super().__init__(name, birth_date, commands)
