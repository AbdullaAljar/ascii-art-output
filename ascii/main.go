package main

import (
	"ascii"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 4 && len(os.Args) > 1 {
		var str string
		var output string
		if len(os.Args) == 4 {
			str = os.Args[2]
			flag.StringVar(&output, "output", "", "Specify the output file")
			flag.Parse()
			ascii.ToAsciiArtFile(str, output)
			os.Exit(0)
		} else {
			str = os.Args[1]
		}
		if strings.HasPrefix(os.Args[1], "--o") && len(os.Args) == 2 {
			fmt.Println("Error: please provide enough arguments\nUsage: go run . [OPTION] [STRING] [BANNER] ")
			os.Exit(0)
		}
		var words []string
		words = ascii.SeparateNewLine(str)
		for i := 0; i < len(words); i++ {
			if words[i] != "N3w L1n3" && words[i] != "" {
				ascii.ToAsciiArt(words[i])
			} else if words[i] == "" {
				fmt.Println()
			}
		}
	} else {
		fmt.Println("please provide maximum of 3 arguments and minimum of 1 argument\nUsage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
}
