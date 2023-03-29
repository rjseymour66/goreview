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

var todoFile = ".todo.json"

// usage information
const usage = `
2-do
Copyright 2023
Usage:
  todo [option]
Options:`

func main() {

	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), usage[1:])
		flag.PrintDefaults()
	}

	// define flags
	add := flag.Bool("add", false, "Add task to list.")
	list := flag.Bool("list", false, "List todo tasks.")
	complete := flag.Int("complete", 0, "Mark a task as completed.")
	flag.Parse()

	// check if filename is defined in env var
	if os.Getenv("TODO_FILENAME") != "" {
		todoFile = os.Getenv("TODO_FILENAME")
	}

	// create list
	l := &todo.List{}
	if err := l.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, "Error: %w", err)
		os.Exit(1)
	}

	// use flags
	switch {
	case *add:
		// use getTask
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			os.Exit(1)
		}

		// tasd task
		l.Add(t)
		// save new list
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			os.Exit(1)
		}
	case *list:
		fmt.Print(l)
	case *complete > 0:
		// mark complete
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			os.Exit(1)
		}
		// save to file
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, "Error: %w", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

}

func getTask(r io.Reader, args ...string) (string, error) {
	// read from Args
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	// read from reader with scanner
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading task")
	}

	if len(scanner.Text()) == 0 {
		return "", fmt.Errorf("input cannot be empty")
	}

	return scanner.Text(), nil
}
