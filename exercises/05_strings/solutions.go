//go:build solutions

// SPOILER ALERT: reference solutions. Try the exercises first!

package stringsx

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	var clean []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(clean)-1; i < j; i, j = i+1, j-1 {
		if clean[i] != clean[j] {
			return false
		}
	}
	return true
}

func CountWords(s string) int {
	return len(strings.Fields(s))
}

func Capitalize(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		if w == "" {
			continue
		}
		runes := []rune(w)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

func ByteAndRuneCount(s string) (bytes, runes int) {
	return len(s), utf8.RuneCountInString(s)
}

func Acronym(s string) string {
	var b strings.Builder
	for _, w := range strings.Fields(s) {
		first, _ := utf8.DecodeRuneInString(w)
		b.WriteRune(unicode.ToUpper(first))
	}
	return b.String()
}

func ParseKeyValue(s string) (key, value string, ok bool) {
	key, value, ok = strings.Cut(s, "=")
	if !ok {
		return "", "", false
	}
	return strings.TrimSpace(key), strings.TrimSpace(value), true
}

func SumCSV(s string) (int, error) {
	if strings.TrimSpace(s) == "" {
		return 0, nil
	}
	sum := 0
	for _, part := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}

func Compress(s string) string {
	runes := []rune(s)
	var b strings.Builder
	for i := 0; i < len(runes); {
		j := i
		for j < len(runes) && runes[j] == runes[i] {
			j++
		}
		b.WriteRune(runes[i])
		b.WriteString(strconv.Itoa(j - i))
		i = j
	}
	return b.String()
}

func MaskCard(card string) string {
	if len(card) <= 4 {
		return card
	}
	return strings.Repeat("*", len(card)-4) + card[len(card)-4:]
}

func CamelToSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			r = unicode.ToLower(r)
		}
		b.WriteRune(r)
	}
	return b.String()
}

func CaesarCipher(s string, shift int) string {
	shift = (shift%26 + 26) % 26
	runes := []rune(s)
	for i, r := range runes {
		switch {
		case r >= 'a' && r <= 'z':
			runes[i] = 'a' + (r-'a'+rune(shift))%26
		case r >= 'A' && r <= 'Z':
			runes[i] = 'A' + (r-'A'+rune(shift))%26
		}
	}
	return string(runes)
}

func Truncate(s string, limit int) string {
	runes := []rune(s)
	if len(runes) <= limit {
		return s
	}
	return string(runes[:limit]) + "..."
}

func CountSubstring(s, sub string) int {
	return strings.Count(strings.ToLower(s), strings.ToLower(sub))
}
