package Job

import (
	. "Front/src/Interfaces"
	"sync"
)

type Job struct {
	workers []IWorkable
	toDo    string
}

func NewJob(ToDo string) Job {
	return Job{toDo: ToDo}
}

func (l *Job) AddWorkers(worker ...IWorkable) {
	l.workers = append(l.workers, worker...)
}

func (l *Job) StartWork(HowMuchWork int) {
	wg := sync.WaitGroup{}
	wg.Add(len(l.workers))
	for _, w := range l.workers {
		go w.DoWork(&wg, l.toDo, HowMuchWork)
	}
	wg.Wait()
}

func DoItWork(wg *sync.WaitGroup, job IJob, workable ...IWorkable) {
	job.AddWorkers(workable...)
	job.StartWork(5)
	wg.Done()
}
