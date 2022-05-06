package Animal

import "time"

type Elephant struct {
	animal
}

func NewElephant() (e *Elephant) {
	return &Elephant{*newAnimal("\u001B[1;35mElephant\u001B[0m", 6, time.Second)}
}
