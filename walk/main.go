package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	ext  string
	size int64
	list bool
	del  bool
	wLog io.Writer
}

const usage = `
gwalker

Usage:
  gwalker [options]
Options:`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), usage[1:])
		flag.PrintDefaults()
	}

	// root for the directory to start walking
	root := flag.String("root", "", "Top-level directory to search.")
	// log for the name of the logfile
	logFile := flag.String("log", "", "Name of log file. Must be of type io.Writer.")
	// list to list the files in a directory only
	list := flag.Bool("list", false, "List files in root directory tree.")
	// del to delete files from a directory
	del := flag.Bool("del", false, "Delete a file in root directory tree.")
	// ext to filter by extension
	ext := flag.String("ext", "", "Extension to filter files.")
	// size to filter by min file size
	size := flag.Int64("size", 0, "Size to filter files.")

	// TODO archive to create an archive file from a directory
	// archive := flag.Bool("archive", false, "Archive filtered files")
	flag.Parse()

	var (
		// log everything to STDOUT by default
		f   = os.Stdout
		err error
	)

	// Check logFile flag and set
	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		defer f.Close()
	}

	// create a config for filter criteria provided by flags
	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		wLog: f,
	}

	// execute the run method
	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	// create a logger and assign to cfg
	delLogger := log.New(cfg.wLog, "DELETED:", log.LstdFlags)

	// Walk only returns functions that return error.
	return filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			// check if func() returns an error while accessing a path. This prevents panics
			if err != nil {
				return err
			}
			// filter files
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}
			// list files
			if cfg.list {
				return listFile(path, out)
			}
			// archive files
			// delete files
			if cfg.del {
				return delFile(path, delLogger)
			}
			// list files again as default option
			return listFile(path, out)
		})
}

/*
In filepath.Walk(root, func(path string, info fs.FileInfo, err error), the function works like this:

The path argument uses the first arg to Walk as a prefix while walking the filesystem. For example:
- root/<path>

The info argument is the file info for the current path value. Essentially, this function calls the following:

info, err := lstat(path)

and makes info available for operations on each path that Walk executes on.

The err argument determines whether the Walk function continues traversing the directory tree. If it returns a
non-nil error, it returns. Otherwise, it uses the function logic to determine if it skips one or all of the
files and directories.
*/
