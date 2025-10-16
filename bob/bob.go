// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// 去除首尾空白字符
	trimmedRemark := strings.TrimSpace(remark)

	// 1. 检查是否为沉默（空字符串）
	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	// 2. 检查是否为大喊大叫（全部大写字母且至少包含一个字母）
	isYelling := isYelling(trimmedRemark)

	// 3. 检查是否为问题（以问号结尾）
	isQuestion := isQuestion(trimmedRemark)

	// 4. 根据条件返回相应回答
	switch {
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

// isYelling 检查是否为大喊大叫
func isYelling(s string) bool {
	// 必须至少包含一个字母
	hasLetter := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetter = true
			break
		}
	}

	if !hasLetter {
		return false
	}

	// 所有字母都必须是大写
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}

	return true
}

// isQuestion 检查是否为问题
func isQuestion(s string) bool {
	// 使用正则表达式检查是否以问号结尾（忽略后面的空白字符）
	pattern := regexp.MustCompile(`\?\s*$`)
	return pattern.MatchString(s)
}
