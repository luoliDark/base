package snow

type singleton struct {
	worker *Worker
}
type Singleton struct {
	*singleton
}

var single *Singleton

func init() {
	single = &Singleton{&singleton{worker: getNewWork()}}
}

func (s *Singleton) GetWorker() *Worker {
	return s.worker
}

func GetInstance() *Singleton {
	return single
}

func getNewWork() *Worker {
	worker, _ := NewWorker(1)
	return worker
}
