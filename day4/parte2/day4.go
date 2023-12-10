package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)
func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	total :=0
	copiesmap:=make(map[int]int)
	cardindex:=1;
	scanner:=bufio.NewScanner(file)
		for scanner.Scan(){
			line :=scanner.Text()
			array:=strings.Split(line,":")
			array[1]=strings.ReplaceAll(array[1]," ",",")
			array[1]=strings.ReplaceAll(array[1],",,",",")
			
			array=strings.Split(array[1],"|");
			mymap :=make(map[string]int)
			wining:= strings.Split(array[0],",")
			card:= strings.Split(array[1],",")
			pts:=0;
			for _,key := range wining{
			
					if len(key)>=1 {
						mymap[key]=1;


					}
			}
			var copiesIndex int
			var Index2 int 
			if copiesmap[cardindex]==0 {
				copiesIndex=1
				Index2=cardindex


			}else{
				copiesmap[cardindex]+=1
				copiesIndex=copiesmap[cardindex]
				Index2=cardindex+1
			}
			for i := 0; i < copiesIndex; i++ 	{
				pts=0
			for _,element := range card{
				_,exist:=mymap[element]
				if exist{
					pts +=1
				}
			}
			for i := Index2; i <=cardindex+ pts; i++ {
				
				copiesmap[i]+=1;
				
			}
			
		}

		cardindex++
		}
		for key,element := range copiesmap{
			if key ==cardindex {
				break
			}
			total +=element
		}
		fmt.Println(copiesmap)
		fmt.Println(total)
		}

	
