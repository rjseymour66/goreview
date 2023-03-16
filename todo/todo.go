package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Item represents a task.
type Item struct {
	Task        string    `json:"task"`
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
	Completed   bool      `json:"completed"`
}

// List is one or more tasks.
type List []Item

// Add task adds a task to the list.
func (l *List) Add(task string) {
	item := Item{
		Task:        task,
		StartedAt:   time.Now(),
		CompletedAt: time.Time{},
		Completed:   false,
	}
	*l = append(*l, item)
}

// Delete task removes a task from the list.
func (l *List) Delete(i int) error {
	if i < 0 || i > len(*l) {
		return fmt.Errorf("Task index %d is out of bounds", i)
	}
	ls := *l
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Complete task marks a task completed.
func (l *List) Complete(i int) error {
	if i < 0 || i > len(*l) {
		return fmt.Errorf("Task index %d is out of bounds", i)
	}

	ls := *l
	ls[i-1].CompletedAt = time.Now()
	ls[i-1].Completed = true

	return nil
}

// Save to file as JSON
func (l *List) Save(filename string) error {
	// encode as JSON
	jsonList, err := json.Marshal(&l)
	if err != nil {
		return err
	}
	// write to file
	return os.WriteFile(filename, jsonList, 0644)
}

// Get from file and populate list

// Stringer (lists just the task)

// Verbose (includes CreatedAt)
