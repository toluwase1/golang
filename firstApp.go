package main

import (
	"fmt"
	"sort"
)
func main() {
	var varName string = "toluwase";
	fmt.Println(varName)
	if len(varName)==4 {
		fmt.Println("The length is 4")
	} else {
		fmt.Println("This is less or greater than 4")
	}
	var x int = 1
	y:=2
	fmt.Println(x+y)
fmt.Println("_____for range_____")
	stringArray := [] string{"a", "b", "c", "d"}

	for i, c:=range stringArray {
		fmt.Println("index = ", i, "value of c = ", c)
		fmt.Println("value of c = ",c)
	}

	fmt.Println("_____Looping through MAPS to dispay key and value_____")
	maps := map [string] int  {"key1": 1, "key2":2, "key3":3, "key4":4}
	for k , v :=range maps {
		fmt.Println("key=", k, "values = ", v )
	}

	fmt.Println("____looping through maps to display keys only____")
	for  k := range maps {
		fmt.Println("keys only", k)
	}

	fmt.Println("____looping through maps to display values only____")
	for v:=range maps{
		fmt.Println("printing only values", v);
	}

	abc:=make(map[string]int)
	abc["toluwase"] = 27
	abc["thomas"] =167
	fmt.Println(abc["toluwase"])
	fmt.Println(abc)

	i, j := 42, 2701
	fmt.Println(&i,& j)

	p,q:=&i, &j
	fmt.Println(p, q)
	fmt.Println(*p)
	*q=*q/37
	//fmt.Println(&p)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(*q)
	fmt.Println(j/37)

squareInputs(20)
//fmt.Println(a)
	fmt.Println(calculateArea(2,4))
	calculateArea(2,5)
	//var car Cars =Cars{"blue", 4. "toyo"}
	carr:=Cars{"dgdg",67,"rtr"}
	fmt.Println(carr)

	ca:= Cars{"jkngjk", 67, ""}
	fmt.Println(ca.color)
	fmt.Println(ca.wheel)
	fmt.Println(ca.name)
	fmt.Println(ca)
}
func TwoOldestAges(ages []int) [2]int {
	//
	sort.Ints(ages)
	return [2]int{}
}

func squareInputs(age int)  {
	ageValue := age*age
	fmt.Println(ageValue)
}

func myFunction(age, year int) (int, int)  {
	age*=age
	year-=2021
	return age,year
}

func functionName(x, y int)(int, int){
	return x*3, y*7
}

func calculateArea(x int, y int)( rectangleArea int, squareArea int )  {
		rectangleArea = x*y
		squareArea = x*x
		return
}

//structs
type Cars struct {
	color string
	wheel int
	name string
}
