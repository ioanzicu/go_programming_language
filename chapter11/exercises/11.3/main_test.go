/*
Modify randomPalindrome to exercise IsPalindrome’s
handling of punctuation and spaces.
*/

package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	m := rng.Intn(n + 1)
	var temp []rune
	temp = append(temp, runes[:m]...)
	temp = append(temp, ' ') // Insert space
	temp = append(temp, runes[m:]...)

	m = rng.Intn(n + 1)
	runes = append(make([]rune, 0), temp[:m]...)
	runes = append(runes, ',') // Insert punctuation
	runes = append(runes, temp[m:]...)
	return string(runes)
}
func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2 // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < n-1; i++ {
		r := rune('A' + rng.Intn('Z'-'A'))
		runes[i] = r
	}
	runes[len(runes)-1] = rune(runes[0] + 1)
	return string(runes)
}

func TestIsPalindromesByRandomeNoisy(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if IsPalindrome(p) == true {
			t.Errorf("IsPalindrome(%q) returns true", p)
		}
	}
}
