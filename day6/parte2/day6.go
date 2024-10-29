package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getWays(Time int64, Record int64) int64 {
	min := int64(0)
	max := int64(0)

	for i := int64(0); i < Time; i++ {
		if ((Time - i) * i) > Record {
			min = i
			break

		}
	}
	for i := Time; i > min; i-- {
		if (Time-i)*i > Record {
			max = i
			break

		}
	}
	return max - min + int64(1)
}

type race struct {
	Time   int64
	Record int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Race := race{0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		TempArray := strings.Split(line, ":")
		TempArray[1] = strings.ReplaceAll(TempArray[1], " ", "")
		fmt.Println(TempArray[1])
		if Race.Time == 0 {
			Race.Time, _ = strconv.ParseInt(TempArray[1], 10, 64)

		} else {
			Race.Record, _ = strconv.ParseInt(TempArray[1], 10, 64)
		}

	}

	result := getWays(int64(Race.Time), int64(Race.Record))
	fmt.Println(Race)
	fmt.Println(result)
}
