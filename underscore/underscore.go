package underscore

import "strings"

func Camel(inp string) string {
	var sb strings.Builder
	for _, ch := range inp {
		if ch >= 'A' && ch <= 'Z' {
			sb.WriteRune('_')
		}
		sb.WriteRune(ch)
	}
	return strings.ToLower(sb.String())
}
