package main

import (
    "fmt"
    "strconv"
    "strings"
    "bufio"
    "os"
)

type Task struct {
    Name string
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
    return append(sliceTasks[:index], sliceTasks[index + 1:]...)
}

func main() {
    var list []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    listData := strings.Split(scanner.Text(), ",")
    appData := listData[0]
    userData := listData[1]
    scanner.Scan()
    secondString := scanner.Text()
    if secondString != "" {
        dividedString := strings.Split(secondString, ",")
        for count := 0; count < len(dividedString); count++ {
            tasks := dividedString[count]
            dividedTasks := strings.Split(tasks, ":")
            completedTag, err := strconv.ParseBool(dividedTasks[1])
            if err != nil {
                fmt.Printf("Invalid data: %v\n", err)
                return
            }
            list = append(list, Task{dividedTasks[0], completedTag})
        }
    }
    fmt.Printf("Welcome to %s, %s!\n", appData, userData)
    fmt.Printf("1. Add Task\n")
    fmt.Printf("2. View Tasks\n")
    fmt.Printf("3. Complete Task\n")
    fmt.Printf("4. Remove Task\n")
    fmt.Printf("5. Exit\n")
    fmt.Printf("Current tasks: %d\n", len(list))
    scanner.Scan()
    actions := strings.Split(scanner.Text(), ",")
    for i := 0; i < len(actions); i++ {
        currentAction := actions[i]
        dividedActions := strings.Split(currentAction, "|")
        switch dividedActions[0] {
        case "add":
            fmt.Printf("--- ADD TASK ---\n")
            list = addTask(list, dividedActions[1])
            fmt.Printf("Task '%s' added!\n", dividedActions[1])
        case "view":
            fmt.Printf("--- VIEW TASKS ---\n")
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
        case "complete":
            fmt.Printf("--- COMPLETE TASK ---\n")
            actionsDigit, err := strconv.Atoi(dividedActions[1])
            if err != nil {
                fmt.Printf("Invalid action: %v\n", err)
                return
            }
            if actionsDigit < 0 || actionsDigit >= len(list) {
                fmt.Printf("Invalid task number\n")
            } else {
                completeTask(&list, actionsDigit)
                fmt.Printf("Task '%s' marked as completed!\n", list[actionsDigit].Name)
            }
        case "remove":
            fmt.Printf("--- REMOVE TASK ---\n")
            actionsDigit, err := strconv.Atoi(dividedActions[1])
            if err != nil {
                fmt.Printf("Invalid action: %v\n", err)
                return
            }
            if actionsDigit < 0 || actionsDigit >= len(list) {
                fmt.Printf("Invalid task number\n")
            } else {
                taskName := list[actionsDigit].Name
                list = removeTask(list, actionsDigit)
                fmt.Printf("Task '%s' removed successfully!\n", taskName)
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
            fmt.Printf("Goodbye, %s!\n", userData)
            return
        }
    }
}
