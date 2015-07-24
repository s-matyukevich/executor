package main

import (
	"sync"
	"time"
)

type Stage struct {
	isParallel bool
	tasks      []*Task
}

func NewStage(tasks []*Task, isParallel bool) *Stage {
	return &Stage{
		isParallel: isParallel,
		tasks:      tasks,
	}
}

func (s *Stage) ExecuteTasks(statusChanel chan *Task) {
	for _, task := range s.tasks {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer func() {
				var status int
				if r := recover(); r != nil {
					status = StatusFailed
				} else {
					status = StatusFinished
				}
				if task.Status != StatusExpired {
					task.Status = status
					statusChanel <- task
					wg.Done()
				}
			}()
			task.Status = StatusRunning
			statusChanel <- task
			task.Func()
		}()
		go func() {
			time.Sleep(TaskTimeout * time.Second)
			if task.Status == StatusRunning {
				task.Status = StatusExpired
				statusChanel <- task
				wg.Done()
			}
		}()
		if !s.isParallel {
			wg.Wait()
		}
	}
}
