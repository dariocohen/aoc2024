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
        skippedUnsafe := false
        for i := 0; i < len(reportStr); i++ {
            level, err := strconv.Atoi(reportStr[i])
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
                    if skippedUnsafe == true {
                        safe = false
                        break
                    } else {
                        skippedUnsafe = true
                        i = -1
                        prevDirection = nil
                        prevLevel = nil
                        continue
                    }
                }
                if prevDirection != nil {
                    if direction != *prevDirection {
                        if skippedUnsafe == true {
                            safe = false
                            break
                        } else {
                            skippedUnsafe = true
                            i = -1
                            prevDirection = nil
                            prevLevel = nil
                            continue
                        }
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
