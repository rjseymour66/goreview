# Goreview

Where I will review Go project books to make sure I actually learned something.

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