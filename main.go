package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	filename := flag.String("o", "", `output filename (default: "$dir/.dockerignore")`)
	dir := flag.String("dir", ".", "build context root directory")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("dignore: ")

	ignore := *filename
	if ignore == "" {
		ignore = filepath.Join(*dir, ".dockerignore")
	}

	if ignore != "-" {
		f, err := os.Create(ignore)
		if err != nil {
			log.Fatalf("failed to create file(path=%v): %v", ignore, err)
		}
		defer f.Close()
		out = f
	}

	directories := flag.Args()
	fileInfos, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatalf("failed to readdir(dir=%v): %v", *dir, err)
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
			appendPrefix(*dir, n)
		} else {
			fmt.Fprintln(out, n)
		}
	}
}

func appendPrefix(dir, name string) {
	f, err := os.Open(filepath.Join(dir, name, ".dockerignore"))
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		log.Fatalf("failed to open .dockerignore on %v: %v", filepath.Join(dir, name), err)
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
		fmt.Fprintln(out, path.Join(name, t))
	}
}
