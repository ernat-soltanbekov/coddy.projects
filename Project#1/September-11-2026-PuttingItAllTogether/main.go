package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func addTask(slice []Task, stroka string) []Task {
	newSlice := append(slice, Task{stroka, false})
	return newSlice
}

func viewAllTasks(parametr []Task) {
	for i := 0; i < len(parametr); i++ {
		if parametr[i].Completed == true {
			fmt.Printf("[x] %s\n", parametr[i].Name)
		} else {
			fmt.Printf("[ ] %s\n", parametr[i].Name)
		}
	}
}

func completeTask(tasklist *[]Task, index int) {
	(*tasklist)[index].Completed = true
}

func removeTask(sliceTasks []Task, index int) []Task {
	return append(sliceTasks[:index], sliceTasks[index+1:]...)
}

func main() {
	var list []Task
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanlist := strings.Split(scanner.Text(), ",")
	appName := scanlist[0]
	userName := scanlist[1]
	scanner.Scan()
	scanion := scanner.Text()
	if scanion != "" {
		taskSlice := strings.Split(scanion, ",")
		for count := 0; count < len(taskSlice); count++ {
			taskData := strings.Split(taskSlice[count], ":")
			booledTask, err := strconv.ParseBool(taskData[1])
			if err != nil {
				fmt.Printf("Invalid value %v", err)
				return
			}
			list = append(list, Task{taskData[0], booledTask})
		}
	}
	fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	fmt.Printf("1. Add Task\n")
	fmt.Printf("2. View Tasks\n")
	fmt.Printf("3. Complete Task\n")
	fmt.Printf("4. Remove Task\n")
	fmt.Printf("5. Exit\n")
	fmt.Printf("Current tasks: %d\n", len(list))
	scanner.Scan()
	actions := strings.Split(scanner.Text(), ",")
	for i := 0; i < len(actions); i++ {
		everySolution := strings.Split(actions[i], "|")
		switch everySolution[0] {
		case "add":
			fmt.Printf("--- ADD TASK ---\n")
			list = addTask(list, everySolution[1])
			fmt.Printf("Task '%s' added!\n", list[len(list)-1].Name)
		case "view":
			fmt.Printf("--- VIEW TASKS ---\n")
			viewAllTasks(list)
			completedCount := 0
			for counter := 0; counter < len(list); counter++ {
				if list[counter].Completed {
					completedCount++
				}
			}
			leftTasks := len(list) - completedCount
			fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(list), completedCount, leftTasks)
		case "complete":
			fmt.Printf("--- COMPLETE TASK ---\n")
			indexSlice, err := strconv.Atoi(everySolution[1])
			if err != nil {
				fmt.Printf("Invalid data:% v\n", err)
				return
			}
			if indexSlice < 0 || indexSlice >= len(list) {
				fmt.Printf("Invalid task number\n")
			} else {
				completeTask(&list, indexSlice)
				fmt.Printf("Task '%s' marked as completed!\n", list[indexSlice].Name)
			}
		case "remove":
			fmt.Printf("--- REMOVE TASK ---\n")
			indexDel, err := strconv.Atoi(everySolution[1])
			if err != nil {
				fmt.Printf("Invalid data: %v\n", err)
				return
			}
			if indexDel < 0 || indexDel >= len(list) {
				fmt.Printf("Invalid task number\n")
			} else {
				savedName := list[indexDel].Name
				list = removeTask(list, indexDel)
				fmt.Printf("Task '%s' removed successfully!\n", savedName)
			}
		case "exit":
			fmt.Printf("--- EXIT ---\n")
			fmt.Printf("Final list:\n")
			viewAllTasks(list)
			totalCount := len(list)
			completedCount := 0
			for i := 0; i < len(list); i++ {
				if list[i].Completed == true {
					completedCount++
				}
			}
			remainingCount := totalCount - completedCount
			fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", totalCount, completedCount, remainingCount)
			fmt.Printf("Goodbye, %s!\n", userName)
			return
		}
	}
}
