package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Filters out the given path if it is a directory, is greater than the minSize,
// or the file extension does not match the ext
func filterOut(path string, ext string, minSize int64, info os.FileInfo) bool {
	// check if it is a directory
	if !info.IsDir() || info.Size() > minSize {
		return true
	}
	// check file extension
	if ext != "" && filepath.Ext(path) != ext {
		return true
	}
	return false
}

// listFile prints the file to out.
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintf(out, path)
	return err
}

// delFile deletes the file and logs the file path.
func delFile(path string, delLogger *log.Logger) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

func archiveFile(destDir, root, path string) error {
	// Creates a zip file out of a file in a destination directory:
	// confirm that the provided destination is a directory
	// get the relative path between the pwd and destination dir
	// create a .gz file name for the zip file
	// construct the targetPath directory structure from the destination path, relative path, and destination file name
	// Make directories for the targetPath directory structure
	// Zip the file
	//// Open the targetPath file
	//// Open the file w the contents to zip
	//// Create a zip writer
	//// Copy the contents from the zip writer to the file
	//// Close the zip writer
	//// return an error on fail
	return nil
}
