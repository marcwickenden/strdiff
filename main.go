package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

var (
	Red      = "\033[31m"
	Reset    = "\033[0m"
	printHex bool
)

func main() {
	flag.BoolVar(&printHex, "hex", false, "Show hex of strings")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Printf("usage: %s <string1> <string2>", os.Args[0])
		os.Exit(1)
	}

	string1 := flag.Arg(0)
	string2 := flag.Arg(1)

	if printHex {
		fmt.Println(hex.EncodeToString([]byte(string1)))
	} else {
		fmt.Println(string1)
	}

	for idx, i := range string1 {
		if len(string2) < idx+1 {
			break
		}

		character := string2[idx]

		if rune(character) != i {
			if printHex {
				fmt.Printf("%s%s%s", Red, hex.EncodeToString([]byte{character}), Reset)
			} else {
				fmt.Printf("%s%U%s", Red, character, Reset)
			}
		} else {
			if printHex {
				fmt.Printf("%s", hex.EncodeToString([]byte{character}))
			} else {
				fmt.Printf("%U", character)
			}
		}

	}
	fmt.Println()

}
