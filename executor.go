package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Executor struct {
	stages []*Stage
}

func (e *Executor) AddStage(stage *Stage) {
	e.stages = append(e.stages, stage)
}

func (e *Executor) Execute() {
	var statusChanel = make(chan *Task)
	var timeoutChanel = time.After(ExecutorTimeout * time.Minute)
	var interuptChanel = make(chan os.Signal)
	signal.Notify(interuptChanel, os.Interrupt, os.Kill)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case task := <-statusChanel:
				fmt.Printf("Task: %s entered status: %d\n", task.Name, task.Status)
			case <-timeoutChanel:
				fmt.Printf("Executor timeout")
				wg.Done()
				return
			case <-interuptChanel:
				fmt.Printf("Exiting")
				wg.Done()
				return
			}
		}
		wg.Done()
	}()
	for _, stage := range e.stages {
		stage.ExecuteTasks(statusChanel)
	}
	wg.Wait()
}
