//go:build !solutions

// Package stringsx covers strings, bytes, runes and the strings/strconv/unicode
// packages. (It's not called "strings" so it doesn't clash with the standard
// library package you'll be importing.)
//
// Quick syntax reminder:
//
//	s := "héllo"
//	len(s)                       // number of BYTES (6 here, é takes 2 bytes)
//	s[0]                         // a byte (uint8)
//	for i, r := range s { }      // r is a rune (int32, a unicode code point)
//	[]rune(s)                    // convert to runes to index characters safely
//	string(runes)                // convert back
//	strings.ToUpper, strings.Split, strings.Fields, strings.Contains,
//	strings.Repeat, strings.Cut, strings.TrimSpace, strings.Join ...
//	strconv.Itoa(42), strconv.Atoi("42")
//	var b strings.Builder; b.WriteString("x"); b.String()
//
// Strings are immutable: you build new strings, you never modify one in place.
package stringsx

// Exercise 1: Reverse
//
// Reverse a string. It must work for any unicode text, not only ASCII.
//
//	Reverse("Hello")    → "olleH"
//	Reverse("héllo 世界") → "界世 olléh"
//
// Concepts: []rune conversion.
func Reverse(s string) string {
	// TODO: implement
	return ""
}

// Exercise 2: IsPalindrome
//
// Report whether s reads the same backward as forward, ignoring case and
// ignoring every character that isn't a letter or a digit.
//
//	IsPalindrome("A man, a plan, a canal: Panama") → true
//
// Concepts: unicode.IsLetter, unicode.IsDigit, unicode.ToLower.
func IsPalindrome(s string) bool {
	// TODO: implement
	return false
}

// Exercise 3: CountWords
//
// Count the words in s. Words are separated by any amount of whitespace.
//
//	CountWords("  the quick   brown fox ") → 4
//
// Concepts: strings.Fields.
func CountWords(s string) int {
	// TODO: implement
	return 0
}

// Exercise 4: Capitalize
//
// Uppercase the first letter of every word (words are separated by single
// spaces). The rest of each word stays unchanged.
//
//	Capitalize("hello wonderful world") → "Hello Wonderful World"
func Capitalize(s string) string {
	// TODO: implement
	return ""
}

// Exercise 5: ByteAndRuneCount
//
// Return the number of bytes and the number of runes (characters) in s.
//
//	ByteAndRuneCount("héllo") → 6, 5
//
// Concepts: len(s) vs utf8.RuneCountInString(s) (package unicode/utf8).
func ByteAndRuneCount(s string) (bytes, runes int) {
	// TODO: implement
	return 0, 0
}

// Exercise 6: Acronym
//
// Build an uppercase acronym from the first letter of each word.
//
//	Acronym("Portable Network Graphics") → "PNG"
//	Acronym("as soon as possible")       → "ASAP"
func Acronym(s string) string {
	// TODO: implement
	return ""
}

// Exercise 7: ParseKeyValue
//
// Split a "key=value" string. Surrounding whitespace around key and value is
// trimmed. If there's no "=", return ok = false.
//
//	ParseKeyValue("name=John")      → "name", "John", true
//	ParseKeyValue(" host = a=b ")   → "host", "a=b", true   (split on the first "=")
//	ParseKeyValue("nothing")        → "", "", false
//
// Concepts: strings.Cut, strings.TrimSpace.
func ParseKeyValue(s string) (key, value string, ok bool) {
	// TODO: implement
	return "", "", false
}

// Exercise 8: SumCSV
//
// Sum a comma-separated list of integers.
// Rules:
// 1) Spaces around numbers are allowed.
// 2) An empty string (i.e. "") is valid, and sums up to 0.
// 3) function should Return an error if any item isn't a number. i.e. a letter, an empty string, etc.
//
//	SumCSV("1,2,3")     → 6, nil
//	SumCSV(" 4 , 5 ")   → 9, nil
//	SumCSV("1,x")       → 0, error
//
// Concepts: strings.Split, strconv.Atoi and its error.
func SumCSV(s string) (int, error) {
	// TODO: implement
	return 0, nil
}

// Exercise 9: Compress
//
// Run-length encode a string: each run of the same character becomes the
// character followed by the run length.
//
//	Compress("aaabcc") → "a3b1c2"
//
// Concepts: strings.Builder, strconv.Itoa.
func Compress(s string) string {
	// TODO: implement
	return ""
}

// Exercise 10: MaskCard
//
// Replace every character except the last 4 with '*'. Strings of length 4 or
// less are returned unchanged.
//
//	MaskCard("1234567812345678") → "************5678"
//
// Concepts: slicing strings (s[i:]), strings.Repeat.
func MaskCard(card string) string {
	// TODO: implement
	return ""
}

// Exercise 11: CamelToSnake
//
// Convert camelCase (or PascalCase) to snake_case.
//
//	CamelToSnake("helloWorld")     → "hello_world"
//	CamelToSnake("MyHTTPServer")   → "my_h_t_t_p_server"  (keep it simple: every uppercase letter starts a new word)
//
// Concepts: unicode.IsUpper, unicode.ToLower, building strings rune by rune.
func CamelToSnake(s string) string {
	// TODO: implement
	return ""
}

// Exercise 12: CaesarCipher
//
// Shift every ASCII letter by shift positions in the alphabet, wrapping
// around and preserving case. Other characters are unchanged. shift can be
// negative or larger than 26.
//
//	CaesarCipher("abc", 1)            → "bcd"
//	CaesarCipher("xyz", 3)            → "abc"
//	CaesarCipher("Hello, World!", 13) → "Uryyb, Jbeyq!"
//
// Concepts: runes are numbers: 'a' + 1 == 'b'.
func CaesarCipher(s string, shift int) string {
	// TODO: implement
	return ""
}

// Exercise 13: Truncate
//
// If s has more than limit characters (runes), cut it to limit characters and
// append "...". Otherwise return s unchanged.
//
//	Truncate("Hello, World", 5) → "Hello..."
//	Truncate("Hi", 5)           → "Hi"
func Truncate(s string, limit int) string {
	// TODO: implement
	return ""
}

// Exercise 14: CountSubstring
//
// Count the non-overlapping occurrences of sub in s, case-insensitively.
//
//	CountSubstring("Go go GO gopher", "go") → 4
//
// Concepts: strings.Count, strings.ToLower.
func CountSubstring(s, sub string) int {
	// TODO: implement
	return 0
}
