package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"io"
)

var file_types = []string {
	".wav",
	".mp3",
	".flac",
	".mp4",
	".m4a",
}

func processDirectory(f io.Writer, name string, num_files *int) int {
	dir,err := ioutil.ReadDir(name)
	if (err != nil) {
		fmt.Fprintln(os.Stderr, "Error reading directory: ", name)
		os.Exit(1)
	}
	for _, i := range(dir) {
		if (i.IsDir()) {
			processDirectory(f, name + "/" + i.Name(), num_files)
			continue
		}
		lower_name := strings.ToLower(i.Name())
		for _, z := range (file_types) {
			if(strings.Contains(lower_name, z) == true) {
				fmt.Fprintf(f, "%s\n", name+"/"+i.Name())
				*num_files ++
				continue
			}
		}
	}
	return *num_files
}

func main() {
	fmt.Println("M3U_Gen Go v1.0")
	if(len(os.Args) != 4) {
		fmt.Fprintln(os.Stderr,"Program requires 3 arguments\nm3ugen playlist.m3u directory [ -s for sort -r for reverse sort ]")
		os.Exit(1)
	}
	f_output,err := os.Create(os.Args[1])
	if(err != nil) {
		fmt.Fprintln(os.Stderr,"Error could not open file to write playlist too: ", err)
		os.Exit(1)
	}
	num_files := 0
	counted := processDirectory(f_output, os.Args[2], &num_files)
	f_output.Close()
	fmt.Println("Processed ", counted, " files written to ", os.Args[1])
	os.Exit(0)
}