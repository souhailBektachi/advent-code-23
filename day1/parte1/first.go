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
	total:=0
	flag=true

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line :=scanner.Text()
		flag=true
		line = strings.ToLower(line)

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