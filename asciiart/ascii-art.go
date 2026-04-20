package asciiart

import (
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