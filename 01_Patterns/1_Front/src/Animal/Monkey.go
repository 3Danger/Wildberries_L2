package Animal

import "time"

type Monkey struct {
	animal
}

func NewMonkey() (e *Monkey) {
	return &Monkey{*newAnimal("\u001B[;32mMonkey\u001B[0m", 4, time.Millisecond*500)}
}
