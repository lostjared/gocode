package main

 import(
	 "fmt"
	 "os"
	 "bufio"
	 "sort"
 )

type sortString struct {
	lines []string
}

func (s sortString) Len() int {
	return len(s.lines)
}

func (s sortString) Less(i int, j int) bool {
	return s.lines[i] < s.lines[j]
}

func (s sortString) Swap(i int, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}


func main() {
	if(len(os.Args) != 3) {
		fmt.Fprintln(os.Stderr, os.Args[0], " source output")
		fmt.Fprintf(os.Stderr, "%s", "Requires two arguments")
		os.Exit(1)
	}
	fptr,err := os.Open(os.Args[1])
	if(err != nil) {
		fmt.Fprintf(os.Stderr, "%s%s", "Error could not open file: ", os.Args[1])
		os.Exit(1)
	}
	input := bufio.NewScanner(fptr)
	var lines = []string {}
	for input.Scan() {
		lines = append(lines, input.Text())
	}
	line_data := sortString{lines}
	sort.Sort(line_data)
	fptr.Close()
	ofptr,err := os.Create(os.Args[2])
	if(err != nil) {
		fmt.Fprintf(os.Stderr, "Error could not create file: %s", os.Args[2])
		os.Exit(1)
	}
	for i := 0; i < len(line_data.lines); i++ {
		fmt.Fprintf(ofptr, "%s\n", line_data.lines[i])
	}
	fmt.Println("Wrote sorted list to file: ", os.Args[2])
	ofptr.Close()
}