/*
Date: 2023-10-29
ProblemID: 1410
ProblemName: HTML 实体解析器
*/
package leetcode

import (
	"strings"
)

func entityParser(text string) string {
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&apos;", "'")

	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&frasl;", "/")

	text = strings.ReplaceAll(text, "&amp;", "&")
	return text
}
