package stringsx

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Hello", "olleH"}, {"a", "a"}, {"", ""}, {"héllo 世界", "界世 olléh"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true}, {"racecar", true}, {"Go", false},
		{"No 'x' in Nixon", true}, {"hello", false}, {"12321", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %t, want %t", tt.in, got, tt.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"  the quick   brown fox ", 4}, {"one", 1}, {"", 0}, {"   ", 0}, {"tab\tand\nnewline", 3},
	}
	for _, tt := range tests {
		if got := CountWords(tt.in); got != tt.want {
			t.Errorf("CountWords(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct{ in, want string }{
		{"hello wonderful world", "Hello Wonderful World"}, {"go", "Go"},
		{"already Capital", "Already Capital"}, {"", ""},
	}
	for _, tt := range tests {
		if got := Capitalize(tt.in); got != tt.want {
			t.Errorf("Capitalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestByteAndRuneCount(t *testing.T) {
	tests := []struct {
		in           string
		bytes, runes int
	}{
		{"hello", 5, 5}, {"héllo", 6, 5}, {"世界", 6, 2}, {"", 0, 0},
	}
	for _, tt := range tests {
		b, r := ByteAndRuneCount(tt.in)
		if b != tt.bytes || r != tt.runes {
			t.Errorf("ByteAndRuneCount(%q) = %d, %d, want %d, %d", tt.in, b, r, tt.bytes, tt.runes)
		}
	}
}

func TestAcronym(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Portable Network Graphics", "PNG"}, {"as soon as possible", "ASAP"}, {"Go", "G"},
	}
	for _, tt := range tests {
		if got := Acronym(tt.in); got != tt.want {
			t.Errorf("Acronym(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestParseKeyValue(t *testing.T) {
	tests := []struct {
		in         string
		key, value string
		ok         bool
	}{
		{"name=John", "name", "John", true},
		{" host = a=b ", "host", "a=b", true},
		{"empty=", "empty", "", true},
		{"nothing", "", "", false},
	}
	for _, tt := range tests {
		k, v, ok := ParseKeyValue(tt.in)
		if k != tt.key || v != tt.value || ok != tt.ok {
			t.Errorf("ParseKeyValue(%q) = %q, %q, %t, want %q, %q, %t", tt.in, k, v, ok, tt.key, tt.value, tt.ok)
		}
	}
}

func TestSumCSV(t *testing.T) {
	tests := []struct {
		in      string
		want    int
		wantErr bool
	}{
		{"1,2,3", 6, false}, {" 4 , 5 ", 9, false}, {"-1,1", 0, false}, {"", 0, false},
		{"1,x", 0, true}, {"1,,2", 0, true},
	}
	for _, tt := range tests {
		got, err := SumCSV(tt.in)
		if tt.wantErr {
			if err == nil {
				t.Errorf("SumCSV(%q) returned no error, want an error", tt.in)
			}
			continue
		}
		if err != nil || got != tt.want {
			t.Errorf("SumCSV(%q) = %d, %v, want %d, nil", tt.in, got, err, tt.want)
		}
	}
}

func TestCompress(t *testing.T) {
	tests := []struct{ in, want string }{
		{"aaabcc", "a3b1c2"}, {"a", "a1"}, {"", ""}, {"zzzzzzzzzzzz", "z12"}, {"abab", "a1b1a1b1"},
	}
	for _, tt := range tests {
		if got := Compress(tt.in); got != tt.want {
			t.Errorf("Compress(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestMaskCard(t *testing.T) {
	tests := []struct{ in, want string }{
		{"1234567812345678", "************5678"}, {"12345", "*2345"}, {"1234", "1234"}, {"12", "12"},
	}
	for _, tt := range tests {
		if got := MaskCard(tt.in); got != tt.want {
			t.Errorf("MaskCard(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct{ in, want string }{
		{"helloWorld", "hello_world"}, {"HelloWorld", "hello_world"},
		{"userID", "user_i_d"}, {"simple", "simple"},
	}
	for _, tt := range tests {
		if got := CamelToSnake(tt.in); got != tt.want {
			t.Errorf("CamelToSnake(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCaesarCipher(t *testing.T) {
	tests := []struct {
		in    string
		shift int
		want  string
	}{
		{"abc", 1, "bcd"}, {"xyz", 3, "abc"}, {"Hello, World!", 13, "Uryyb, Jbeyq!"},
		{"bcd", -1, "abc"}, {"abc", 27, "bcd"}, {"ABC", -27, "ZAB"},
	}
	for _, tt := range tests {
		if got := CaesarCipher(tt.in, tt.shift); got != tt.want {
			t.Errorf("CaesarCipher(%q, %d) = %q, want %q", tt.in, tt.shift, got, tt.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		in    string
		limit int
		want  string
	}{
		{"Hello, World", 5, "Hello..."}, {"Hi", 5, "Hi"}, {"exact", 5, "exact"}, {"héllo wörld", 4, "héll..."},
	}
	for _, tt := range tests {
		if got := Truncate(tt.in, tt.limit); got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.in, tt.limit, got, tt.want)
		}
	}
}

func TestCountSubstring(t *testing.T) {
	tests := []struct {
		s, sub string
		want   int
	}{
		{"Go go GO gopher", "go", 4}, {"banana", "ana", 1}, {"abc", "x", 0},
	}
	for _, tt := range tests {
		if got := CountSubstring(tt.s, tt.sub); got != tt.want {
			t.Errorf("CountSubstring(%q, %q) = %d, want %d", tt.s, tt.sub, got, tt.want)
		}
	}
}
