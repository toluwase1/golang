package main

import "fmt"

func main() {

	var num int = 1
	var email string = "bob@example.com"
	concatenated:=fmt.Sprintf("%d%s", num, email)

	var name string = "Tolu"
	var squad string = "squad 8"
	var year int = 2021
	fmt.Println("My name is" + name)
	fmt.Println("My squad is: "+ squad)
	fmt.Println(concatenated)
	//fmt.Println("My year is "+ year)
	concatenated1 := fmt.Sprint("My name is ", name, " my squad is ", squad, " admitted in the year ", year)
	fmt.Println(concatenated1)
	fmt.Sprint(year, name, squad)
}

func goProgram(name string, squad string, year int)  {
	name = "Tolu"
	squad = "squad 08"
	year = 2021
}
