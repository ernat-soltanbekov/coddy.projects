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

func removeTask(slice []Task, index int) []Task {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
    var sliceTasks []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scannedText := scanner.Text()
    convertedText, err := strconv.Atoi(scannedText)
    if err != nil {
        fmt.Printf("Invalid digit: %v\n", err)
        return
    }
    scanner.Scan()
    scannedText = scanner.Text()
    secondString := strings.Split(scannedText, ",")
    for count := 0; count < convertedText; count++ {
        splitedText := strings.Split(secondString[count], ":")
        completed, err := strconv.ParseBool(splitedText[1])
        if err != nil {
            fmt.Printf("Invalid status: %v", err)
            return
        }
        sliceTasks = append(sliceTasks, Task{splitedText[0], completed})
    }
    scanner.Scan()
    convertedInt, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Printf("Invalid number: %v\n", err)
        return
    }
    savedName := sliceTasks[convertedInt].Name
    sliceTasks = removeTask(sliceTasks, convertedInt)
    finishedCount := 0
    for i := 0; i < len(sliceTasks); i++ {
        if sliceTasks[i].Completed == true {
            finishedCount++
            fmt.Printf("[x] %s\n", sliceTasks[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", sliceTasks[i].Name)
        }
    }
    fmt.Printf("Task '%s' removed successfully!\n", savedName)
    left := len(sliceTasks) - finishedCount
    fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(sliceTasks), finishedCount, left)
}

