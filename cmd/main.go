package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	targetPath := "./"
	var maxDepth int
	flag.IntVar(&maxDepth, "d", 10, "Max depth from CWD. Default is 10")
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), "Usage: rootdir [options] file file2 ...\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	fileNames := flag.Args()

	for i := 0; i < maxDepth; i++ {
		path, _ := filepath.Abs(targetPath)

		for _, fileName := range fileNames {
			if fileExists(path + "/" + fileName) {
				fmt.Printf("%s/%s\n", path, fileName)
				return
			}
		}

		if path == "/" {
			break
		}

		targetPath += "/.."
	}

}
