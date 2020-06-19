package strmod

import (
	"strings"
	"unicode"
)

// Title is executing stdlib strings.Title().
func Title() Modifier {
	return func(s string) string {
		return strings.Title(s)
	}
}

// ToLower is executing stdlib strings.ToLower().
func ToLower() Modifier {
	return func(s string) string {
		return strings.ToLower(s)
	}
}

// ToLowerSpecial is executing stdlib strings.ToLowerSpecial().
func ToLowerSpecial(c unicode.SpecialCase) Modifier {
	return func(s string) string {
		return strings.ToLowerSpecial(c, s)
	}
}

// ToTitle is executing stdlib strings.ToTitle().
func ToTitle() Modifier {
	return func(s string) string {
		return strings.ToTitle(s)
	}
}

// ToTitleSpecial is executing stdlib strings.ToTitleSpecial().
func ToTitleSpecial(c unicode.SpecialCase) Modifier {
	return func(s string) string {
		return strings.ToTitleSpecial(c, s)
	}
}

// ToUpper is executing stdlib strings.ToUpper().
func ToUpper() Modifier {
	return func(s string) string {
		return strings.ToUpper(s)
	}
}

// ToUpperSpecial is executing stdlib strings.ToUpperSpecial().
func ToUpperSpecial(c unicode.SpecialCase) Modifier {
	return func(s string) string {
		return strings.ToUpperSpecial(c, s)
	}
}

// ToValidUTF8 is executing stdlib strings.ToValidUTF8().
func ToValidUTF8(replacement string) Modifier {
	return func(s string) string {
		return strings.ToValidUTF8(s, replacement)
	}
}

// Trim is executing stdlib strings.Trim().
func Trim(cutset string) Modifier {
	return func(s string) string {
		return strings.Trim(s, cutset)
	}
}

// TrimFunc is executing stdlib strings.TrimFunc().
func TrimFunc(f func(rune) bool) Modifier {
	return func(s string) string {
		return strings.TrimFunc(s, f)
	}
}

// TrimLeft is executing stdlib strings.TrimLeft().
func TrimLeft(cutset string) Modifier {
	return func(s string) string {
		return strings.TrimLeft(s, cutset)
	}
}

// TrimLeftFunc is executing stdlib strings.TrimLeftFunc().
func TrimLeftFunc(f func(rune) bool) Modifier {
	return func(s string) string {
		return strings.TrimLeftFunc(s, f)
	}
}

// TrimPrefix is executing stdlib strings.TrimPrefix().
func TrimPrefix(prefix string) Modifier {
	return func(s string) string {
		return strings.TrimPrefix(s, prefix)
	}
}

// TrimRight is executing stdlib strings.TrimRight().
func TrimRight(cutset string) Modifier {
	return func(s string) string {
		return strings.TrimRight(s, cutset)
	}
}

// TrimRightFunc is executing stdlib strings.TrimRightFunc().
func TrimRightFunc(f func(rune) bool) Modifier {
	return func(s string) string {
		return strings.TrimRightFunc(s, f)
	}
}

// TrimSpace is executing stdlib strings.TrimSpace().
func TrimSpace() Modifier {
	return func(s string) string {
		return strings.TrimSpace(s)
	}
}

// TrimSuffix is executing stdlib strings.TrimSuffix().
func TrimSuffix(suffix string) Modifier {
	return func(s string) string {
		return strings.TrimSuffix(s, suffix)
	}
}

// Replace is executing stdlib strings.Replace().
func Replace(old, new string, n int) Modifier {
	return func(in string) string {
		return strings.Replace(in, old, new, n)
	}
}

// ReplaceAll is executing stdlib strings.ReplaceAll().
func ReplaceAll(old, new string) Modifier {
	return func(in string) string {
		return strings.ReplaceAll(in, old, new)
	}
}
