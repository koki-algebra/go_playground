package main

import (
	"container/heap"
	"fmt"
)

func main() {
	tasks := map[string]int{
		"AtCoder":       80,
		"Study AWS":     30,
		"Study English": 50,
	}

	// Create a priority queue, put the items in it, and establish the priority queue invariants.
	var taskQueue TaskQueue
	for name, priority := range tasks {
		task, err := NewTask(name, priority)
		if err != nil {
			panic(err)
		}
		taskQueue = append(taskQueue, task)
	}
	heap.Init(&taskQueue)

	// Insert a new task
	task, err := NewTask("Research", 70)
	if err != nil {
		panic(err)
	}
	heap.Push(&taskQueue, task)

	// Take the task out; they arrive in decreasing priority order
	for taskQueue.Len() > 0 {
		task := heap.Pop(&taskQueue).(*Task)
		fmt.Printf("name: %s, priority: %d\n", task.GetName(), task.GetPriority())
	}
}
