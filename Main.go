package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Name    string
	DueDate string
	ID      int
}

func main() {
	var Tasks []Task
	var NewTask Task
	var counter int = 0
	for {
		i := printChoices()
		if i == 1 {
			counter++
			NewTask = createTask(counter)
			Tasks = append(Tasks, NewTask)
		} else if i == 2 {
			for j := 0; j < len(Tasks); j++ {
				fmt.Println(Tasks[j])
			}
		} else if i == 3 {
			var num int
			fmt.Println("Which task would you like to remove?")
			for j := 0; j < len(Tasks); j++ {
				fmt.Println(Tasks[j])
			}
			fmt.Print("ID Number: ")
			fmt.Scan(&num)
			num -= 1
			Tasks = append(Tasks[:num], Tasks[num+1:]...)
		} else if i == 4 {
			fmt.Println("Saving Tasks\n")
			saveJSON(Tasks)
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

func createTask(num int) Task {
	var name string
	var dueDate string
	fmt.Print("What is the name of the Task?\nEnter Name: ")
	fmt.Scan(&name)
	fmt.Print("\nWhen do you want the task to be done?\nEnter Date: ")
	fmt.Scan(&dueDate)
	NewTask := Task{
		Name:    name,
		DueDate: dueDate,
		ID:      num,
	}
	return NewTask
}

func saveJSON(Tasks []Task) {
	data, err := json.MarshalIndent(Tasks, "", "  ")
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	}
	fmt.Println("Tasks saved successfully")
}
