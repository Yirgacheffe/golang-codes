package main

import (
	"fmt"
	"regexp"
)

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
}

type Schema struct {
	Databases []Database
}

type Database struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []string
}

func Search(m DatabaseMiner) error {
	s, err := m.GetSchema()
	if err != nil {
		return nil
	}

	re := getRegex()

	for _, db := range s.Databases {
		for _, table := range db.Tables {
			for _, field := range table.Columns {
				for _, r := range re {
					if r.MatchString(field) {
						fmt.Println(db)
						fmt.Printf("[+] HIT: %s\n", field)
					}
				}
			}
		}
	}

	return nil
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)pass(word)`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)hash`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)card`),
		regexp.MustCompile(`(?i)key`),
		regexp.MustCompile(`(?i)security`),
	}
}
