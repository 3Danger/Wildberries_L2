package Interfaces

import "sync"

type IWorkable interface {
	DoWork(wg *sync.WaitGroup, todo string, count int)
}
