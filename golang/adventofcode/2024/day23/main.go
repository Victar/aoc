package main

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"sort"
	"strings"
)

var DAY = "23"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	graph := make(map[string]map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if graph[parts[0]] == nil {
			graph[parts[0]] = make(map[string]bool)
		}
		if graph[parts[1]] == nil {
			graph[parts[1]] = make(map[string]bool)
		}
		graph[parts[0]][parts[1]] = true
		graph[parts[1]][parts[0]] = true
	}
	count := 0
	visited := make(map[string]bool)
	for c1 := range graph {
		for c2 := range graph[c1] {
			for c3 := range graph[c2] {
				if c3 != c1 && graph[c3][c1] && (strings.HasPrefix(c1, "t") || strings.HasPrefix(c2, "t") || strings.HasPrefix(c3, "t")) {
					strs := []string{c1, c2, c3}
					slices.Sort(strs)
					key := strings.Join(strs, ",")
					if !visited[key] {
						count++
						visited[key] = true
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	graph := make(map[string]map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if graph[parts[0]] == nil {
			graph[parts[0]] = make(map[string]bool)
		}
		if graph[parts[1]] == nil {
			graph[parts[1]] = make(map[string]bool)
		}
		graph[parts[0]][parts[1]] = true
		graph[parts[1]][parts[0]] = true
	}

	cliques := [][]string{}
	//https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
	R, P, X := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	for node := range graph {
		P[node] = true
	}
	bronKerbosch(R, P, X, graph, &cliques)

	maxClique := []string{}
	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}
	sort.Strings(maxClique)
	fmt.Println(strings.Join(maxClique, ","))

}

// That is, in pseudocode, the algorithm performs the following steps:
// algorithm BronKerbosch1(R, P, X) is
// if P and X are both empty then
// report R as a maximal clique
// for each vertex v in P do
// BronKerbosch1(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
// P := P \ {v}
// X := X ⋃ {v}
func bronKerbosch(R, P, X map[string]bool, graph map[string]map[string]bool, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		clique := slice(R)
		*cliques = append(*cliques, clique)
		return
	}
	for v := range P {
		RNew := union(R, set(v))
		PNew := intersection(P, graph[v])
		XNew := intersection(X, graph[v])
		bronKerbosch(RNew, PNew, XNew, graph, cliques)
		P = difference(P, set(v))
		X = union(X, set(v))
	}
}

func union(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range a {
		res[k] = true
	}
	for k := range b {
		res[k] = true
	}
	return res
}

func intersection(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range a {
		if b[k] {
			res[k] = true
		}
	}
	return res
}

func difference(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range a {
		if !b[k] {
			res[k] = true
		}
	}
	return res
}
func set(a string) map[string]bool {
	return map[string]bool{a: true}
}

func slice(m map[string]bool) []string {
	s := make([]string, 0, len(m))
	for node := range m {
		s = append(s, node)
	}
	return s
}
