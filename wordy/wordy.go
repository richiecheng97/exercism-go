package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is ") {
		return 0, false
	}

	question = strings.TrimPrefix(question, "What is ")
	question = strings.TrimSuffix(question, "?")

	if len(question) == 0 {
		return 0, false
	}

	return parseExpression(question)
}

func parseExpression(question string) (int, bool) {
	question = strings.TrimSpace(question)

	first, remaining, b := extractFirstNumber(question)
	if !b {
		return 0, false
	}
	if remaining == "" {
		return first, true
	}

	// 从左到右计算
	result := first
	for remaining != "" {
		var operation string
		var operand int

		// 确定操作符
		if strings.HasPrefix(remaining, " plus ") {
			operation = "plus"
			remaining = strings.TrimPrefix(remaining, " plus ")
		} else if strings.HasPrefix(remaining, " minus ") {
			operation = "minus"
			remaining = strings.TrimPrefix(remaining, " minus ")
		} else if strings.HasPrefix(remaining, " multiplied by ") {
			operation = "multiplied by"
			remaining = strings.TrimPrefix(remaining, " multiplied by ")
		} else if strings.HasPrefix(remaining, " divided by ") {
			operation = "divided by"
			remaining = strings.TrimPrefix(remaining, " divided by ")
		} else {
			return 0, false
		}

		// 提取操作数
		operand, remaining, b = extractFirstNumber(remaining)
		if !b {
			return 0, false
		}

		// 执行操作
		switch operation {
		case "plus":
			result += operand
		case "minus":
			result -= operand
		case "multiplied by":
			result *= operand
		case "divided by":
			if operand == 0 {
				return 0, false
			}
			result /= operand
		}
	}
	return result, true
}

func extractFirstNumber(s string) (int, string, bool) {
	s = strings.TrimSpace(s)

	re := regexp.MustCompile(`^(-?\d+)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) < 1 {
		return 0, "", false
	}

	numStr := matches[0]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, "", false
	}

	remaining := s[len(numStr):]
	return num, remaining, true
}
