package main

import (
        "fmt"
        "bufio"
        "os"
        "strconv"
        "strings"
    )

type Task struct {
    Name string
    Completed bool
}

func completedTask(tasklist *[]Task, index int) {
    (*tasklist)[index].Completed = true
} 

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        lines := scanner.Text()
    }

    


}
