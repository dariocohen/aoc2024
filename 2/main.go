package main

import (
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    fi, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fi.Close()

    scanner := bufio.NewScanner(fi)
    safeReports := 0
    for scanner.Scan() {
        reportStr := strings.Split(scanner.Text(), " ")
        var prevLevel *int
        var prevDirection *int
        safe := true
        for _, i := range reportStr {
            level, err := strconv.Atoi(i)
            if err != nil {
                panic(err)
            }
            if prevLevel != nil {
                diff := level - *prevLevel
                var direction int
                if diff < 0 {
                    direction = 1
                } else {
                    direction = -1
                }

                absDiff := absDiffInt(level,*prevLevel) 
                if absDiff > 3 || absDiff < 1 {
                    safe = false
                    break
                }
                if prevDirection != nil {
                    if direction != *prevDirection {
                        safe = false
                        break
                    }
                }

                prevDirection = &direction
            }

            prevLevel = &level
        }
        if safe == true {
            safeReports++
        }

    }

    fmt.Println(safeReports)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func absInt(x int) int {
   return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
   if x < y {
      return y - x
   }
   return x - y
}
