package job

//DancerJob работа где танцуют
type DancerJob struct {
	Job
}

//NewDancerJob конструктор
func NewDancerJob() *DancerJob {
	return &DancerJob{NewJob("Dancing")}
}
