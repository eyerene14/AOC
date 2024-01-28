package main

import (
	"fmt"
)

//TEST
//var mt=[...]int{7,15,30}
//var md=[...]int{9,40,200}

//REAL
//Time:        35     93     73     66
//Distance:   212   2060   1201   1044
var mt=[...]int{35,93,73,66}
var md=[...]int{212,2060,1201,1044}

var total=0

func mltp(num int){

}

func iterRace() int{
	var tc [4]int
	for i:=0; i < 4; i++{
		tc[i]=calcDistance(i)	
	}
		fmt.Print(tc[0],tc[1],tc[2],tc[3])
	
	return tc[0]*tc[1]*tc[2]*tc[3]
}

func calcDistance(i int) int{
	var maxT,maxD=mt[i],md[i]
	var distance int
	count:=0
	m1,m2:=maxT,0
	for ;m1 >= 0;m1-- {
		distance=m1*m2
		m2++
		fmt.Println(distance)

		if distance > maxD{
			count++
		}else{
			continue
		}
	}
	return count
}

func main(){
	rt:=iterRace()
	fmt.Println("Count:", rt)
}
