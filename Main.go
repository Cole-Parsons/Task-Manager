package main

import (
	"fmt"
	"os"
)

type task struct {
	name    string
	dueDate string
	ID      int
}

func main() {
	var tasks []task
	var newTask task
	var counter int = 0
	for {
		i := printChoices()
		if i == 1 {
			counter++
			newTask = createTask(counter)
			tasks = append(tasks, newTask)
		} else if i == 2 {
			for j := 0; j < len(tasks); j++ {
				fmt.Println(tasks[j])
			}
		} else if i == 3 {
			var num int
			fmt.Println("Which task would you like to remove?")
			for j := 0; j < len(tasks); j++ {
				fmt.Println(tasks[j])
			}
			fmt.Print("ID Number: ")
			fmt.Scan(&num)
			num -= 1
			tasks = append(tasks[:num], tasks[num+1:]...)
		} else if i == 4 {
			fmt.Println("Saving Tasks\n")
		} else if i == 5 {
			os.Exit(0)
		} else {
			fmt.Println("Invalid input try again\n")
		}
	}
}

func printChoices() int {
	var num int
	fmt.Println("What do you want to do?\n")
	fmt.Println("1. Create Task \n2. List Tasks\n3. Remove Task\n4. Save Tasks (Warning: Tasks will disappear without saving on close)\n5. Exit")
	fmt.Print("\nEnter choice: ")
	fmt.Scan(&num)

	return num
}

func createTask(num int) task {
	var name string
	var dueDate string
	fmt.Print("What is the name of the Task?\nEnter Name: ")
	fmt.Scan(&name)
	fmt.Print("\nWhen do you want the task to be done?\nEnter Date: ")
	fmt.Scan(&dueDate)
	newTask := task{
		name:    name,
		dueDate: dueDate,
		ID:      num,
	}
	return newTask
}
