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

	dir_name, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	json_str, err := gen_rename_list(dir_name)
	if err != nil {
		log.Fatal(err)
	}

	err = rename_by_list(json_str)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done.")
}

func gen_rename_list(path_base string) ([]byte, error) {
	log.Println("generate file list in: ", path_base)

	list := rename_list{}
	list.Path_base = path_base
	var json_str []byte

	files, err := os.ReadDir(path_base)
	if err != nil {
		log.Fatal(err)
		return json_str, err
	}

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
		fmt.Println(filepath.Join(path_base, old_name))
		name := strings.Split(old_name, ".")
		name[0] += "_new"
		new_name := strings.Join(name, ".")
		fmt.Println(filepath.Join(path_base, new_name))
		list.V_file = append(list.V_file, rename_candidate{old_name, new_name})
	}

	json_str, err = json.Marshal(list)
	if err != nil {
		log.Fatal(err)
		return json_str, err
	}
	fmt.Printf("%s\n", json_str)
	return json_str, nil
}

func rename_by_list(rlist []byte) error {
	log.Println("Start to rename.")
	list := rename_list{}
	err := json.Unmarshal(rlist, &list)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, file := range list.V_file {
		fmt.Println("rename: ", filepath.Join(list.Path_base, file.Old_name), "-->", filepath.Join(list.Path_base, file.New_name))
		/*
			err := os.Rename(filepath.Join(dir_name, old_name), filepath.Join(dir_name, new_name))
			if err != nil {
				log.Fatal(err)
			}
		*/
	}
	return nil
}
