package main

import "fmt"

func deepest() {
	fmt.Println("deepest")
}

func recursiveCall(depth int) {
	if depth == 0 {
		deepest()
		return
	}
	recursiveCall(depth - 1)
}

func main() {
	recursiveCall(1000)
}
