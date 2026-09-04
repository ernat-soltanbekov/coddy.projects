package main

import "fmt"

func main() {
    // Чтение входных данных
    var appName string
    var userName string
    fmt.Scanln(&appName)
    fmt.Scanln(&userName)
    
    // TODO: Напишите свой код ниже
    // Определение структуры Task
    type Task struct {
        Name string
        Completed bool
    }
    // Создание среза структур Task
    var tasks []Task
    // Отображение приветственного сообщения и меню
    fmt.Printf("Welcome to %s, %s!\n", appName, userName)
    fmt.Println("1. Add Task")
    fmt.Println("2. View Tasks")
    fmt.Println("3. Complete Task")
    fmt.Println("4. Remove Task")
    fmt.Println("5. Exit")
    // Вывод текущего количества задач
    fmt.Printf("Current tasks: %d\n", len(tasks))
}
