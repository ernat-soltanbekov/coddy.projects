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

func viewAllTasks(parametr []Task) {
    for i := 0; i < len(parametr); i++ {
        if parametr[i].Completed == true {
            fmt.Printf("[x] %s\n", parametr[i].Name)
        } else {
            fmt.Printf("[ ] %s\n", parametr[i].Name)
        }
    }
}

func main() {
    var kusok []Task
    scanirovatel := bufio.NewScanner(os.Stdin)
    scanirovatel.Scan()
    otskanirovannyiText, oshibka := strconv.Atoi(scanirovatel.Text())
    if oshibka != nil {
        fmt.Println("Oshibka, cyka blyat.")
        return
    }
    scanirovatel.Scan()
    vtorayaStroka := scanirovatel.Text()
    razdelyennayaStroka := strings.Split(vtorayaStroka, ",")
    for i := 0; i < otskanirovannyiText; i++ {
        currentNote := razdelyennayaStroka[i]
        divideNote := strings.Split(currentNote, ":")
        boolDefine, mistake := strconv.ParseBool(divideNote[1])
        if mistake != nil {
            fmt.Println("Davai po-novoi, Miwa. Vse xyunia.")
        }
        kusok = append(kusok, Task{divideNote[0], boolDefine})
    }
    completedCount := 0
    for count := 0; count < len(kusok); count++ {
        if kusok[count].Completed == true {
            completedCount++ 
        }
    }
    viewAllTasks(kusok)
    fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", otskanirovannyiText, completedCount, otskanirovannyiText - completedCount)
}
