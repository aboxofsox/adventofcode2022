package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func priorityMap() map[string]int {
	mp := make(map[string]int)
	p := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ps := strings.Split(p, "")

	for i, s := range ps {
		mp[s] = i + 1
	}

	return mp
}

func sum(i []int) int {
	sm := 0
	for _, ii := range i {
		sm += ii
	}
	return sm
}

func splitter(s string) (string, string) {
	ss := strings.Split(s, "")

	first := ss[:len(s)/2]
	second := ss[len(s)/2:]

	return strings.Join(first, ""), strings.Join(second, "")
}

func findCommon(first, second string) int {
	mp := priorityMap()
	fs := strings.Split(first, "")
	ss := strings.Split(second, "")

	sort.Strings(fs)
	sort.Strings(ss)

	for _, f := range fs {
		for _, s := range ss {
			if f == s {
				return mp[f]
			}
		}
	}

	return 0
}

func findCommonGroup(ss []string) int {
	mp := priorityMap()
	first := strings.Split(ss[0], "")
	second := strings.Split(ss[1], "")
	third := strings.Split(ss[2], "")

	sort.Strings(first)
	sort.Strings(second)
	sort.Strings(third)

	// it's not about efficiency. it's about getting the right answer...🤮🤮🤮
	for _, f := range first {
		for _, s := range second {
			for _, t := range third {
				if f == s && f == t {
					return mp[f]
				}
			}
		}
	}
	return 0
}

func threePart(s []string, z int) [][]string {
	var ss [][]string
	var j int
	for i := 0; i < len(s); i += z {
		j += z
		if j > len(s) {
			j = len(s)
		}
		ss = append(ss, s[i:j])
	}

	return ss
}

func main() {
	var lines []string
	var sums []int
	var gsums []int
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		first, second := splitter(scanner.Text())
		n := findCommon(first, second)

		sums = append(sums, n)
	}

	// part 1
	smm := sum(sums)

	fmt.Println(smm)

	// part 2
	tp := threePart(lines, 3)

	for _, t := range tp {
		n := findCommonGroup(t)
		gsums = append(gsums, n)
	}

	gs := sum(gsums)

	fmt.Println(gs)

}
