package main

import (
	"fmt"
	"os"
	"os/signal"
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
	go func(){
		for {
			select {
			case task := <- statusChanel:
				fmt.Printf("Task: %s entered status: %d\n", task.Name, task.Status)
			case <- timeoutChanel:
				fmt.Printf("Executor timeout")
				break;
			case <- interuptChanel:
				fmt.Printf("Exiting")
				break;
			}
		}
	}()
	for _, stage := range e.stages {
		stage.ExecuteTasks(statusChanel)
	}
}
