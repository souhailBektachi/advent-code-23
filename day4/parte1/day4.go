package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, ":")
		array[1] = strings.ReplaceAll(array[1], " ", ",")
		array[1] = strings.ReplaceAll(array[1], ",,", ",")

		array = strings.Split(array[1], "|")
		mymap := make(map[string]int)
		wining := strings.Split(array[0], ",")
		card := strings.Split(array[1], ",")
		pts := 0
		for index, key := range wining {

			if len(key) >= 1 {
				mymap[key] = index

			}

		}
		for _, element := range card {
			_, exist := mymap[element]
			if exist {
				if pts == 0 {
					pts = 1
				} else {
					pts *= 2
				}
			}
		}
		total += pts

	}
	fmt.Println("total:", total)
}
