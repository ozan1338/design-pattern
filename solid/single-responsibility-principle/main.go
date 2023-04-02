package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

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
	// ..
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// This one break single principle
// separation of concerns
// God Object
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url url.URL) {}

// Just Create Another Func To Do persistanve cause journal just need doing about entries
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

// or just create another struct for implementing about persitance
type Persistance struct {
	lineSeparator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I studied today")
	j.AddEntry("I Sahur Today")
	fmt.Println(j.String())

	// So when to deal with persistance
	SaveToFile(&j, "journal.txt")

	// Or Create persistance object
	p := Persistance{}
	p.lineSeparator = "\r\n"
	p.SaveToFile(&j, "journal-1.txt")
}
