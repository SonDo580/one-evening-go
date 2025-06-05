package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	all := flag.Bool("a", false, "list all files")
	flag.Parse()

	files := listFiles("testdata", *all)
	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(dirname string, all bool) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		name := f.Name()
		if strings.HasPrefix(name, ".") && !all {
			continue
		}
		dirs = append(dirs, name)
	}

	return dirs
}
