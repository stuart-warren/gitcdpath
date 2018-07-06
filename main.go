package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, arg := range args {
		must(tree(arg))
	}
}

func tree(path string) error {
	set := make(map[string]struct{})
	err := filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() && fi.Name() == ".git" {
			set[filepath.Dir(filepath.Dir(path))] = struct{}{}
			return filepath.SkipDir
		}
		return nil
	})
	fmt.Println(joinSet(set))
	return err
}

func joinSet(set map[string]struct{}) string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return strings.Join(keys, string(filepath.ListSeparator))
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
