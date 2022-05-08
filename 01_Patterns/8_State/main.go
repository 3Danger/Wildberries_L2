package main

import (
	"errors"
	"fmt"
)

/*
Краткое описание:
	Состояние — это поведенческий паттерн проектирования,
	который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта.
Плюсы:
	1. Избавляет от множества больших условных операторов машины состояний.
	2. Концентрирует в одном месте код, связанный с определённым состоянием.
	3. Упрощает код контекста.
Минусы:
	1. Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

// State определим интерфейс для состояний
type State interface {
	On() error
	Off() error
}

// On Off возможные состоянии лампочки
type On struct{}
type Off struct{}

// On Off Различное поведение в зависимости от состояния
func (On) On() error {
	return errors.New(Red("err light bulb is already on"))
}
func (On) Off() error {
	fmt.Println(Grn("light bulb is switched OFF"))
	return nil
}
func (Off) On() error {
	fmt.Println(Grn("light bulb is switched ON!"))
	return nil
}
func (Off) Off() error {
	return errors.New(Red("err light bulb is already off"))
}

// LightBulb наша лампочка где будут применяться состояния
type LightBulb struct {
	state State
}

func NewLightBulb() *LightBulb {
	return &LightBulb{Off{}}
}

// On Off Переключение состояний
func (l *LightBulb) On() {
	if ok := l.state.On(); ok != nil {
		fmt.Println(ok)
	} else {
		l.state = On{}
	}
}
func (l *LightBulb) Off() {
	if ok := l.state.Off(); ok != nil {
		fmt.Println(ok)
	} else {
		l.state = Off{}
	}
}

func main() {
	bulb := NewLightBulb() // Создаем объекта с состоянием выкл
	bulb.Off()             // Поскольку она уже выключена, будет сообщение об ошибке
	bulb.On()              // Изначально была выкл, поэтому без проблем переключится... итд
	bulb.On()
	bulb.Off()
	bulb.On()
	bulb.Off()
	bulb.Off()
	bulb.Off()
}

// Red Grn Цвета для удобного чтения
func Red(s string) string { return "\033[0;31m" + s + "\033[0;0m" }
func Grn(s string) string { return "\033[0;32m" + s + "\033[0;0m" }
