package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	Time   int
	Record int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Races := make([]race, 0)
	NumWays := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		TempArray := strings.Split(line, ":")
		TempArray[1] = strings.Trim(TempArray[1], " ")
		line = strings.Join(strings.Fields(TempArray[1]), " ")
		for index, element := range strings.Split(line, " ") {
			elementInt, _ := strconv.Atoi(element)
			if len(Races) <= index {
				Races = append(Races, race{elementInt, 0})

			} else {
				Races[index].Record = elementInt
			}
		}

	}
	fmt.Println(Races)

	for index, element := range Races {
		NumWays = append(NumWays, element.Record)
		min := 0
		max := 0
		for i := 0; i < element.Time; i++ {
			if ((element.Time - i) * i) > element.Record {
				min = i
				break

			}
		}
		for i := element.Time; i > min; i-- {
			if (element.Time-i)*i > element.Record {
				max = i
				break

			}
		}
		NumWays[index] = max - min + 1
	}
	fmt.Println(NumWays)
	result := 1
	for _, element := range NumWays {
		result *= element
	}
	fmt.Println(result)
}
