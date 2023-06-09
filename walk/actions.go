package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Filters out the given path if it is a directory, is less than the minSize,
// or the file extension does not match the ext
func filterOut(path string, ext string, minSize int64, info os.FileInfo) bool {
	// check if it is a directory
	if info.IsDir() || info.Size() < minSize {
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
	_, err := fmt.Fprintln(out, path)
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

// archiveFile preserves the relative direcory tree so files are correctly archived relative
// to the root, and compresses data.
// - destDir: The destination directory for the archived files. The target.
// - root: Directory where the search was started. root is used to determine the relative path of the files
// that you are archiving.
// - path: Path of the file that is being archived.
func archiveFile(destDir, root, path string) error {
	// confirm that the provided destination is a directory
	info, err := os.Stat(destDir)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", info)
	}

	// get the relative directory from root of the file to be archived
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		return err
	}

	// create new name for .gz file
	dest := fmt.Sprintf("%s.gz", filepath.Base(path))

	// Create the path for the final zipped file
	targetPath := filepath.Join(destDir, relDir, dest)

	// create the target directory tree
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}
	/*
		GZIP Implementation
	*/

	// Open the targetPath so you can write to it
	out, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer out.Close()

	// open the source file to read from it
	in, err := os.Open(path)
	if err != nil {
		return nil
	}

	defer in.Close()

	// create a new gzip writer and name
	zw := gzip.NewWriter(out)
	zw.Name = filepath.Base(path)

	// write compressed data with io.Copy
	if _, err := io.Copy(zw, in); err != nil {
		return err
	}

	// close the gzip writer
	if err := zw.Close(); err != nil {
		return err
	}

	// close the destination (zipped file)
	return out.Close()
}
