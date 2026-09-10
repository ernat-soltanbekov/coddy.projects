package main

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
    )

type Task struct {
    Name string
    Completed bool
}

func addTask(tasks []Task, taskName string) []Task {
    newTasks := append(tasks, Task{taskName, false})
    return newTasks
}

func viewAllTasks(tasks []Task) {
    for i := 0; i < len(tasks); i++ {
        if tasks[i].Completed {
            fmt.Printf("[x] %s\n", tasks[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", tasks[i].Name)
        }
    }
}

func completeTask(tasks *[]Task, index int) {
    if index >= len(*tasks) || index < 0 {
        fmt.Printf("Invalid task number\n")
        return
    }
    (*tasks)[index].Completed = true
}

func removeTask(tasks []Task, index int) []Task {
    return append(tasks[:index], tasks[index +1:]...)
}

func main() {
    var list []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanlist := strings.Split(scanner.Text(), ",")
    appName := scanlist[0]
    userName := scanlist[1]
    scanner.Scan()
    taskSlice := strings.Split(scanner.Text(), ",")
    for count := 0; count < len(taskSlice); count++ {
        taskData := strings.Split(taskSlice[count], ":")
        booledTask, err := strconv.ParseBool(taskData[1])
        if err != nil {
            fmt.Printf("Invalid value %v", err)
            return
        }
        list = append(list, Task{taskData[0], booledTask})
    }
    fmt.Printf("Welcome to %s, %s!\n", appName, userName)
    fmt.Printf("1. Add Task\n")
    fmt.Printf("2. View Tasks\n")
    fmt.Printf("3. Complete Task\n")
    fmt.Printf("4. Remove Task\n")
    fmt.Printf("5. Exit\n")
    fmt.Printf("Current tasks: %d\n", len(list))
    scanner.Scan()
    actions := strings.Split(scanner.Text(), "|")
    desicion := strings.Split(actions[0], "|")
    for i := 0; i < len(actions); i++ {
        switch desicion {
            case "add":
            fmt.Printf("--- ADD TASK ---\n")
            case "view":
            fmt.Printf("--- VIEW TASKS ---\n")
            case "complete":
            fmt.Printf("--- COMPLETE TASK ---\n")
            case "remove":
            fmt.Printf("--- REMOVE TASK ---\n")
            case "exit":
            fmt.Printf("--- EXIT ---\n")
        }
    }
}
