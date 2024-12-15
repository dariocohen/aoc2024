package main

import (
    "strings"
    "strconv"
    "bufio"
    "fmt"
    "log"
    "slices"
    "os"
)

func main() {
    fi, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer fi.Close()

    scanner := bufio.NewScanner(fi)
    var leftValues []int
    var rightValues []int
    for scanner.Scan() {
        pairFromInput := strings.Split(scanner.Text(), "   ")
        leftInt, err := strconv.Atoi(pairFromInput[0])
        if err != nil {
            log.Fatal(err)
        }

        rightInt, err := strconv.Atoi(pairFromInput[1])
        if err != nil {
            log.Fatal(err)
        }

        leftValues = append(leftValues, leftInt)
        rightValues = append(rightValues, rightInt)
    }
    slices.Sort(leftValues)
    slices.Sort(rightValues)
    var acum = 0
    var score = 0
    for i := 0; i < len(leftValues); i++ {
        acum = acum + absInt(leftValues[i] - rightValues[i])
        var leftValue = leftValues[i]
        var numberOfRightCoincidences = 0
        for j := 0; j < len(rightValues); j++ {
            if leftValue == rightValues[j] {
                numberOfRightCoincidences++
            }
        }
        score = score + (numberOfRightCoincidences * leftValue)
    }

    fmt.Println(acum)
    fmt.Println(score)

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
