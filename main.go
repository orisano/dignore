package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("dignore: ")

	directories := flag.Args()
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal("failed to readdir:", err)
	}
	required := map[string]struct{}{}
	for _, d := range directories {
		required[d] = struct{}{}
	}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			continue
		}
		n := fi.Name()
		if _, ok := required[n]; ok {
			appendPrefix(n)
		} else {
			fmt.Println(n)
		}
	}
}

func appendPrefix(name string) {
	f, err := os.Open(filepath.Join(name, ".dockerignore"))
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatalf("failed to open .dockerignore on %v: %v", name, err)
	}
	defer f.Close()
	bs := bufio.NewScanner(f)
	for bs.Scan() {
		t := bs.Text()
		if t == "" {
			continue
		}
		if strings.HasPrefix(t, "#") {
			continue
		}
		if strings.HasPrefix(t, "!") {
			fmt.Fprintf(os.Stderr, "warning: unsupported ! syntax: %v", t)
			continue
		}
		fmt.Println(filepath.Join(name, t))
	}
}
