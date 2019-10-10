package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("usage:")
	fmt.Println("to indent with tabs: indent +")
	fmt.Println("to unindent with 4 spaces: indent - 4")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(0)
	}

	op := args[0]
	var spaces bool
	var amount int
	if len(args) == 2 {
		spaces = true
		thisAmount, err := strconv.Atoi(args[1])
		amount = thisAmount
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			os.Exit(1)
		}
	} else {
		spaces = false
		amount = 1
	}

	if op != "+" && op != "-" {
		fmt.Println("operator must be '+' or '-'!")
		usage()
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		dieAfterwards := false
		if err == io.EOF {
			dieAfterwards = true
		}
		if err != nil && err != io.EOF {
			fmt.Printf("err: %s\n", err.Error())
			os.Exit(1)
		}

		if op == "+" {
			var char string
			if spaces {
				char = " "
			} else {
				char = "\t"
			}

			for i := 0; i < amount; i++ {
				fmt.Printf(char)
			}
			fmt.Printf("%s", line)
		}

		if op == "-" {
			canUnindent := true
			for i := 0; i < len(line) && i < amount; i++ {
				if line[i] != ' ' && line[i] != '\t' {
					canUnindent = false
					break
				}
			}

			if canUnindent {
				fmt.Printf("%s", line[amount:])
			} else {
				fmt.Printf("%s", line)
			}
		}

		if dieAfterwards {
			break
		}

	}
}
