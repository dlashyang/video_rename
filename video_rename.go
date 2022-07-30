package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	_, err = gen_rename_list(dir_name)
	if err != nil {
		log.Fatal(err)
	}

	err = rename_by_list("info.json")
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
		new_name := gen_new_name(old_name)
		list.V_file = append(list.V_file, rename_candidate{old_name, new_name})
	}

	//json_str, err = json.Marshal(list)
	json_str, err = json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Fatal(err)
		return json_str, err
	}

	write_rename_list_file("info.json", json_str)
	//fmt.Printf("%s\n", json_str)
	return json_str, nil
}

func gen_new_name(old_name string) string {
	name := strings.Split(old_name, ".")
	name[0] += "_new"
	new_name := strings.Join(name, ".")
	return new_name
}

func rename_by_list(json_file_name string) error {
	log.Println("rename below files:")
	list := rename_list{}

	info := read_rename_list_file(json_file_name)
	err := json.Unmarshal(info, &list)
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

func read_rename_list_file(file_name string) []byte {
	info, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}

	return info
}

func write_rename_list_file(json_file_name string, info []byte) error {
	filePtr, err := os.Create(json_file_name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer filePtr.Close()

	_, err = filePtr.Write(info)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
