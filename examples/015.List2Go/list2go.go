package main

import (
	"fmt"
	"os"
	"bufio"
)

func fixEscape(s string) string {
	var temp string
	for i := 0; i < len(s); i++  {
		if(s[i] == '\\' || s[i] == '"') {
			temp += string('\\')
			temp += string(s[i])
		} else {
			temp += string(s[i])
		}

	}
	return temp
}

func readData(name string) {
	fmt.Printf("func %s () (int, []string) {\n\n", name)
	fmt.Printf("\t%s_arr := []string {\n", name)
	scan := bufio.NewScanner(os.Stdin)
	var scanlines []string
	count := 0
	for scan.Scan() {
		var t string = fixEscape(scan.Text())
		scanlines = append(scanlines, t)
		count ++
	}
	for i := 0; i < count; i++ {
		if (i < count-1) {
			fmt.Fprintf(os.Stdout, "\t\"%s\",\n", scanlines[i])
		} else {
			fmt.Fprintf(os.Stdout, "\t\"%s\" }\n", scanlines[i])
		}
	}
	fmt.Fprintf(os.Stdout, "\n\t%s_count := %d\n", name, count)
	fmt.Fprintf(os.Stdout, "\treturn %s_count, %s_arr\n\n}\n\n", name, name)
}

func main() {
	if len(os.Args) == 2 {
		readData(os.Args[1])

	} else {
		fmt.Println("Arguments: varname")
		os.Exit(1)
	}
}
