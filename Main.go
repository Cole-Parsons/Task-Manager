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
	var testTask task
	for {
		i := printChoices()
		if i == 1 {
			fmt.Println("Creating task\n")
			testTask = task{
				name:    "laundry",
				dueDate: "tomorrow",
				ID:      1,
			}
			tasks = append(tasks, testTask)
		} else if i == 2 {
			fmt.Println("Listing tasks\n")
			fmt.Println(testTask)
		} else if i == 3 {
			fmt.Println("Marking tast as complete\n")
		} else if i == 4 {
			fmt.Println("Deleting Tasks\n")
		} else if i == 5 {
			fmt.Println("Saving Tasks\n")
		} else if i == 6 {
			os.Exit(0)
		} else {
			fmt.Println("Invalid input try again\n")
		}
	}
}

func printChoices() int {
	var num int
	fmt.Println("What do you want to do?\n")
	fmt.Println("1. Create Task \n2. List Tasks\n3. Mark Task as Complete\n4. Delete Task\n5. Save Tasks (Warning: Tasks will disappear without saving on close)\n6. Exit")
	fmt.Print("\nEnter choice: ")
	fmt.Scan(&num)

	return num
}
