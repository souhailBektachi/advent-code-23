package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	gamemap :=map[string]int{
		"blue":14,
		"red":12,
		"green":13,
		
		
	}
	total:=0;
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		 possible :=false

		line :=scanner.Text()
		linearray :=strings.Split(line,":")
		game:=linearray[0]
		gamenum,_:=strconv.Atoi(strings.Split(game," ")[1])
		
		firstloop:
			for _,game := range strings.Split(linearray[1],";"){
				
				for _,color:=range strings.Split(game,","){
					colorarray:= strings.Split(color," ")
					colors:=colorarray[2]
					colorNum,_ := strconv.Atoi(colorarray[1])
					if colorNum>gamemap[colors] {
						possible = false;
						break firstloop
						
					}else{
						possible =true;
					}
				}

			}
			if possible {
				total +=gamenum
				
			}
	}
	fmt.Println("total",total)
	
}