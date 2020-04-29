package dom

import "strings"

func ConvertToParagraphElements(str string) string {
	segments := strings.Split(str, "\n")

	var builder strings.Builder

	for _, seg := range segments {
		builder.WriteString("<p>")
		builder.WriteString(strings.TrimSpace(seg))
		builder.WriteString("</p>")
	}

	return builder.String()
}
