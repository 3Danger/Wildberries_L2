package main

import (
	"Builder/pkg/Address"
)

/*
	Паттерн 'строитель' необходим при конструировании сложного объекта,
	для которого может потребоваться много действий.
	Вместо того что бы создавать конструктор с огромным количеством параметров,
	которые могут быть еще и опциональными, в подобных случаях
	следует использовать шаблон проектирования 'строитель'
*/

func main() {
	address := Address.NewAddressBuilder().
		SetUserName("csamuro").
		SetCountry("Russia").
		SetRegion("Tatarstan").
		SetPhone("+7 999 532 2588").
		SetPost("224334").
		SetCity("Kazan").
		SetHome("44B").
		Build()
	address.PrintInfo()
}
