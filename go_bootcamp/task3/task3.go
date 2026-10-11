package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseToIntSlice(line string) ([]int, error) {
	fields := strings.Fields(line)
	var result []int

	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	line1 := scanner.Text()

	list1, err := parseToIntSlice(line1)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if !scanner.Scan() {
		return
	}
	line2 := scanner.Text()

	list2, err := parseToIntSlice(line2)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	set2 := make(map[int]struct{})
	for _, num := range list2 {
		set2[num] = struct{}{}
	}

	seen := make(map[int]struct{})
	var intersection []int

	for _, num := range list1 {
		if _, existsInSecond := set2[num]; existsInSecond {
			if _, alreadyAdded := seen[num]; !alreadyAdded {
				intersection = append(intersection, num)
				seen[num] = struct{}{}
			}
		}
	}

	var resultStrings []string
	for _, num := range intersection {
		resultStrings = append(resultStrings, strconv.Itoa(num))
	}

	if len(resultStrings) == 0 {
		fmt.Println("Empty intersection")
	} else {
		fmt.Printf("%s\n", strings.Join(resultStrings, " "))
	}
}
