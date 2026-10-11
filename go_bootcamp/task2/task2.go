package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func TopKFrequentWords(text string, k int) []string {
	if k <= 0 {
		return []string{}
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{}
	}

	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}

	type pair struct {
		word  string
		count int
	}
	pairs := make([]pair, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, pair{w, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].word < pairs[j].word
		}
		return pairs[i].count > pairs[j].count
	})

	if k > len(pairs) {
		k = len(pairs)
	}

	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i].word
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		fmt.Println()
		return
	}
	text := scanner.Text()

	if !scanner.Scan() {
		fmt.Println()
		return
	}
	kStr := strings.TrimSpace(scanner.Text())
	k, err := strconv.Atoi(kStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: K must be an integer\n")
		os.Exit(1)
	}

	result := TopKFrequentWords(text, k)
	fmt.Println(strings.Join(result, " "))
}
