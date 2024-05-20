package ascii

import (
	"strings"
	"os"
	"fmt"
	"bufio"
)

// ToAsciiArtFile takes a string and converts it into ASCII art based on the specified feature,
// then saves the result to an output file.
func ToAsciiArtFile(str string, output string) {
	var filename string

	// Get the feature argument from command line
	feature := strings.ToLower(os.Args[3])
	switch feature {
	case "shadow":
		filename = "shadow.txt"
	case "thinkertoy":
		filename = "thinkertoy.txt"
	case "standard":
		filename = "standard.txt"
	default:
		fmt.Println("Error: please check your arguments :)\nUsage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}

	// Split the input string into words
	words := SeparateNewLine(str)

	// Create or open the output file for writing
	filev, err := os.Create(output)
	if err != nil {
		fmt.Println("Error: failed to create output file:", err)
		os.Exit(1)
	}
	defer filev.Close()

	for i := 0; i < len(words); i++ {
		// Check for special marker "N3w L1n3" to handle newlines
		if words[i] != "N3w L1n3" && words[i] != "" {
			// Open the ASCII art file
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println("Error: failed to open file:", err)
				os.Exit(1)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			lineNumber := 1
			runes := []rune(words[i])

			for y := 0; y < 8; y++ {
				for i := 0; i < len(runes); i++ {
					lineNumber = 0
					ValueOfletter := int(runes[i])
					lineToPrint := (ValueOfletter-33)*9 + 11 + y
					for scanner.Scan() {
						lineNumber++
						if lineNumber == lineToPrint {
							line := scanner.Text()
							_, err := filev.WriteString(line)
							if err != nil {
								fmt.Println("Error: failed to write to output file:", err)
								os.Exit(1)
							}
						}
					}
					// Reset the file scanner to the beginning
					file.Seek(0, 0)
					scanner = bufio.NewScanner(file)
				}
				// Write a newline character after each line of ASCII art
				_, _ = filev.WriteString("\n")
			}
		} else if words[i] == "" {
			// Write a newline character for empty lines
			_, err := filev.WriteString("\n")
			if err != nil {
				fmt.Println("Error: failed to write to output file:", err)
				os.Exit(1)
			}
		}
	}
}

// SeparateNewLine takes a string and splits it into words, recognizing the "\\n" marker
// as a newline character.
func SeparateNewLine(str string) []string {
	var words []string
	word := ""
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 {
			// Check for "\\n" marker and split the words
			if str[i] == '\\' && str[i+1] == 'n' {
				words = append(words, word)
				words = append(words, "N3w L1n3")
				word = ""
				i++
			} else {
				word += string(str[i])
			}

		} else {
			word += string(str[i])
			words = append(words, word)
		}
	}
	return words
}

