package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type rename_candidate struct {
	Old_name string `json:"old_name"`
	New_name string `json:"new_name"`
}

type rename_list struct {
	Path_base string             `json:"path"`
	V_file    []rename_candidate `json:"list"`
}

func main() {
	log.Println("Start.")
	list := rename_list{}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	dir_name, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	list.Path_base = dir_name
	fmt.Println(dir_name)

	for _, file := range files {
		if true == file.IsDir() {
			continue
		}
		if true != strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
			continue
		}
		if true == strings.Contains(file.Name(), "_new") {
			continue
		}

		old_name := file.Name()
		fmt.Println(filepath.Join(dir_name, old_name))
		name := strings.Split(old_name, ".")
		name[0] += "_new"
		new_name := strings.Join(name, ".")
		fmt.Println(filepath.Join(dir_name, new_name))
		list.V_file = append(list.V_file, rename_candidate{old_name, new_name})
	}

	json_str, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", json_str)

	rename_by_list(json_str)
	log.Print("Done.")
}

func rename_by_list(rlist []byte) error {
	log.Println("Start to rename.")
	list := rename_list{}
	err := json.Unmarshal(rlist, &list)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range list.V_file {
		fmt.Println("old name: ", filepath.Join(list.Path_base, file.Old_name))
		fmt.Println("new name: ", filepath.Join(list.Path_base, file.New_name))
		/*
			err := os.Rename(filepath.Join(dir_name, old_name), filepath.Join(dir_name, new_name))
			if err != nil {
				log.Fatal(err)
			}
		*/
	}
	return nil
}
