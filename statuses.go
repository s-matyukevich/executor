package main

const (
	StatusWaiting = iota
	StatusRunning 
	StatusFailed 
	StatusExpired 
	StatusFinished 
)

const TaskTimeout = 1
const ExecutorTimeout = 10
