// ma creates files and their needed directories automatically.
//
// Examples
// 1. `ma a/b/c`       # `mkdir -p a/b/c`
// 2. `ma a/b/c/d.txt` # `mkdir -p a/b/c && touch a/b/c/d.txt`
package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: ma <things to create>")
		os.Exit(1)
	}

	d, f := split(os.Args[1])

	if d != "" {
		if err := os.MkdirAll(d, 0755); err != nil {
			log.Printf("error creating directories %v\n", err)
			os.Exit(1)
		}
	}

	if f != "" {
		ff := filepath.Join(d, f)
		if fileExists(ff) {
			log.Printf("file already exists %q", ff)
			os.Exit(0)
		}

		fh, err := os.Create(ff)
		if err != nil {
			log.Printf("error creating file %v\n", err)
		}

		defer fh.Close()

	}

}

// isDir takes a guess if input is a dir or file
func isDir(input string) bool {
	r := regexp.MustCompile(`\.\w+$`)
	return !r.MatchString(input)
}

// split splits input to its directory and file parts
func split(input string) (string, string) {
	if isDir(input) {
		return input, ""
	}

	return filepath.Split(input)
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
