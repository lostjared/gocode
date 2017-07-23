
/*

Go programming practice example
M3U Generator (by Filename)
The Program should be run from the directory you want to generate the playlist for.
It ues relative path names for the filenames it outputs to the list

m3ugen playlistname.m3u path

I use the program by entering the directory I want to generate a playlist for and

$ m3ugen playlist.m3u .

- Jared

*/


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
	if(len(os.Args) != 3) {
		fmt.Fprintln(os.Stderr,"Program requires 2 arguments\nm3ugen playlist.m3u directory")
		os.Exit(1)
	}
	f_output,err := os.Create(os.Args[1])
	if(err != nil) {
		fmt.Fprintln(os.Stderr,"Error could not open file to write playlist to: ", err)
		os.Exit(1)
	}
	num_files := 0
	counted := processDirectory(f_output, os.Args[2], &num_files)
	f_output.Close()
	fmt.Println("Processed ", counted, " files written to ", os.Args[1])
	os.Exit(0)
}