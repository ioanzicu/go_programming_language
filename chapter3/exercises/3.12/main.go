// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.
package main

import "fmt"

func main() {
	fmt.Println(isAnagram("mama", "amam"))     // true
	fmt.Println(isAnagram("banana", "nanaba")) // true
	fmt.Println(isAnagram("mama", "mana"))     // false
	fmt.Println(isAnagram("bird", "birt"))     // false
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Map := map[rune]int{}
	for _, char := range s1 {
		s1Map[char]++
	}

	s2Map := map[rune]int{}
	for _, char := range s2 {
		s2Map[char]++
	}

	for _, char := range s1 {
		count1, ok := s1Map[char]
		if !ok {
			return false
		}

		count2, ok := s2Map[char]
		if !ok {
			return false
		}

		if count1 != count2 {
			return false
		}
	}
	return true
}
