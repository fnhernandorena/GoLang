package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	name        string
	description string
	complete    bool
}

func main() {
	tasks := make(map[int]task)
	lastID := 0

	for {
		fmt.Println("\n Pick the wanted option:")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Complete task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Option: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			if len(tasks) == 0 {
				fmt.Println("There are no tasks")
			} else {
				for id, task := range tasks {
					fmt.Printf(" %d. %s: %s (completed: %t)\n", id, task.name, task.description, task.complete)
				}
			}

		case 2:
			fmt.Print("Task name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Task description: ")
			scanner.Scan()
			description := scanner.Text()

			lastID++

			tasks[lastID] = task{
				name:        name,
				description: description,
				complete:    false,
			}
			fmt.Printf("Task %d added\n", lastID)

		case 3:
			fmt.Print("Write completed Task ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if task, ok := tasks[id]; ok {
				task.complete = true
				tasks[id] = task
				fmt.Printf("Task %d completed\n", id)
			} else {
				fmt.Println("No valid ID task")
			}
		case 4:
			fmt.Print("Write completed Task ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())

			if _, ok := tasks[id]; ok {
				delete(tasks, id)
				fmt.Printf("Task with id: %d, deleted.\n", id)
			} else {
				fmt.Println("No valid ID task")
			}

		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		}

	}
}
