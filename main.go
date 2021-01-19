package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("./HRS <CONFIG-FILE> <INPUT-DATA-FILE>")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid number of arguments")
		usage()
	}

	rules, err := parseRules(os.Args[1])
	if err != nil {
		panic(err)
	}

	in, err := parseData(os.Args[2])
	if err != nil {
		panic(err)
	}

	for _, r := range rules {
		iPasses, err := r.check(in)
		if err != nil {
			panic(err)
		}
		for i, passes := range iPasses {
			if passes {
				fmt.Printf("%d@%d\n", r.ID, i)
			}
		}
	}
}
