package Interfaces

type IJob interface {
	AddWorkers(...IWorkable)
	StartWork(int)
}
