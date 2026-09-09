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

func removeTask(sliceTasks []Task, index int) []Task {
    return append(sliceTasks[:index], sliceTasks[index + 1:]...)
}

func main() {
    var testSlice []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scannedText := scanner.Text()
    count, err := strconv.Atoi(scannedText)
    if err != nil {
        fmt.Printf("Invalid count: %v\n", err)
        return
    }
    scanner.Scan()
    dividedText := strings.Split(scanner.Text(), ",")
    for cycle := 0; cycle < count; cycle++ {
        taskText := dividedText[cycle]
        taskData := strings.Split(taskText, ":")
        taskCompleted, err := strconv.ParseBool(taskData[1])
        if err != nil {
            fmt.Printf("Invalid data: %v\n", err)
            return
        }
        testSlice = append(testSlice, Task{taskData[0], taskCompleted})
    }
    scanner.Scan()
    taskIndex, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Printf("Invalid index: %v\n", err)
        return
    }
    savedName := testSlice[taskIndex].Name
    testSlice = removeTask(testSlice, taskIndex)
    completedCount := 0
    for i := 0; i < len(testSlice); i++ {
        if testSlice[i].Completed {
            completedCount++
            fmt.Printf("[x] %s\n", testSlice[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", testSlice[i].Name)
        }
    }
    totalCount := len(testSlice)
    leftTasks := totalCount - completedCount
    fmt.Printf("Task '%s' removed successfully!\n", savedName)
    fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", totalCount, completedCount, leftTasks)
}
