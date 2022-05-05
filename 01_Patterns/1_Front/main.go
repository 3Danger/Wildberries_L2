package main

import (
	"Front/src/Animal"
	"Front/src/Job"
	"sync"
)

/*
	Фасад позволяет скрыть сложные части системы,
	давая доступ только к необходимым рычагам,
	тем самым упрощая использование системы для клиента,

	В нашем примере в качестве фасада выступает функция DoItWork() из пакета Job
*/

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	// DoItWork() Это наш фасад найма на работу, и начала работы
	go Job.DoItWork(&wg, Job.NewDancerJob(), Animal.NewElephant())
	go Job.DoItWork(&wg, Job.NewLoaderJob(), Animal.NewBull())
	go Job.DoItWork(&wg, Job.NewPosterJob(), Animal.NewMonkey())
	wg.Wait()
}
