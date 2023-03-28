package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	// arrange
	task := "Task 1"

	// act
	l := List{}
	l.Add(task)

	// assert
	if l[0].Task != task {
		t.Errorf("\ngot:  %s\nwant: %s", l[0].Task, task)
	}
}
func TestDelete(t *testing.T) {
	// arrange
	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	l := List{}
	for _, t := range tasks {
		l.Add(t)
	}

	// act
	l.Delete(2)

	// assert
	if l[1].Task != tasks[2] {
		t.Errorf("\ngot:  %s\nwant: %s", l[1].Task, tasks[2])
	}
}

func TestComplete(t *testing.T) {
	// arrange
	l := List{}
	l.Add("Task 1")

	// act
	l.Complete(1)

	// assert
	if !l[0].Done {
		t.Errorf("\ngot:  %t\nwant: %t", l[0].Done, true)
	}
}

// Save the task to a file
func TestSaveAndGet(t *testing.T) {
	// arrange
	l1 := List{}
	l2 := List{}

	l1.Add("Task 1 of 1")
	temp, err := os.CreateTemp("", "testFile")
	if err != nil {
		t.Errorf("error writing to file")
	}
	defer temp.Close()
	defer os.Remove(temp.Name())

	// act

	l1.Save(temp.Name())
	l2.Get(temp.Name())

	// assert
	if l1[0].Task != l2[0].Task {
		t.Errorf("\ngot:  %s\nwant: %s", l1[0].Task, l2[0].Task)
	}

}

func BenchmarkString(b *testing.B) {
	b.ReportAllocs()
	b.Logf("Loop %d times\n", b.N)

	l := List{}
	l.Add("Task 1")
	l.Add("Task 2")
	l.Complete(2)

	for i := 0; i < b.N; i++ {
		_ = l.String()
	}
}
