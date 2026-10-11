package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Visit struct {
	Specialty string
	Date      string
}

type PatientNotFoundError struct {
	PatientName string
}

func (e *PatientNotFoundError) Error() string {
	return "patient not found"
}

var medicalRecords = make(map[string][]Visit)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}

		cmd := strings.TrimSpace(scanner.Text())
		if cmd == "" {
			continue
		}

		switch cmd {
		case "Save":
			handleSave(scanner)
		case "GetHistory":
			handleGetHistory(scanner)
		case "GetLastVisit":
			handleGetLastVisit(scanner)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения ввода: %v\n", err)
			break
		}
	}
}

func handleSave(scanner *bufio.Scanner) {
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	scanner.Scan()
	specialty := strings.TrimSpace(scanner.Text())

	scanner.Scan()
	date := strings.TrimSpace(scanner.Text())

	medicalRecords[name] = append(medicalRecords[name], Visit{
		Specialty: specialty,
		Date:      date,
	})
}

func handleGetHistory(scanner *bufio.Scanner) {
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	visits, exists := medicalRecords[name]
	if !exists || len(visits) == 0 {
		fmt.Println(&PatientNotFoundError{PatientName: name})
		return
	}

	for _, v := range visits {
		fmt.Printf("%s %s\n", v.Specialty, v.Date)
	}
}

func handleGetLastVisit(scanner *bufio.Scanner) {
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	scanner.Scan()
	specialty := strings.TrimSpace(scanner.Text())

	visits, exists := medicalRecords[name]
	if !exists {
		fmt.Println(&PatientNotFoundError{PatientName: name})
		return
	}

	var lastDate string
	found := false

	for _, v := range visits {
		if v.Specialty == specialty {
			if !found || v.Date > lastDate {
				lastDate = v.Date
				found = true
			}
		}
	}

	if !found {
		return
	}

	fmt.Println(lastDate)
}
