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

type Interval struct {
	Min int64
	Max int64
}

func fillNextMap(MapUsed map[Interval]Interval, NexMap map[Interval]Interval) map[Interval]Interval {
	for _, element := range MapUsed {
		NexMap[element] = element
	}
	return NexMap
}
func getIntervals(line string, InterVals []map[Interval]Interval, index int) {
	tempArray := strings.Split(line, " ")
	Range, _ := strconv.ParseInt(tempArray[2], 10, 64)
	Destination, _ := strconv.ParseInt(tempArray[0], 10, 64)
	Source, _ := strconv.ParseInt(tempArray[1], 10, 64)
	SourceInterval := Interval{Source, Source + (Range - 1)}
	DestinationInterval := Interval{Destination, Destination + (Range - 1)}
	if len(InterVals[index]) == 0 {
		InterVals[index] = make(map[Interval]Interval)
	}
	InterVals[index][SourceInterval] = DestinationInterval
}

func DestinationToSource(interval map[Interval]Interval, MapUsed map[Interval]Interval) map[Interval]Interval {
	tempmap := make(map[Interval]Interval)
	for NewSource := range MapUsed {
		overlaps := false

		for Source, Destination := range interval {
			if NewSource.Max < Source.Min || NewSource.Min > Source.Max {
				continue
			} else {
				overlaps = true

				if NewSource.Min > Source.Min {
					if NewSource.Max < Source.Max {
						tempmap[NewSource] = Interval{Destination.Min + (NewSource.Min - Source.Min), Destination.Max - (Source.Max - NewSource.Max)}
					} else {
						tempmap[Interval{NewSource.Min, Source.Max}] = Interval{Destination.Min + (NewSource.Min - Source.Min), Destination.Max}
					}
				} else {
					if NewSource.Max < Source.Max {
						tempmap[Interval{Source.Min, NewSource.Max}] = Interval{Destination.Min, Destination.Max - (Source.Max - NewSource.Max)}
					} else {
						tempmap[Interval{Source.Min, Source.Max}] = Interval{Destination.Min, Destination.Max}
					}
				}
			}
		}

		if !overlaps {
			tempmap[NewSource] = NewSource
		}
	}
	tempmap2 := make(map[Interval]Interval)
	for key := range MapUsed {
		overlaps := false

		for key2 := range tempmap {
			if key2.Min > key.Min && key2.Min <= key.Max {
				tempmap2[Interval{key.Min, key2.Min - 1}] = Interval{key.Min, key2.Min - 1}
				overlaps = true
			}
			if key2.Max >= key.Min && key2.Max < key.Max {
				tempmap2[Interval{key2.Max + 1, key.Max}] = Interval{key2.Max + 1, key.Max}
				overlaps = true
			}
		}

		if !overlaps {
			tempmap2[key] = key
		}
		delete(MapUsed, key)
	}
loop:
	for key, element := range tempmap2 {

		_, exist := tempmap[key]

		if !exist {
			for key2 := range tempmap {
				if key.Min == key2.Min || key.Max == key2.Max {
					continue loop
				}

			}
			tempmap[key] = element

		}

	}

	if len(tempmap) == 0 {
		return MapUsed

	}
	return tempmap
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	Seeds := make([]Interval, 0)
	MapUsed := make(map[Interval]Interval)

	InterVals := make([]map[Interval]Interval, 7)
	index := -1

	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			continue
		}

		if len(Seeds) == 0 {

			tempArray := strings.Split(line, ":")
			tempArray = strings.Split(tempArray[1], " ")[1:]

			for i, element := range tempArray {
				if (i+1)%2 != 0 {
					elementInt, _ := strconv.ParseInt(element, 10, 64)

					Seeds = append(Seeds, Interval{elementInt, 0})

				} else {
					tempInt, _ := strconv.ParseInt(element, 10, 64)
					Seeds[len(Seeds)-1].Max = tempInt + Seeds[len(Seeds)-1].Min - 1

				}
			}
			for _, seed := range Seeds {
				MapUsed[seed] = seed

			}
			continue
		}

		if !unicode.IsDigit(rune(line[0])) {
			fmt.Println(line)

			index++
			continue
		}

		if unicode.IsDigit(rune(line[0])) {

			getIntervals(line, InterVals, index)
			continue
		}

	}

	for i := 0; i < len(InterVals); i++ {
		MapUsed = DestinationToSource(InterVals[i], MapUsed)
		tempMap := make(map[Interval]Interval)
		if i < len(InterVals)-1 {
			MapUsed = fillNextMap(MapUsed, tempMap)
		}
	}
	lowDest := int64(math.MaxInt64)
	for _, dest := range MapUsed {

		if dest.Min < lowDest {
			lowDest = dest.Min

		}
	}
	// fmt.Println("Seeds: ", Seeds)
	// fmt.Println("intervals: ", InterVals)

	// fmt.Println("MapUsed: ", MapUsed)

	fmt.Println("lowDest: ", lowDest)
}
