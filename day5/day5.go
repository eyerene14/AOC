package main

import (
	"adventofcode/user/hello/puzzle"
	"bufio"
	"fmt"
	"slices"
	//"regexp"
	"strconv"
	"strings"
	//"golang.org/x/exp/slices"
)

var mapName string

var seeds = []int{2906961955, 52237479, 1600322402, 372221628, 2347782594, 164705568, 541904540, 89745770, 126821306, 192539923, 3411274151, 496169308, 919015581, 8667739, 654599767, 160781040, 3945616935, 85197451, 999146581, 344584779}

//var seeds = []int{79, 14, 55, 13 }

type Map struct {
	seed   int
	soil int
}

type lookup_tbl struct {
	x int
	y int
	r int
}

type mapLines struct {
	name string
	lines []lookup_tbl
}

//var mapKeys=map[string]string{"seed-to-soil","soil-to-fertilizer","fertilizer-to-water","water-to-light","light-to-temperature","temperature-to-humidity","humidity-to-location"}

func LinesPerMap(line []string) lookup_tbl {
	y, _ := strconv.Atoi(line[0])
	x, _ := strconv.Atoi(line[1])
	r, _ := strconv.Atoi(line[2])
	map_as_tbl := lookup_tbl{x, y, r}
	return map_as_tbl
}

func ParseMap(tbl map[string][]lookup_tbl) []int{
	var a int
	var b bool
	var seed int
	htol:=make([]int,len(seeds))

	for i,s:= range seeds{
		seed=s
		a=s
		for _,v:= range tbl["humidity-to-location "]{
			a,b=CompareBetween(v,a)	
			if b{
				continue
			}
			//fmt.Println(a)
		}
		htol[i]=a
	}
	fmt.Println(seed,a,htol)
	return htol
}

func CompareBetween(tbl lookup_tbl, seed int) (int,bool){
	//fmt.Println("seed:", seed, " x:", tbl.x, ", y:", tbl.y, ", r:", tbl.r, " -- ")
	
	if seed > tbl.x && seed <= tbl.x+tbl.r{
		return seed + (tbl.y - tbl.x),true
	}else{
		return seed,true
	}
	//return seed,seed + (tbl.y - tbl.x)
}

func seedToSoil(line string, i int)Map{
	n:=Map{}
	if strings.Contains(line,"map:"){
		mapName = strings.Split(line, "map:")[0]
	}else{
		//nums:=strings.Fields(line)
		//m:=CompareBetween(nums, i)
	}
	return n
}


func ReadLines(scanner *bufio.Scanner) map[string][]lookup_tbl{
	//var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	linecount := 0
	tbls := make(map[string][]lookup_tbl)
	//sm := make(map[int][]int)
	//m:=Map{}
	var nums []string
	// Read each line and search for the first integer
	var lines []lookup_tbl
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line,"map:"){
			mapName = strings.Split(line, "map:")[0]
		}else{
			
			//d:=digitRegexp.FindAllString(line,-1)
			nums= strings.Fields(line)
			lines=append(lines,LinesPerMap(nums))
		}
		tbls[mapName]=lines
		
		linecount++

	}
	return tbls
}
func main() {
	scanner:=puzzle.ReadFile("./day5/day5data.txt")
	tables:= ReadLines(scanner)
	loc:=ParseMap(tables)
	fmt.Println(slices.Min(loc))
}
