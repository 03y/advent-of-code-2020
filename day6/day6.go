package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    input, _ := readLines("input.txt")
    input = append(input, "")
    answers, values:= make(map[string]bool), []int{}

    for i := range input {
        if input[i] != "" {
            for j := range input[i] {
                 answers[string(input[i][j])] = true
                 
            }
        } else {
            values = append(values, len(answers))
            answers = make(map[string]bool)
        }
    }

    fmt.Println("Part 1:", sum(values)) 
}

func deepCopyMap(original map[string]bool) map[string]bool {
    newMap := make(map[string]bool)
    for key, value := range original {
        newMap[key] = value
    }
    return newMap
}

func sum(arr []int) int {
    c := 0
    for _, n := range arr {
        c += n
    }
    return c
}

func readLines(path string) ([]string, error) {
    file, _ := os.Open(path)
    defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}
