package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func fillNextMap(MapUsed map[string]string, NexMap map[string]string) {
	for _, element := range MapUsed {
		NexMap[element] = element
	}
}
func DestinationToSource(line string, DestOrg map[string]string) {

	tempArray := strings.Split(line, " ")
	Range, _ := strconv.ParseInt(tempArray[2], 10, 64)
	Destination, _ := strconv.ParseInt(tempArray[0], 10, 64)
	Source, _ := strconv.ParseInt(tempArray[1], 10, 64)
	offset := int64(0)

	if Destination > Source {
		offset = Destination - Source
	} else if Source > Destination {
		offset = Source - Destination
		offset *= -1
	}
	for element := range DestOrg {

		SourceNum, _ := strconv.ParseInt(element, 10, 64)
		if SourceNum < Source+Range && SourceNum >= Source {
			DestOrg[strconv.FormatInt(SourceNum, 10)] = strconv.FormatInt(SourceNum+offset, 10)

		}

	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	MapUsed := make(map[string]string)

	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			continue
		}
		if len(MapUsed) == 0 {
			tempArray := strings.Split(line, ":")
			tempArray = strings.Split(tempArray[1], " ")[1:]
			for _, element := range tempArray {
				MapUsed[element] = element

			}
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			fmt.Println(line)
			tempMap := make(map[string]string)
			fillNextMap(MapUsed, tempMap)
			MapUsed = tempMap

			continue
		}

		if unicode.IsDigit(rune(line[0])) {
			DestinationToSource(line, MapUsed)

		}

	}
	fmt.Println(MapUsed)
	lowLocation := math.MaxInt
	for _, element := range MapUsed {

		LocationInt, _ := strconv.Atoi(element)
		if lowLocation > LocationInt {

			lowLocation = LocationInt

		}
	}

	fmt.Println("low location", lowLocation)
}
