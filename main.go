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
		} else {
			return inputTask
		}
	}
}

func validateMarkInput(length int) int {
	for {
		fmt.Print("Choose the task number to be marked: ")
		var inputMark int
		fmt.Scanf("%d", &inputMark)

		if inputMark > length {
			fmt.Println("Please input the correct task")
		} else {
			return inputMark
		}
	}
}

func main() {
	var menu int = 0
	var tasks []Task

	for {
		showMenu()
		fmt.Scanf("%d", &menu)

		if menu == 1 {
			fmt.Println("===== SHOW TASKS =====")

			showTasks(tasks)
			fmt.Println("")

			pressEnterToContinue()
			continue

		} else if menu == 2 {
			fmt.Println("===== ADD TASK =====")

			var inputTask string = validateNewTask()
			var newTask = Task{inputTask, false}
			tasks = append(tasks, newTask)

			fmt.Println(" ")
			pressEnterToContinue()
			continue

		} else if menu == 3 {
			fmt.Println("===== MARK TASK =====")

			showTasks(tasks)
			fmt.Println("")

			// validate and mark task
			inputMark := validateMarkInput(len(tasks))

			tasks[inputMark-1].Done = true

			fmt.Println("Task has been marked successfully!")

			fmt.Println(" ")
			pressEnterToContinue()
		} else if menu == 4 {
			fmt.Println("===== DELETE TASK =====")

		} else if menu == 5 {
			fmt.Println("===== EXIT =====")
			fmt.Println("See you next time!")
			break
		} else {
			fmt.Println("Incorrect menu input!")
			pressEnterToContinue()
		}
	}
}
