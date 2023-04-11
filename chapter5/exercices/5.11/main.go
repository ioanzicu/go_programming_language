// The instructor of the linear algebra course decides that calculus is now a
// prerequisite. Extend the topoSort function to report cycles.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"}, // cycle
}

func main() {
	sortReqs, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
	}

	for i, course := range sortReqs {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for key := range m {
		path, circleFound := findCycle(key, nil, m)
		if circleFound {
			return nil, fmt.Errorf("Circle found: %s", strings.Join(path, " => "))
		}
		visitAll([]string{key})
	}
	// sort.Strings(keys)
	// visitAll(keys)
	return order, nil
}

func findCycle(key string, path []string, m map[string][]string) ([]string, bool) {
	if path == nil {
		path = append(path, key)
	}

	for _, item := range m[key] {
		for i := 0; i < len(path); i++ {
			if path[i] == item {
				path = append(path, item+" @#"+strconv.Itoa(i))
				return path, true
			}
		}
		path = append(path, item)
		return findCycle(item, path, m)
	}
	return nil, false
}
