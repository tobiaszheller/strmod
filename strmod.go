package strmod

import (
	"strings"
)

func Modify(in string, mods ...Modifier) string {
	s := in
	for _, mod := range mods {
		s = mod(s)
	}
	return s
}

type Modifier func(string) string

func SplitAndReturnLastPart(sep string) Modifier {
	return func(in string) string {
		if sep == "" && in == "" {
			return ""
		}
		splitted := strings.Split(in, sep)
		return splitted[len(splitted)-1]
	}
}

func SplitAndReturnFirstPart(sep string) Modifier {
	return func(in string) string {
		if sep == "" && in == "" {
			return ""
		}
		splitted := strings.Split(in, sep)
		return splitted[0]
	}
}
