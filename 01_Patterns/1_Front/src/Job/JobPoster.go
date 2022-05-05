package Job

type PosterJob struct {
	Job
}

func NewPosterJob() *PosterJob {
	return &PosterJob{NewJob("Sends mail")}
}
