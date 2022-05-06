package Job

type LoaderJob struct {
	Job
}

func NewLoaderJob() *LoaderJob {
	return &LoaderJob{NewJob("mowing Garbage")}
}
