package main

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "os"
    )

// Определение структуры Task
type Task struct {
    Name string
    Completed bool
}

func addTask(newTask []Task, nameTask string) []Task {
    newTask = append(newTask, Task{nameTask, false})
    return newTask
}

// TODO: Напишите свой код ниже

func main() {
    // Создание среза структур Task
    var tasks []Task
    // Отображение приветственного сообщения и меню
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    countTask, err := strconv.Atoi(scanner.Text())

    if err != nil {
        fmt.Println("Invalid number")
        return
    }

    for i := 0; i < countTask; i++ {
        tasks = append(tasks, Task{fmt.Sprintf("Existing Task %d", i+1), false})
    }
    // Вывод текущего количества задач
    if scanner.Scan() {
        taskNames := scanner.Text()
        totalTasks := strings.Split(taskNames, ",")
        for t := 0; t < len(totalTasks); t++ {
            tasks = addTask(tasks, totalTasks[t])
        }
    }
    for print := 0; print < len(tasks); print++ {
        fmt.Printf("Task: %s, Completed: %t\n", tasks[print].Name, tasks[print].Completed)
    }
    fmt.Printf("Total tasks: %d\n", len(tasks))
}
