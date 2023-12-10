package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	symbols := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "[", "]", "{", "}", ";", ":", ",", "<", ">", "/", "?", "`", "~"}

	var EngineArray2 [][]string
	scanner := bufio.NewScanner(file)
	ligneindex := 0
	for scanner.Scan() {

		temp2 := []string{}
		line := scanner.Text()
		EngineArray2 = append(EngineArray2, temp2)

		for _, char := range line {
			EngineArray2[ligneindex] = append(EngineArray2[ligneindex], string(char))

		}
		ligneindex++
	}
	total := 0
	indexI := 0
	indexJ := 0
	offsetI := 0
	offsetJ := 0
	temp := 0
	condition := false
	for I, array := range EngineArray2 {
		if len(array) == 0 {
			continue
		}
		for J, element := range array {
			condition = false
			if J != 0 && unicode.IsDigit(rune(element[0])) && !unicode.IsDigit(rune(array[J-1][0])) {
				condition = true
			} else if J == 0 && unicode.IsDigit(rune(element[0])) {
				condition = true
			}
			if condition {
				if I == 0 {
					indexI = I

				} else {
					indexI = I - 1
				}
				if I == len(EngineArray2)-1 {
					offsetI = I
				} else {
					offsetI = I + 1
				}

				if J == 0 {
					indexJ = J
				} else {
					indexJ = J - 1
				}

				tempj := J

				for unicode.IsDigit(rune(string(array[tempj])[0])) {
					if tempj < len(array)-1 {
						tempj++
					} else {
						break
					}

				}
				offsetJ = tempj

			loop:
				for i := indexI; i <= offsetI; i++ {
					for j := indexJ; j <= offsetJ; j++ {
						for _, symbol := range symbols {
							if symbol == EngineArray2[i][j] {
								tempj := J
								tempc := ""
								for unicode.IsDigit(rune(array[tempj][0])) {
									tempc += string(array[tempj])
									if tempj < len(array)-1 {
										tempj++
									} else {
										break
									}
									fmt.Println(tempj)

								}
								// fmt.Println(tempc)
								temp, _ = strconv.Atoi(tempc)
								total += temp
								break loop
							}
						}
					}
				}
			}

		}
	}

	fmt.Println("total: ", total)
}
