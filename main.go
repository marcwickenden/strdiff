package main

import (
	"fmt"
	"os"
)

var (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <string1> <string2>", os.Args[0])
		os.Exit(1)
	}

	string1 := os.Args[1]
	string2 := os.Args[2]

	fmt.Println(string1)

	for idx, i := range string1 {
		if len(string2) < idx+1 {
			break
		}

		if rune(string2[idx]) != i {
			fmt.Printf(Red + string(i) + Reset)
		} else {
			fmt.Printf("%s", string(i))
		}
	}
	fmt.Println()

}