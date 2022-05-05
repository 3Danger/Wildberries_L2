package Job

type DancerJob struct {
	Job
}

func NewDancerJob() *DancerJob {
	return &DancerJob{NewJob("Dancing")}
}
