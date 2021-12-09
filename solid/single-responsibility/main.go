package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// Single Responsibility principle
// A type should have only one reason to change
// Separation of concerns - different types / packages handling
// different, independent tasks / problems

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Separation of Concerns

// persistence hhere is not a responsibility of this type
func (j *Journal) Save(filename string) {
	// ...
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("No any news")
	fmt.Println(j.String())

	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
