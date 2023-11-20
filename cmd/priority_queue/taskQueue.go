package main

import (
	"fmt"
)

const (
	MaxPriority = 100
	MinPriority = 0
)

type Task struct {
	name     string
	priority int
}

func NewTask(name string, priority int) (*Task, error) {
	if priority > MaxPriority || priority < MinPriority {
		return nil, fmt.Errorf("priority must be between %d and %d", MinPriority, MaxPriority)
	}
	task := &Task{
		name:     name,
		priority: priority,
	}
	return task, nil
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetPriority() int {
	return t.priority
}

type TaskQueue []*Task

func (tq TaskQueue) Len() int {
	return len(tq)
}

func (tq TaskQueue) Less(i, j int) bool {
	return tq[i].priority > tq[j].priority
}

func (tq TaskQueue) Swap(i, j int) {
	tq[i], tq[j] = tq[j], tq[i]
}

func (tq *TaskQueue) Push(x any) {
	task := x.(*Task)
	*tq = append(*tq, task)
}

func (tq *TaskQueue) Pop() any {
	old := *tq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // avoid memory leak
	*tq = old[0 : n-1]
	return task
}
