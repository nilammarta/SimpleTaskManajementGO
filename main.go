package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Title string
	Done  bool
}

/*
*
Helper Function
*/
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
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Add new task: ")
		inputTask, _ := reader.ReadString('\n')

		// remove newline di akhir
		inputTask = strings.TrimSpace(inputTask)

		if inputTask == "" {
			fmt.Println("Please input the correct task")
			continue
		} else {
			return inputTask
		}
	}
}

func validateChooseTask(length int, ask string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(ask)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// convert number string input menjadi int
		numberTask, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please input a number")
			continue
		}

		if numberTask <= length && numberTask >= 1 {
			return numberTask - 1
		} else {
			fmt.Println("Please choose the correct task")
			continue
		}
	}
}

/*
MAIN FUNCTION
*/
func addTask(tasks *[]Task) {
	fmt.Println("===== ADD TASK =====")

	// Validate the input task
	var inputTask = validateNewTask()
	// Create new task
	var newTask = Task{inputTask, false}
	// Save the new task
	*tasks = append(*tasks, newTask)

	fmt.Println(" ")
	pressEnterToContinue()
}

func markTaskAsDone(tasks *[]Task) {
	fmt.Println("===== MARK TASK =====")

	// Show Task
	showTasks(*tasks)
	fmt.Println("")

	// Validate the input task
	inputTask := validateChooseTask(len(*tasks), "Choose the task number to be marked: ")

	// Mark task
	(*tasks)[inputTask].Done = true
	fmt.Println("Task marked successfully!")

	fmt.Println(" ")
	pressEnterToContinue()
}

func deleteTask(tasks *[]Task) {
	fmt.Println("===== DELETE TASK =====")

	// Show Task
	showTasks(*tasks)

	// Validate the input task
	inputTask := validateChooseTask(len(*tasks), "Choose the task number to be deleted: ")

	// Delete Task
	*tasks = append((*tasks)[:inputTask], (*tasks)[inputTask+1:]...)
	fmt.Println("Task has been deleted!")

	fmt.Println(" ")
	pressEnterToContinue()
}

/*
Main App
*/
func main() {
	var menu string
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
			addTask(&tasks)
			continue
		} else if menu == "3" {
			markTaskAsDone(&tasks)
			continue
		} else if menu == "4" {
			deleteTask(&tasks)
			continue
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
