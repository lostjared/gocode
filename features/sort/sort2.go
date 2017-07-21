package main

import(
	"fmt"
	"sort"
)

type Name struct {
	first, last string
}

type sortData struct {
	names []*Name
}

func (s sortData) Len() int {
	return len(s.names)
}
func (s sortData) Less(i int, j int) bool {
	if s.names[i].last != s.names[j].last {
		return s.names[i].last < s.names[j].last;
	}
	if s.names[i].first != s.names[j].first {
		return s.names[i].first < s.names[j].first;
	}
	return false;
}
func (s sortData) Swap(i, j int) {
	s.names[i], s.names[j] = s.names[j], s.names[i]
}

func printNames(name []*Name) {
	fmt.Println("Name list: ")
	for _, i := range(name) {
		fmt.Println(i.last, ", ", i.first)
	}
}

func main() {
	fmt.Println("Sorting names")
	var names = []*Name {
		{"Jared", "Bruni"},
		{"Johnny", "Appleseed"},
		{"Bilbo", "Baggins"},
		{"Freddy", "Krueger"},
	}
	printNames(names)
	data := sortData{names}
	sort.Sort(data)
	printNames(names)
}