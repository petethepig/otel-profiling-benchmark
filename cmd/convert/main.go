package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filepaths, err := filepath.Glob("./profiles/src/*")

	if err != nil {
		panic(err)
	}

	for _, path := range filepaths {
		if strings.HasSuffix(path, ".collapsed") {
			copyFile(path, strings.Replace(path, "src", "intermediary", 1))
		} else {
			panic("Unexpected file extension")
		}
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
