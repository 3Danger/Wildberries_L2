package job

//LoaderJob грузчик
type LoaderJob struct {
	Job
}

//NewLoaderJob конструктор
func NewLoaderJob() *LoaderJob {
	return &LoaderJob{NewJob("mowing Garbage")}
}
