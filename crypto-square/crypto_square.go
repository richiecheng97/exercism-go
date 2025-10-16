package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	input := normalize(pt)
	if len(input) == 0 {
		return ""
	}
	r, c := dimensions(len(input))
	text := input
	for len(text) < r*c {
		text += " "
	}
	res := readColumns(text, r, c)
	out := formatOutput(res, r, c)
	return out
}

func normalize(pt string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := re.ReplaceAllString(pt, "")
	return strings.ToLower(cleaned)
}

func dimensions(length int) (int, int) {
	c := int(math.Ceil(math.Sqrt(float64(length))))
	r := int(math.Ceil(float64(length) / float64(c)))

	// for c < r || c-r > 1 {
	// 	c++
	// 	r = int(math.Ceil(float64(length) / float64(c)))
	// }
	return r, c
}

func createRectangle(text string, r, c int) []string {
	rectangle := make([]string, r)
	for i := 0; i < r; i++ {
		start := i * c
		end := start + c

		if start >= len(text) {
			rectangle[i] = strings.Repeat(" ", c)
		} else if end > len(text) {
			row := text[start:]
			rectangle[i] = row + strings.Repeat(" ", c-len(row))
		} else {
			rectangle[i] = text[start:end]
		}
	}
	return rectangle
}

func readColumns(text string, r, c int) [][]rune {
	grid := make([][]rune, r)
	for i := range grid {
		grid[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			idx := i*c + j
			if idx < len(text) {
				grid[i][j] = rune(text[idx])
			} else {
				grid[i][j] = ' '
			}
		}
	}

	return grid
}

func formatOutput(grid [][]rune, r, c int) string {
	var result strings.Builder
	for j := 0; j < c; j++ {
		if j > 0 {
			result.WriteRune(' ')
		}
		for i := 0; i < r; i++ {
			result.WriteRune(grid[i][j])
		}
	}

	return result.String()
}
