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
// create functionality for writing to a txt file
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

	return n.name, t.task, d.dateTodo
}

func write(a, b, c string) {
	d := "data.txt"

	file, err := os.Create(d)

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

}

func main() {
	var nameTodo, taskTodo, dateTodo todo
	name, task, date := create(&nameTodo, &taskTodo, &dateTodo)
	// check if input is passed
	if name != "" && task != "" && date != "" {
		fmt.Printf("Task created: %s\nDescription: %s\nDate: %s\n", name, task, date)
	} else {
		fmt.Println("Failed to create task.")
	}
	write(name, task, date)
}
