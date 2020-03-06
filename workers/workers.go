package workers

import (
	"fmt"
	"sync"
)

type Pool struct {
	Workers     int
	TotalRecord int
	Records     [][]string
}

func NewPool(workers int, totalRecord int, records [][]string) Pool {
	return Pool{Workers: workers, TotalRecord: totalRecord, Records: records}
}

func (p Pool) Run() {
	tasks := make(chan []string, p.TotalRecord)
	errorlog := make(chan error, 1)
	var wg sync.WaitGroup
	for w := 1; w <= p.Workers; w++ {
		go p.worker(w, &wg, tasks, errorlog)
	}

	for t := 0; t < p.TotalRecord; t++ {
		wg.Add(1)
		tasks <- p.Records[t]
	}

	close(tasks)

	for t := 0; t < p.TotalRecord; t++ {
		err := <-errorlog
		if err != nil {
			fmt.Println(err)
		}

	}

	wg.Wait()

}

func (p Pool) worker(wid int, wg *sync.WaitGroup, tasks <-chan []string, errorlog chan<- error) {
	var err error
	for t := range tasks {
		if wid == 5 {
			err = fmt.Errorf("not insert data WorkerID: %d", wid)
		} else {
			fmt.Println("worker", wid, "processing job", t)
		}
		errorlog <- err
		wg.Done()
	}
}
