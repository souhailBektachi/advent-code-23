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


	total:=0;
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){

		line :=scanner.Text()
		linearray :=strings.Split(line,":")
		game:=linearray[0]
		gamenum,_:=strconv.Atoi(strings.Split(game," ")[1])
		maxmap:=map[string]int{
			"blue":0,"red":0,"green":0,
		}
		tempPower:=1;

			for _,game := range strings.Split(linearray[1],";"){
				
				for _,color:=range strings.Split(game,","){
					colorarray:= strings.Split(color," ")
					colors:=colorarray[2]
					colorNum,_ := strconv.Atoi(colorarray[1])
					
					if maxmap[colors]<colorNum{
						maxmap[colors]=colorNum
					}
		
				}
				
			}
			for color,max := range maxmap{
				fmt.Println(gamenum,color,maxmap[color])
				tempPower *=max
			}
			total +=tempPower
		
	}
	fmt.Println("total",total)
	
}