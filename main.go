package main

import (
	"fmt"
	//"strconv"
	"strings"
	"time"
)

func main()  {
	//go process("order")
	//go  process("refund")
	//fmt.Scanln()
	fmt.Println(StatusAccepted)

	z:= Point{12, 34, true}
	var abce = Point{1, 2, false}
	//abce:= Point{a:2}

	fmt.Println(abce.c)
	fmt.Println(abce)
	fmt.Println(z.a)

	var sb strings.Builder

	for i := 0; i < 10; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())
}

func process(item string)  {
	for i:=1; true; i++ {
		time.Sleep(time.Second/2)
		fmt.Println("processed", i, item)
	}
}

const Pi = 3.14
const (
	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
)

type Point struct {
	a int
	b int
	c bool
}

/**
Primitive data types in Golang
In golang, we have int, float, byte, string,
rune & bool(boolean if you are coming from other programming languages).
 */



func MakeNegative(x int) int {
	if x*1==x {
		return x*-1
	} else {
		return -x
	}
	return 0
}
