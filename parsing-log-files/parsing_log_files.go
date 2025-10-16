package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\s*\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var total int
	re := regexp.MustCompile(`"[^"]*(?i)password[^"]*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			total++
		}
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
    
    result := make([]string, len(lines))
    for i, line := range lines {
        matches := re.FindStringSubmatch(line)
        if len(matches) > 1 {
            username := matches[1]
            result[i] = fmt.Sprintf("[USR] %s %s", username, line)
        } else {
            result[i] = line
        }
    }
    
    return result
}
