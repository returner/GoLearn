package main

import (
	"fmt"
	"time"
	// "os"
	// "text/scanner"
)

func f(msg string, delay time.Duration) {
	for {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}

func main() {
	// var s scanner.Scanner
	// s.Init(os.Stdin)
	// for {
	// 	switch s.Scan() {
	// 	case scanner.EOF:
	// 		return //all done
	// 	case scanner.Ident:
	// 		fmt.Println(s.TokenText())
	// 	}
	// }
	go f("A--", 300*time.Millisecond)
	go f("-B-", 500*time.Millisecond)
	go f("--C", 1100*time.Millisecond)
	time.Sleep(20 * time.Second)
}
