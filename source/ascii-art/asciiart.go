package asciiart

import (
	"bufio"
	"os"
	"strings"
)

// function to load files
func GetBanner(input string, num int) (string, error) {
	file, err := os.Open(input)
	if err != nil {
		return "", err
	}
	defer file.Close()
	//create a new scanner to read content of the file
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

// fuction to generate the asciiart
func GenerateArt(input, filename string) (string, error) {
	banner := "banners/" + filename + ".txt"
	var result strings.Builder
	arg := strings.Split(input, "\n")
	for _, word := range arg {
		for i := range 8 {
			for _, runes := range word {
				//called the getbanner function help generate the art.
				asciiline, err := GetBanner(banner, 1+int(runes-' ')*9+i)
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
