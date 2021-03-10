package window

import (
	"fmt"
	"html"
	"strconv"
)

func Sprintf(format string, a ...interface{}) string {
	safeArgs := make([]interface{}, len(a))
	for i, arg := range a {
		safeArgs[i] = safe{arg}
	}
	return fmt.Sprintf(format, safeArgs...)
}

type safe struct {
	wrapped interface{}
}

func (safe safe) Format(f fmt.State, verb rune) {
	_, _ = f.Write([]byte(html.EscapeString(fmt.Sprintf(recreateFormatVerb(f, verb), safe.wrapped))))
}

func recreateFormatVerb(f fmt.State, verb rune) string {
	var array [50]rune
	verbWithFlags := append(array[:0], '%')

	for _, flag := range "+-# 0" {
		if f.Flag(int(flag)) {
			verbWithFlags = append(verbWithFlags, flag)
		}
	}

	if w, isSet := f.Width(); isSet {
		verbWithFlags = append(verbWithFlags, []rune(strconv.Itoa(w))...)
	}
	if w, isSet := f.Precision(); isSet {
		verbWithFlags = append(verbWithFlags, '.')
		verbWithFlags = append(verbWithFlags, []rune(strconv.Itoa(w))...)
	}
	verbWithFlags = append(verbWithFlags, verb)

	return string(verbWithFlags)
}
