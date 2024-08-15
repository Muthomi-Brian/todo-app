package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// simple function to create,read,update,delete a list of tast

type todo struct {
	name     string
	task     string
	dateTodo string
}

// create func(✅)
//
//	create logic for getting input(✅)
//
// create functionality for writing to a txt file(✅)
func create(n *todo, t *todo, d *todo) (string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the name of the task:")
	nName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", " ", ""
	}
	n.name = strings.TrimSpace(nName)

	fmt.Println("Enter the task description:")
	tTask, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", ""
	}
	t.task = strings.TrimSpace(tTask)

	fmt.Println("Enter the date you want to do it:")
	dDate, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", "", ""
	}
	d.dateTodo = strings.TrimSpace(dDate)

	if n.name != " " {
		fmt.Printf("The task %s was successfully created \n", n.name)
	}
	if t.task != " " {
		fmt.Print("The task description was successfully created \n")
	}

	if d.dateTodo != " " {
		fmt.Print("The task date was succesfully created \n")
	}

	return n.name, t.task, d.dateTodo
}

func write(a, b, c string, d string) {

	file, err := os.Create(d)
	// creates a new file

	if err != nil {
		fmt.Println("an error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(a + "\n")

	if err != nil {
		fmt.Println("an error creating file:", err)
		return
	}

	_, err = file.WriteString(b + "\n")

	if err != nil {
		fmt.Println("an error creating file:", err)
		return
	}
	_, err = file.WriteString(c + "\n")

	if err != nil {
		fmt.Println("an error creating file:", err)
		return
	}
	// the value d causes a side effect the file data.txt
}

// read file functionality
func read(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("an error occured trying to read file :", err)
		return
	}
	fmt.Println("file content:\n", string(data))
}

func update(a todo, b todo, c todo) (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("what is your new task:")
	nName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", " ", ""
	}
	a.name = strings.TrimSpace(nName)
	fmt.Println("what is your new task description:")
	btask, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", " ", ""
	}
	b.task = strings.TrimSpace(btask)

	fmt.Println("what is your new task date:")
	ctask, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return "", " ", ""
	}
	c.dateTodo = strings.TrimSpace(ctask)

	return a.name, b.task, c.dateTodo
}

func delete(a string) string {
	filename := a
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("error deleting file:", err)
		return ""
	}
	return a
}
func todoLogic() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose an operation: create, read, update, delete")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	var nameTodo, taskTodo, dateTodo todo
	fileName := "data.txt"

	if operation == "create" {
		name, task, date := create(&nameTodo, &taskTodo, &dateTodo)
		write(name, task, date, fileName)
	} else if operation == "read" {
		read(fileName)
	} else if operation == "update" {
		update(nameTodo, taskTodo, dateTodo)
	} else if operation == "delete" {
		delete(fileName)
	} else {
		fmt.Println("Invalid operation selected.")
	}
}

func main() {

	todoLogic()

}
