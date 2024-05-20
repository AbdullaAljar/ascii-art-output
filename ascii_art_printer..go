package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ToAsciiArt converts a string to ASCII art and prints it to the console.
func ToAsciiArt(str string) {
	var filename string

	// Check if there are command-line arguments
	if len(os.Args) == 3 {
		// Get the feature argument from the command line
		feature := strings.ToLower(os.Args[2])

		// Determine the filename based on the feature
		switch feature {
		case "shadow":
			filename = "shadow.txt"
		case "thinkertoy":
			filename = "thinkertoy.txt"
		case "standard":
			filename = "standard.txt"
		default:
			fmt.Println("Error: please check your arguments :/\nUsage: go run . [OPTION] [STRING] [BANNER] ")
			os.Exit(0)
		}

		// Open the chosen ASCII art file
		file, _ := os.Open(filename)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		lineNumber := 1
		runes := []rune(str)

		// Loop through the lines of ASCII art
		for y := 0; y < 8; y++ {
			for i := 0; i < len(runes); i++ {
				lineNumber = 0
				ValueOfletter := int(runes[i])
				lineToPrint := (ValueOfletter-33)*9 + 11 + y

				// Scan lines from the file and print the matching line
				for scanner.Scan() {
					lineNumber++
					if lineNumber == lineToPrint {
						line := scanner.Text()
						fmt.Print(line)
					}
				}

				// Reset the file scanner to the beginning
				file.Seek(0, 0)
				scanner = bufio.NewScanner(file)
			}
			fmt.Println()
		}
	} else {
		// If no feature specified, use the "standard" ASCII art file
		file, _ := os.Open("standard.txt")
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		lineNumber := 1
		runes := []rune(str)

		// Loop through the lines of ASCII art
		for y := 0; y < 8; y++ {
			for i := 0; i < len(runes); i++ {
				lineNumber = 0
				ValueOfletter := int(runes[i])
				lineToPrint := (ValueOfletter-33)*9 + 11 + y

				// Scan lines from the file and print the matching line
				for scanner.Scan() {
					lineNumber++
					if lineNumber == lineToPrint {
						line := scanner.Text()
						fmt.Print(line)
					}
				}

				// Reset the file scanner to the beginning
				file.Seek(0, 0)
				scanner = bufio.NewScanner(file)
			}
			fmt.Println()
		}
	}
}
