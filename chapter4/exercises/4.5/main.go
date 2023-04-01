// Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func main() {
	// s := []string{"one", "one", "two", "three"}
	s := []string{"one", "one", "two", "two", "two", "three"}
	fmt.Println(s)
	t := removeAdjDuplicatest(s)
	fmt.Println(t)

	fmt.Println("")

	s2 := []string{"one", "one", "two", "two", "two", "three"}
	fmt.Println(s2)
	t2 := removeAdjDuplicatest2(s2)
	fmt.Println(t2)
}

func removeAdjDuplicatest(strings []string) []string {
	k := 0
	for i := 0; i < len(strings); {
		j := i
		// move to the index that is not adjiacent duplicate
		for ; j < len(strings) && strings[i] == strings[j]; j++ {
		}

		// assign non duplicated values to the k position
		strings[k] = strings[i]
		k++
		i = j // update i to the index where are non duplicated item starts
	}
	return strings[:k]
}

func removeAdjDuplicatest2(strings []string) []string {
	out := strings[:0]
	for i := 0; i < len(strings); {
		j := i
		// move to the index that is not adjiacent duplicate
		for ; j < len(strings) && strings[i] == strings[j]; j++ {
		}

		// append non duplicated values
		out = append(out, strings[i])
		i = j // update i to the index where are non duplicated item starts
	}
	return out
}
