package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Краткое описание:
	Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.
Плюсы:
	1. Избавляет класс от привязки к конкретным классам продуктов.
	2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	3. Упрощает добавление новых продуктов в программу.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Может привести к созданию больших параллельных иерархий классов,
		так как для каждого класса продукта надо создать свой подкласс создателя.
*/

type IComputer interface {
	GetType() string
	PrintInfo()
}

type Computer struct{ core, ram, vram int }

type ServerComputer struct {
	Computer
}
type LaptopComputer struct {
	Computer
	monitor bool
}
type Playstation struct {
	Computer
	joystick bool
}

func NewServerComputer() *ServerComputer {
	return &ServerComputer{Computer{48, 256, 1}}
}
func NewLaptopComputer() *LaptopComputer {
	return &LaptopComputer{Computer{8, 16, 6}, true}
}
func NewPlaystation(joystick bool) *Playstation {
	return &Playstation{Computer{8, 16, 8}, joystick}
}

func (ServerComputer) GetType() string { return "ServerComputer" }
func (LaptopComputer) GetType() string { return "LaptopComputer" }
func (Playstation) GetType() string    { return "Playstation___" }

func (s ServerComputer) PrintInfo() {
	fmt.Printf("%s, Cores %d, RAM %dgb, video RAM %dgb\n", s.GetType(), s.core, s.ram, s.vram)
}
func (l LaptopComputer) PrintInfo() {
	fmt.Printf("%s, Cores %d, RAM %dgb, video RAM %dgb, monitor[%v]\n", l.GetType(), l.core, l.ram, l.vram, l.monitor)
}
func (p Playstation) PrintInfo() {
	fmt.Printf("%s, Cores %d, RAM %dgb, video RAM %dgb, joystick[%v]\n", p.GetType(), p.core, p.ram, p.vram, p.joystick)
}

//FabricMethod наш фабричный метод
func FabricMethod() IComputer {
	rand.Seed(time.Now().UnixNano())
	switch rand.Int() % 3 {
	case 0:
		return NewLaptopComputer()
	case 1:
		return NewPlaystation(rand.Int()&1 == 1)
	case 2:
		return NewServerComputer()
	}
	return nil
}

func main() {
	const N = 10
	computers := make([]IComputer, N)
	for i := 0; i < N; i++ {
		computers[i] = FabricMethod()
	}
	for _, v := range computers {
		v.PrintInfo()
	}
}
