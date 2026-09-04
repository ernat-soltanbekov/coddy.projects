package main

import (
        "fmt"
        "bufio"
        "os"
        "strconv"
        "strings"
    )

func addTask(slice []Task, stroka string) []Task {
    newSlice := append(slice, Task{stroka, false})
    return newSlice
}

type Task struct {
    Name string
    Completed bool
}

func main() {
    var tasks []Task
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    scanText := scanner.Text()
    countTask, _ := strconv.Atoi(scanText)
    for i := 0; i < countTask; i++ {
        taskName := fmt.Sprintf("Existing Task %d", i + 1)
        tasks = addTask(tasks, taskName)
    }
    hasInput := scanner.Scan()
    if hasInput {
        finalScan := scanner.Text()
        splStr := strings.Split(finalScan, ",")
        for _, element := range splStr {
            tasks = addTask(tasks, element)
        }
    }

    for e := 0; e < len(tasks); e++ {
        task := tasks[e]
        fmt.Printf("Task: %s, Completed: %t\n", task.Name, task.Completed)
        }
        fmt.Printf("Total tasks: %d\n", len(tasks))
}
