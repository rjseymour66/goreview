package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type item struct {
	Task        string    `json:"Task"`
	Done        bool      `json:"Done"`
	StartedAt   time.Time `json:"StartedAt"`
	CompletedAt time.Time `json:"CompletedAt"`
}

// List is a slice of todo items
type List []item

// Add a task
func (l *List) Add(task string) {
	if len(task) == 0 {
		fmt.Println("task cannot be empty.")
	}
	i := item{
		Task:        task,
		Done:        false,
		StartedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, i)
}

// Delete a task
func (l *List) Delete(i int) error {
	if i < 0 || i > len(*l) {
		return errors.New("task index is out of bounds")
	}
	ls := *l
	// adjust for 0 indexing
	// {1, [2], (3)} l.Delete(2)

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Complete a task
func (l *List) Complete(i int) error {
	if i < 0 || i > len(*l) {
		return errors.New("task index is out of bounds")
	}
	ls := *l
	ls[i-1].Done = true

	return nil
}

// Save the task to a file
func (l *List) Save(filename string) error {
	// marshal into json
	jsonList, err := json.Marshal(&l)
	if err != nil {
		return errors.New("error marshalling List")
	}
	// write to file
	return os.WriteFile(filename, jsonList, 0644)
}

// Get tasks from the file and populate a List
func (l *List) Get(filename string) error {
	// read from file
	contents, err := os.ReadFile(filename)
	if err != nil {
		// if the file does not exist, return without error
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errors.New("error reading file")
	}

	if len(contents) == 0 {
		return nil
	}

	// populate list
	return json.Unmarshal(contents, &l)
}

// Stringer to list the todos
func (l *List) String() string {
	var formatted strings.Builder

	for k, t := range *l {
		prefix := "[ ] "
		if t.Done {
			prefix = "[x] "
		}
		fmt.Fprintf(&formatted, "%s%d. %s\n", prefix, k+1, t.Task)
	}
	return formatted.String()
}
