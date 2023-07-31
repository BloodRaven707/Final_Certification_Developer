package main


import "sync"


// Работает с реестром животных, считает добавленные в него методом AddAnimal объекты типа Animal
type Counter struct {
    value int
    mu    sync.Mutex
    lock  bool
}

func NewCounter() Counter {
    return Counter{
		value: 0,
	}
}

/*func (c *Counter) get() int {
    return c.value
}*/

func (c *Counter) add() {
    if !c.lock {
        panic("Объект Counter не был закрыт")
    }

    // Блокируем для других потоков = открываем = начинаем работать с объектом
    c.mu.Lock()
    c.lock = !c.lock

    c.value++

    // Разблокируем для других потоков = закрывем = прекращаем работать с объектом
    c.mu.Unlock()
    c.lock = !c.lock
    // c.lock == true // В данный момент объект закрыт, если не закроется при следующем вызове будет паника(Ошибка в Go прграмме)
}
