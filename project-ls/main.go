package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	all := flag.Bool("a", false, "show all")
	flag.Parse()

	for _, file := range listFiles("testdata", *all) {
		fmt.Println(file)
	}

}

func listFiles(dirname string, all bool) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !all && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		dirs = append(dirs, f.Name())
	}

	return dirs
}
