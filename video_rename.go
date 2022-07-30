package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	dir_name, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir_name)

	for _, file := range files {
		if true == file.IsDir() {
			continue
		}
		old_name := file.Name()
		if true == strings.HasSuffix(strings.ToLower(old_name), ".mp4") {
			fmt.Println(filepath.Join(dir_name, old_name))
			if true == strings.Contains(old_name, "_new") {
				continue
			}
			name := strings.Split(old_name, ".")
			name[0] += "_new"
			new_name := strings.Join(name, ".")
			fmt.Println(filepath.Join(dir_name, new_name))
			err := os.Rename(filepath.Join(dir_name, old_name), filepath.Join(dir_name, new_name))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Print("Done")
}
