package job

//PosterJob работа почтальон
type PosterJob struct {
	Job
}

//NewPosterJob конструктор
func NewPosterJob() *PosterJob {
	return &PosterJob{NewJob("Sends mail")}
}
