package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main(){
	
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	 var firstdigit string
	 var lastdigit string
	 var temp int
	 var flag bool
	 numericWords := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
		
	}
	total:=0
	flag=true

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line :=scanner.Text()
		flag=true
		line = strings.ToLower(line)

		for word,value := range numericWords{
			line=strings.Replace(line,word,word[0:1]+string(value)+word[len(word)-1:len(word)],-1)
		}
		fmt.Println(line)
		for index:= range line{

			if unicode.IsDigit(rune(line[index]))  {
				if flag {
					firstdigit=string(line[index])
					lastdigit=string(line[index])
					flag=false;

				}else{
					lastdigit=string(line[index])
				}

				
			}
		}
		temp,_ =strconv.Atoi((firstdigit+lastdigit))
		
		
		total +=temp

		
	}
	fmt.Println("The total: ",total)

	
}