package main

import (
	_ "expvar"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	name     string
	duration time.Duration
}

type worker struct {
	id int
}

func (w worker) process(j job) {
	//fmt.Printf("worker%d: started %s, working for %fs\n", w.id, j.name, j.duration.Seconds())
	//time.Sleep(j.duration)
    aVal := 0
    for cnt :=0; cnt < 1000; cnt++{
        //Just doing nothing
        //fmt.Printf("%v", cnt)
        aVal = cnt * aVal
    }
	//fmt.Printf("worker%d: completed %s!\n", w.id, j.name)
}

func main() {
    startTime := time.Now()
	wg := &sync.WaitGroup{}
	jobCh := make(chan job)

	// start workers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		w := worker{i}
		go func(w worker) {
			for j := range jobCh {
				w.process(j)
			}
			wg.Done()
		}(w)
	}

	// add jobs to queue
	for i := 0; i < 1000; i++ {
		name := fmt.Sprintf("job-%d", i)
		duration := time.Duration(rand.Intn(1000)) * time.Millisecond
		//fmt.Printf("adding: %s %s\n", name, duration)
		jobCh <- job{name, duration}
	}

	// close jobCh and wait for workers to complete
	close(jobCh)
	wg.Wait()
    endTime := time.Now()
    fmt.Println("Time Taken: ", endTime.Sub(startTime))
    
    /* startTime = time.Now()
    // dd := 100
    for cnt :=0; cnt < 1000; cnt++{
        //Just doing nothing
        fmt.Printf("%v", cnt)
    }
    endTime = time.Now()
    fmt.Println("Time Taken - loop: ", endTime.Sub(startTime))*/
}