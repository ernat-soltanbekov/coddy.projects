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

func completeTask(tasklist *[]Task, index int) {
	(*tasklist)[index].Completed = true
}

func viewAllTasks(tasklist []Task) {
	for i := 0; i < len(tasklist); i++ {
		if tasklist[i].Completed == true {
			fmt.Printf("[x] %s\n", tasklist[i].Name)
		} else {
			fmt.Printf("[ ] %s\n", tasklist[i].Name)
		}
	}
}

func main() {
	var tasklist []Task

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	taskCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	scanner.Scan()
	taskString := scanner.Text()

	dividedString := strings.Split(taskString, ",")

	for i := 0; i < taskCount; i++ {
		currentTask := dividedString[i]
		dividedTask := strings.Split(currentTask, ":")

		completed, err := strconv.ParseBool(dividedTask[1])
		if err != nil {
			return
		}

		tasklist = append(tasklist, Task{
			Name:      dividedTask[0],
			Completed: completed,
		})
	}

	scanner.Scan()
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	completeTask(&tasklist, index)

	viewAllTasks(tasklist)

	fmt.Printf("Task '%s' marked as completed!\n", tasklist[index].Name)

	completedCount := 0

	for i := 0; i < len(tasklist); i++ {
		if tasklist[i].Completed == true {
			completedCount++
		}
	}

	fmt.Printf(
		"Total: %d tasks (%d completed, %d remaining)\n",
		taskCount,
		completedCount,
		taskCount-completedCount,
	)
}
