package spartaapproves

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q got %q", want, got)
	}
}
func assertIntegers(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %d got %d", want, got)
	}
}

var cases = []struct {
	// description string
	Arabic int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{14, "XIV"},
	{15, "XV"},
	{16, "XVI"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestD2R(t *testing.T) {
	for _, test := range cases {
		t.Run("Test", func(t *testing.T) {
			got := A2R(test.Arabic)
			want := test.Roman
			assertStrings(t, got, want)
		})
	}
	for _, test := range cases {
		t.Run("Test", func(t *testing.T) {
			got := R2A(test.Roman)
			want := test.Arabic
			assertIntegers(t, got, want)
		})
	}
}
