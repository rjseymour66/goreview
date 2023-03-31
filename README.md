# Goreview

Where I will review Go project books to make sure I actually learned something.

## Walk

[Link to original project](https://github.com/rjseymour66/command-line/tree/master/walk)

### File logic

Logic for file system crawler. Desends into a directory tree to look for files that match a criteria.

- [ ] Filter path out if:
  - [ ] is a directory
  - [ ] greater than minimum size
  - [ ] wrong extension
- [ ] Lists files at the specified path
- [ ] Deletes files at the specified path and logs the path
- [ ] Creates a zip file out of a file in a destination directory:
  - [ ] confirm that the provided destination is a directory
  - [ ] get the relative path between the pwd and destination dir
  - [ ] create a .gz file name for the zip file
  - [ ] construct the targetPath directory structure from the destination path, relative path, and destination file name
  - [ ] Make directories for the targetPath directory structure
  - [ ] Zip the file
  - [ ] Open the targetPath file
     - [ ] Open the file w the contents to zip
     - [ ] Create a zip writer
     - [ ] Copy the contents from the zip writer to the file
     - [ ] Close the zip writer
     - [ ] return an error on fail

### File actions tests
- [ ] Test filtering with extension and size

### main logic

- [ ] Model the filter criteria
- [ ] Create flags for the following:
  - [ ] root for the directory to start walking
  - [ ] log for the name of the logfile
  - [ ] list to list the files in a directory only
  - [ ] archive to create an archive file from a directory
  - [ ] del to delete files from a directory
  - [ ] ext to filter by extension
  - [ ] size to filter by min file size
- [ ] Open the logfile
- [ ] create a config for filter criteria provided by flags
- [ ] execute the run method
- [ ] Write the run method with currying
  - [ ] use filepath.Walk to check the flags

### main test

- [ ] TestRun
- [ ] TestRunDeleteExtension
- [ ] createTempDir helper for delete tests
- [ ] TestRunArchive

## Todo app

[Link to original project](https://github.com/rjseymour66/command-line/tree/master/todo)

### Todo logic

Logic for the todo application. Works with a file that stores todos in JSON format.
- [ ] Model a task item 
- [ ] Model a list of tasks
- [ ] Add a task to the list 
- [ ] Delete a task from the list 
- [ ] Mark a task as completed
- [ ] Save the list to a file 
- [ ] Read the todo list from a file and populate a list
- [ ] Implement Stringer so you can view the output 


### Executable and flags

`cmd/todo/todo.go` (main.go)

Manages the executable, which is primarily parsing flags. Accepts input from STDIN or flags.
- [ ] Create custom usage info 
- [ ] create command line flags:
  - [ ] add 
  - [ ] list
  - [ ] complete 
- [ ] check for env var for todo filename 
- [ ] Create a list from the file 
- [ ] evaluate flags 