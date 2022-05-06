package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

/*
Описание:
	Это поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов,
	ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
Преимущества:
	1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	2. Позволяет реализовать простую отмену и повтор операций.
	3. Позволяет реализовать отложенный запуск операций.
	4. Позволяет собирать сложные команды из простых.
	5. Реализует принцип открытости/закрытости.
Недостатки:
	1. Усложняет код программы из-за введения множества дополнительных классов.
*/

// Account некий аккаунт над которым будут производиться команды
type Account struct {
	name  string
	money float64
}

// NewAccount Конструктор
func NewAccount(name string, money float64) *Account {
	return &Account{name, money}
}

// aCommand родитель для всех команд, вспомогательный
type aCommand struct {
	account        *Account
	isCompleted    bool
	moneyToOperate float64
}

func (a *aCommand) IsCompleted() bool {
	return a.isCompleted
}

// newACommandAccount Конструктор
func newACommandAccount(account *Account, money float64) *aCommand {
	return &aCommand{account, false, money}
}

// Deposit Команда для пополнения баланса к аккаунту
type Deposit struct {
	aCommand
}

func NewDeposit(toDeposit float64, account *Account) *Deposit {
	return &Deposit{*newACommandAccount(account, toDeposit)}
}
func (d *Deposit) Execute() error {
	if d.isCompleted {
		return errors.New(fmt.Sprintf("deposit +$%f was completed", d.moneyToOperate))
	} else {
		d.isCompleted = true
		d.account.money += d.moneyToOperate
		fmt.Printf("%s: put money in account +$%f, ", d.account.name, d.moneyToOperate)
		fmt.Printf("new balance $%f\n", d.account.money)
		return nil
	}
}

// WithDraw Команда для снятия денег с аккаунта
type WithDraw struct {
	aCommand
}

func NewWithdraw(toWithDraw float64, account *Account) *WithDraw {
	return &WithDraw{*newACommandAccount(account, toWithDraw)}
}
func (w WithDraw) Execute() error {
	if w.isCompleted {
		return errors.New(fmt.Sprintf("withdraw -%f was completed", w.moneyToOperate))
	} else if w.account.money < w.moneyToOperate {
		return errors.New(w.account.name + ": haven't enough money not withdraw")
	} else {
		w.isCompleted = true
		w.account.money -= w.moneyToOperate
		fmt.Printf("%s: withdrawed from card -$%f, ", w.account.name, w.moneyToOperate)
		fmt.Printf("new balance $%f\n", w.account.money)
		return nil
	}
}

// Executable Интерфейс команд для взаимодействия с: CmdManager
type Executable interface {
	Execute() error
}

// CmdManager Менеджер по исполнения всех команд
type CmdManager struct {
	commands []Executable
}

func (e *CmdManager) Add(execute Executable) *CmdManager {
	e.commands = append(e.commands, execute)
	return e
}
func (e *CmdManager) Run() (ok error) {
	for _, command := range e.commands {
		if ok = command.Execute(); ok != nil {
			return ok
		}
	}
	return nil
}

func main() {
	fmt.Println(strings.Repeat("-", 65))
	rasul := NewAccount("Rasul", 0)
	ildar := NewAccount("Ildar", 0)
	cmdManager := CmdManager{}
	if err := cmdManager.
		Add(NewDeposit(820, rasul)).
		Add(NewWithdraw(139, rasul)).
		Add(NewDeposit(321, ildar)).
		Add(NewDeposit(132.3, rasul)).
		Add(NewWithdraw(192, ildar)).
		Run(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(strings.Repeat("-", 65))

	fmt.Println(*rasul)
	fmt.Println(*ildar)
}
