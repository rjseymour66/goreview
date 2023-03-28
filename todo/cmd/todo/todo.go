package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

const usage = `
todo tool. Additional info
Copyright 2023
Usage:
  todo [options]
Options:`

// todofile name
var todoFile = ".todofile.json"

// create command line flags:
// add
func main() {
	flag.Usage = func() {
		// fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed to learn the Go programming language\n", os.Args[0])
		// fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2023\n")
		// fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		fmt.Fprintln(flag.CommandLine.Output(), usage[1:])
		flag.PrintDefaults()
	}

	add := flag.Bool("add", false, "Task to add")
	list := flag.Bool("list", false, "List the todos")
	complete := flag.Int("complete", 0, "Complete the todo")
	flag.Parse()

	// convert complete into an int

	// check for env var for todo filename
	if os.Getenv("TODO_FILE") != "" {
		todoFile = os.Getenv("TODO_FILE")
	}
	// Create a list from the file
	l := &todo.List{}
	if err := l.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// evaluate flags
	switch {
	case *add:
		// use get task to get the task
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// add the task to the list
		l.Add(task)
		// save the new list
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *list:
		// print the list with stringer
		fmt.Print(l)
	case *complete > 0:
		// mark the item completed
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// save the new list
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}

// getTask gets the task from STDIN or cli args
func getTask(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Input cannot be empty")
	}

	return s.Text(), nil
}
