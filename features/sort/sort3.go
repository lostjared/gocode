package main

import(
	"fmt"
	"sort"
)

type Data struct {
	index int
	name string
}

type sortedData struct {
	names []*Data
	less func(i, j *Data) bool

}

func (s sortedData) Len() int {
	return len(s.names)
}
func (s sortedData) Less(i int, j int) bool {
	return s.less(s.names[i], s.names[j]);
}
func (s sortedData) Swap(i, j int) {
	s.names[i], s.names[j] = s.names[j], s.names[i]
}

func printNames(name []*Data) {
	fmt.Println("Number list: ")
	for _, i := range(name) {
		fmt.Println(i.index, ": ", i.name)
	}
}

func main() {
	var values = []*Data{
		{6, "Six"},
		{1, "One"},
		{3, "Three"},
		{2, "Two"},
		{4, "Four"},
		{5, "Five"},
	}
	fmt.Println("Unsorted: ")
	printNames(values)
	dat := sortedData{values, func(i, j *Data) bool {
		if i.index != j.index {
			return i.index < j.index;
		}
		return false;
	}}
	sort.Sort(dat)
	fmt.Println("Sorted: ")
	printNames(dat.names)

}