package main

import (
	"fmt"
)

type Task struct {
	Title string
	Done  bool
}

func showMenu() {
	fmt.Println("===== SIMPLE TASK MANAGEMENT =====")
	fmt.Println("1. SHOW TASKS")
	fmt.Println("2. ADD TASK")
	fmt.Println("3. MARK TASK AS DONE")
	fmt.Println("4. DELETE TASK")
	fmt.Println("5. EXIT")
	fmt.Print("Enter number: ")
}

func pressEnterToContinue() {
	var input string
	fmt.Print("Press ENTER to continue")
	fmt.Scanf("%s", &input)
	fmt.Println(" ")
}

/*
MENU FUNCTION
*/
func showTasks(tasks []Task) {
	if tasks == nil || len(tasks) == 0 {
		fmt.Println("No tasks found!")
	} else {
		for i, task := range tasks {
			fmt.Print(i + 1)
			if task.Done == true {
				fmt.Println(". [x] " + task.Title)
			} else {
				fmt.Println(". [ ] " + task.Title)
			}
		}
	}
}

func validateNewTask() string {
	var inputTask string

	for {
		fmt.Print("Add new task: ")
		fmt.Scanf("%s", &inputTask)

		if inputTask == "" {
			fmt.Println("Please input the correct task")
			continue
		} else {
			return inputTask
		}
	}
}

func validateChooseTask(length int, ask string) int {
	var inputTask int
	
	for {
		fmt.Print(ask)
		fmt.Scanf("%d", &inputTask)

		if inputTask <= length && inputTask >= 1 {
			return inputTask - 1
		} else {
			fmt.Println("Please choose the correct task")
			continue
		}
	}
}

func main() {
	var menu string = ""
	// inisiasi slice
	var tasks []Task

	for {
		showMenu()
		fmt.Scanf("%s", &menu)

		if menu == "1" {
			fmt.Println("===== SHOW TASKS =====")

			// Show Tasks
			showTasks(tasks)
			fmt.Println("")

			pressEnterToContinue()
			continue

		} else if menu == "2" {
			fmt.Println("===== ADD TASK =====")

			// Validate the input task
			var inputTask = validateNewTask()
			// Create new task
			var newTask = Task{inputTask, false}
			// Save the new task
			tasks = append(tasks, newTask)

			fmt.Println(" ")
			pressEnterToContinue()
			continue

		} else if menu == "3" {
			fmt.Println("===== MARK TASK =====")

			// Show Task
			showTasks(tasks)
			fmt.Println("")

			// Validate the input task
			inputTask := validateChooseTask(len(tasks), "Choose the task number to be marked: ")

			// Mark task
			tasks[inputTask].Done = true
			fmt.Println("Task marked successfully!")

			fmt.Println(" ")
			pressEnterToContinue()
		} else if menu == "4" {
			fmt.Println("===== DELETE TASK =====")

			// Show Task
			showTasks(tasks)

			// Validate the input task
			inputTask := validateChooseTask(len(tasks), "Choose the task number to be deleted: ")

			// Delete Task
			tasks = append(tasks[:inputTask], tasks[inputTask+1:]...)
			fmt.Println("Task has been deleted!")

			fmt.Println(" ")
			pressEnterToContinue()
		} else if menu == "5" {
			fmt.Println("===== EXIT =====")
			fmt.Println("See you next time!")
			break
		} else {
			fmt.Println("Incorrect menu input!")
			pressEnterToContinue()
		}
	}
}
