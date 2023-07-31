class Counter:
    def __init__(self):
        self.count = 0

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, traceback):
        if self.count != 0:
            raise Exception("Counter not closed correctly!")

    def add(self):
        self.count += 1
