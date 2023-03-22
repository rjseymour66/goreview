package todo

import (
	"os"
	"testing"
)

// Add task
func TestAdd(t *testing.T) {
	// arrange
	tc := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
	// act
	l := List{}
	for _, t := range tc {
		l.Add(t)
	}
	// assert
	if tc[0] != l[0].Task {
		t.Errorf("\ngot:  %s\nwant: %s", tc[0], l[0].Task)
	}

	if tc[1] != l[1].Task {
		t.Errorf("\ngot:  %s\nwant: %s", tc[1], l[1].Task)
	}

	if tc[2] != l[2].Task {
		t.Errorf("\ngot:  %s\nwant: %s", tc[2], l[2].Task)
	}

}

// Delete task
func TestDelete(t *testing.T) {
	// arrange
	tc := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	// act
	l := List{}
	for _, t := range tc {
		l.Add(t)
	}
	l.Delete(2)

	// assert
	if l[1].Task != tc[2] {
		t.Errorf("\ngot:  %s\nwant: %s", l[1].Task, tc[2])
	}
}

// Complete task
func TestComplete(t *testing.T) {
	// arrange
	tc := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}
	// act
	l := List{}
	for _, t := range tc {
		l.Add(t)
	}

	l.Complete(1)

	// assert
	if l[0].Completed != true {
		t.Errorf("\ngot:  %t\nwant: %t", l[0].Completed, true)
	}
}

// Get from file and populate list
func TestSaveGet(t *testing.T) {
	// create two lists
	l1 := List{}
	l2 := List{}

	// Add a task to l1
	l1.Add("Task 1")

	// create a temp file, defer its removal
	temp, err := os.CreateTemp("", "taskfile")
	if err != nil {
		t.Fatalf("Error creating temp file")
	}
	defer temp.Close()
	defer os.Remove(temp.Name())
	// save l1 to the temp file
	l1.Save(temp.Name())
	// get l1 from the temp file, save it in l2
	if err := l2.Get(temp.Name()); err != nil {
		t.Errorf("Error in Get method")
	}
	// compare l1 and l2 tasks
	if l1[0].Task != l2[0].Task {
		t.Errorf("\ngot:  %s\nwant: %s", l1[0].Task, l2[0].Task)
	}
}
