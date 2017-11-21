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
	fmt.Printf("%s := []string {\n", name)
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
			fmt.Fprintf(os.Stdout, "\"%s\",\n", scanlines[i])
		} else {
			fmt.Fprintf(os.Stdout, "\"%s\" }\n", scanlines[i])
		}
	}
	fmt.Fprintf(os.Stdout, "\n%s_count := %d\n\n", name, count)
}

func main() {

	testCode()

	if len(os.Args) == 2 {
		readData(os.Args[1])

	} else {

		fmt.Println("Arguments: varname")
		os.Exit(1)
	}
}

func testCode() {
	test := []string {
		"package main",
		"",
		"",
		"import (",
		"	\"fmt\"",
		"	\"os\"",
		"	\"bufio\"",
		")",
		"",
		"",
		"func fixEscape(s string) string {",
		"",
		"	var temp string",
		"",
		"	for i := 0; i < len(s); i++  {",
		"		if(s[i] == '\\\\' || s[i] == '\"') {",
		"			temp += string('\\\\')",
		"			temp += string(s[i])",
		"		} else {",
		"			temp += string(s[i])",
		"		}",
		"",
		"	}",
		"	return temp",
		"",
		"}",
		"",
		"",
		"func readData(name string) {",
		"	fmt.Printf(\"%s := []string {\\n\", name)",
		"	scan := bufio.NewScanner(os.Stdin)",
		"	var scanlines []string",
		"	count := 0",
		"	for scan.Scan() {",
		"		var t string = fixEscape(scan.Text())",
		"		scanlines = append(scanlines, t)",
		"		count ++",
		"	}",
		"	for i := 0; i < count; i++ {",
		"		if (i < count-1) {",
		"			fmt.Fprintf(os.Stdout, \"\\\"%s\\\",\\n\", scanlines[i])",
		"		} else {",
		"			fmt.Fprintf(os.Stdout, \"\\\"%s\\\" }\\n\", scanlines[i])",
		"		}",
		"	}",
		"	fmt.Fprintf(os.Stdout, \"\\n%s_count := %d\\n\\n\", name, count)",
		"}",
		"",
		"func main() {",
		"",
		"	if len(os.Args) == 2 {",
		"		readData(os.Args[1])",
		"",
		"	} else {",
		"",
		"		fmt.Println(\"Arguments: varname\")",
		"		os.Exit(1)",
		"	}",
		"}",
		"" }

	test_count := 60
	for i := 00; i < test_count; i++ {
		fmt.Printf("%s\n", test[i])
	}
}