package ascii

import (
	"bufio"
	"os"
	"strings"
)

func GetBanner(input string, num int) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := 0
	text := ""
	for scanner.Scan() {
		if linenum == num {
			text = scanner.Text()
		}
		linenum++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return text, nil
}
func GenerateArt(input, filename string) (string, error) {
	banner := "banners/" + filename + ".txt"
	var result strings.Builder
	arg := strings.Split(input, "\n")
	for _, word := range arg {
		for i := range 8 {
			for _, rune := range word {
				start := (int(rune)-32)*9 + 1
				asciiline, err := GetBanner(banner, start+i)
				if err != nil {
					return "", err
				}
				result.WriteString(asciiline)
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
