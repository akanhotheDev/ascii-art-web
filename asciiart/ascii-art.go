package asciiart

import (
	// "fmt"
	"os"
	"strings"
)

func getChar(c rune, line []string) []string {
	startLine := (int(c)-32)*9 + 1
	if startLine < 0 || startLine+8 > len(line) {
		return make([]string, 8)
	}
	return line[startLine : startLine+8]
}

// func renderAscii () string{
// 	if len(os.Args) < 2 {
// 		fmt.Println("Pass In your Text.")
// 	}
// 	textArg := os.Args[1: len(os.Args) -1]
// 	fileName := "banner/" + "standard" + ".txt"

// 	if len(os.Args) > 2 {
// 		fileName = "banner/" + os.Args[2] + ".txt"
// 	}
// 	data, err := os.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println("Error reading Font-file", err)

// 	}
// 	content := strings.Split(string(data), "\n")
// 	input := strings.Join(textArg, " ")
// 	input = strings.ReplaceAll(input,`\n`, "\n")

// 	return (Render(input, content))
// }
func Render(input string, fontFile string) string {

	data, err := os.ReadFile(fontFile)
	if err != nil {
		return "error loading font"
	}

	raw := strings.Split(string(data), "\n")

	var lines [] string

	for _, line := range raw {
		lines = append(lines, strings.TrimRight(line,"\r"))
	}

	input = strings.ReplaceAll(input, `\n`, "\n")
	parts := strings.Split(input, "\n")

	var builder strings.Builder

	for i, part := range parts {

		if part == "" {
			if i < len(parts)-1 {
				builder.WriteString("\n")
			}
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range part {
				charLine := getChar(char, lines)
				builder.WriteString(charLine[row])
			}
			builder.WriteString("\n")
		}
	}

	return builder.String()
}