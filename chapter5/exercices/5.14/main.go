// Use the breadthFirst function to explore a different structure. For example,
// you could use the course dependencies from the topoSort example (a directed graph), the file
// system hierarchy on your computer (a tree), or a list of bus or subway routes downloaded from
// your city government’s web site (an undirected graph).
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var stdout io.Writer = os.Stdout // modified during testing

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
		"Assebmler 101",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization", "Golang"},
}

func main() {
	fmt.Print("--- DFS\n\n")
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	fmt.Print("\n\n--- BFS\n")
	for i, course := range breadthFirst(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// dfs
func topoSort(m map[string][]string) []string {
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

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// bfs
func breadthFirst(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var worklist []string
	for key := range m {
		worklist = append(worklist, key)
	}
	sort.Strings(worklist)

	for len(worklist) > 0 {
		// substitute the queue - we append at the end, and consume iterator from begining
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				for _, neighbor := range m[item] {
					worklist = append(worklist, neighbor)
				}
				order = append(order, item)
			}
		}
	}

	return order
}
