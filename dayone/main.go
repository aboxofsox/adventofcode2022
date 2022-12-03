package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("%s\n", string(b))
}

func splitter(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n\n")
}

func splitLine(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}

func sum(ns []string) int {
	var s int
	for _, n := range ns {
		if n != "" {
			s += toInt(n)
		}
	}
	return s
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}

func grouper(ss []string) map[int][]string {
	mp := make(map[int][]string, len(ss))
	for i, s := range ss {
		mp[i] = splitLine(s)
	}
	return mp
}

func sortMap(mp map[int]int) []int {
	sorted := make(map[int]int, len(mp))
	keys := make([]int, 0, len(mp))

	for k := range mp {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return mp[keys[i]] < mp[keys[j]]
	})

	for i := range keys {
		sorted[keys[i]] = mp[keys[i]]
	}

	return keys
}

func main() {
	var top3 int
	input := readInput()
	sp := splitter(input)
	gp := grouper(sp)
	mp := make(map[int]int, len(sp))

	for k, v := range gp {
		mp[k] = sum(v)
	}

	sorted := sortMap(mp)

	for i := range sorted {
		fmt.Printf("{ %d: %d }\n", sorted[i], mp[sorted[i]])

	}

	top3 += mp[sorted[len(sorted)-1]] + mp[sorted[len(sorted)-2]] + mp[sorted[len(sorted)-3]]

	fmt.Println(top3)
}
