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
    actions := strings.Split(scanner.Text(), ",")
    for i := 0; i < len(actions); i++ {
        everySolution := strings.Split(actions[i], "|")
        switch everySolution[0] {
            case "add":
            fmt.Printf("--- ADD TASK ---\n")
            list = addTask(list, everySolution[1])
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
            fmt.Printf("Total: %d tasks (%d completed, %d remaining)", len(list), completedCount, leftTasks)
            case "complete":
            fmt.Printf("--- COMPLETE TASK ---\n")
            completeTask()
            case "remove":
            fmt.Printf("--- REMOVE TASK ---\n")
            removeTask()
            case "exit":
            fmt.Printf("--- EXIT ---\n")
        }
    }
}
