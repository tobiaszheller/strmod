package strmod

import "strings"

func Modify(in string, mods ...Modifier) string {
	s := in
	for _, mod := range mods {
		s = mod(s)
	}
	return s
}

type Modifier func(string) string

func TrimSpace() Modifier {
	return func(in string) string {
		return strings.TrimSpace(in)
	}
}

func ToLower() Modifier {
	return func(in string) string {
		return strings.ToLower(in)
	}
}

func TrimPrefix(prefix string) Modifier {
	return func(in string) string {
		return strings.TrimPrefix(in, prefix)
	}
}

func TrimSuffix(suffix string) Modifier {
	return func(in string) string {
		return strings.TrimSuffix(in, suffix)
	}
}

func Replace(old, new string, n int) Modifier {
	return func(in string) string {
		return strings.Replace(in, old, new, n)
	}
}
func ReplaceAll(old, new string) Modifier {
	return func(in string) string {
		return strings.ReplaceAll(in, old, new)
	}
}

func SplitAndReturnLast(sep string) Modifier {
	return func(in string) string {
		if sep == "" && in == "" {
			return ""
		}
		splitted := strings.Split(in, sep)
		return splitted[len(splitted)-1]
	}
}

func SplitAndReturnFirst(sep string) Modifier {
	return func(in string) string {
		if sep == "" && in == "" {
			return ""
		}
		splitted := strings.Split(in, sep)
		return splitted[0]
	}
}
