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

func completeTask(sliceStruct *[]Task, index int) {
    (*sliceStruct)[index].Completed = true
}

func main() {
    var sliceTasks []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scannedText := scanner.Text()
    scannedInt, err := strconv.Atoi(scannedText)
    if err != nil {
        fmt.Printf("Invalid integer: %v", err)
        return
    }
    scanner.Scan()
    scannedText = scanner.Text()
    secondString := strings.Split(scannedText, ",")
    for i := 0; i < scannedInt; i++ {
        secondSplit := strings.Split(secondString[i], ":")
        boolean, err := strconv.ParseBool(secondSplit[1])
        if err != nil {
            fmt.Printf("Invalid status: %v\n", err)
            return
        }
        sliceTasks = append(sliceTasks, Task{secondSplit[0], boolean})
    }
    scanner.Scan()
    indexInt, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Printf("Invalid index: %v\n", err)
        return
    }
    completeTask(&sliceTasks, indexInt)
    completedCount := 0
    for i := 0; i < len(sliceTasks); i++ {
        if sliceTasks[i].Completed {
            completedCount++
            fmt.Printf("[x] %s\n", sliceTasks[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", sliceTasks[i].Name)
        }
    }
    remaining := len(sliceTasks) - completedCount
    fmt.Printf("Task '%s' marked as completed!\n", sliceTasks[indexInt].Name)
    fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(sliceTasks), completedCount, remaining)
}
