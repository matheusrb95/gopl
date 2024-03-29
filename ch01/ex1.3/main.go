package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("echo1:")
	start := time.Now()
	echo1()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("elapsed time: %s\n", elapsed.String())

	fmt.Println("------\necho2:")
	start = time.Now()
	echo2()
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("elapsed time: %s\n", elapsed.String())

	fmt.Println("------\necho3:")
	start = time.Now()
	echo3()
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Printf("elapsed time: %s\n", elapsed.String())
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
