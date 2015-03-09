package main

import (
	"sync"
	"testing"
)

type taskOrderItem struct {
	Key    string
	Status int
}

func TestSequentialExecutiong(t *testing.T) {
	order := []*taskOrderItem{
		&taskOrderItem{Key: "task1-1", Status: StatusRunning},
		&taskOrderItem{Key: "execution", Status: StatusRunning},
		&taskOrderItem{Key: "task1-1", Status: StatusFinished},
		&taskOrderItem{Key: "task1-2", Status: StatusRunning},
		&taskOrderItem{Key: "execution", Status: StatusRunning},
		&taskOrderItem{Key: "task1-2", Status: StatusFailed},
	}
	var statusChanel = make(chan *Task)
	executionTask := &Task{
		Name:   "execution",
		Status: StatusRunning,
	}
	task1 := &Task{
		Name: "task1-1",
		Func: func() {
			statusChanel <- executionTask
		},
	}
	task2 := &Task{
		Name: "task1-2",
		Func: func() {
			statusChanel <- executionTask
			panic("panic")
		},
	}
	stage := NewStage([]*Task{task1, task2}, false)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for _, item := range order {
			task := <-statusChanel
			if task.Status != item.Status || task.Name != item.Key {
				t.Errorf("Wrong order, expected: {Name: %s, Status: %d}, got: {Name: %s, Status: %d}", item.Key, item.Status, task.Name, task.Status)
			}
		}
		wg.Done()
	}()
	t.Log("ExecuteTask started")
	stage.ExecuteTasks(statusChanel)
	wg.Wait()
}
