package main

import (
	"fmt"
)

var mt=[...]int{7,15,30}
var md=[...]int{9,40,200}
var maxT int

func iterRace()int{
	var mid int
	for i:=0; i < 3; i++{
		maxT=mt[i]
		mid=maxT+1/2
	}
	return mid
}

func calcDistance(m int) int{
	var distance int
	count:=1
	m1,m2:=maxT-1,1
	for ;m > 0;m-- {
		distance=m1*m2
		m2++
		m1--
		fmt.Println(distance)

		if distance < maxT{
			continue
		}else{
			count++
		}
	}
	return count
}

func main(){
	maxT=mt[2]
	fmt.Println(calcDistance((maxT+1)/2))
}
