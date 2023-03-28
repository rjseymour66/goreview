package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item represents a task.
type item struct {
	Task        string    `json:"Task"`
	StartedAt   time.Time `json:"StartedAt"`
	CompletedAt time.Time `json:"CompletedAt"`
	Completed   bool      `json:"Completed"`
}

// List is one or more tasks.
type List []item

// Add task adds a task to the list.
func (l *List) Add(task string) {
	item := item{
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
func (l *List) Get(filename string) error {

	// read the file as a slice of bytes
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(fileContents) == 0 {
		return nil
	}
	// unmarshal into a struct
	return json.Unmarshal(fileContents, l)
}

// Stringer (lists just the task)
func (l *List) String() string {

	formatted := ""

	// loop through tasks in list
	for k, t := range *l {

		prefix := "[ ] "
		if t.Completed {
			prefix = "[X] "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)

	}

	return formatted
}

// Verbose (includes CreatedAt)
