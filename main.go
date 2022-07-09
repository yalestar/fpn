package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var copyToClipboard bool

func main() {
	flag.BoolVar(&copyToClipboard, "c", false, "Copy to clipboard")

	flag.Parse()

	filenamesArg := flag.Args()

	if len(filenamesArg) == 0 {
		pwd, _ := os.Getwd()
		listEverything(pwd)
		os.Exit(0)
	}
	listTheseFiles(filenamesArg...)
}

func listEverything(path string) {
	var allFiles []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		filename := filepath.Join(path, file.Name())
		allFiles = append(allFiles, filename)
		fmt.Println(filename)
	}

	if copyToClipboard {
		toClipboard(allFiles)
	}
}

func toClipboard(ls []string) {
	clipboard.WriteAll(strings.Join(ls, "\n"))
	fmt.Println("And it's on the clipboard")
}

func listTheseFiles(filenames ...string) {

	var allFiles []string
	for _, f := range filenames {
		fullPath, _ := filepath.Abs(f)
		allFiles = append(allFiles, fullPath)
		fmt.Println(fullPath)
	}

	if copyToClipboard {
		toClipboard(allFiles)
	}
}
