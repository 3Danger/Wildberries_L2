package Animal

import "time"

type Bull struct {
	animal
}

func NewBull() (b *Bull) {
	return &Bull{*newAnimal("\u001B[1;31mBull\u001B[0m", 3, time.Millisecond*600)}
}
