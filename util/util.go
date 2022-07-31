package util

import (
	"encoding/json"
	"fmt"
	"github.com/tigrato/go-mediainfo"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type rename_candidate struct {
	Old_name string `json:"old_name"`
	New_name string `json:"new_name"`
}

type rename_list struct {
	Path_base string             `json:"path"`
	V_file    []rename_candidate `json:"list"`
}

/*
func main() {
	log.Println("Start.")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough arguments.")
	}
	dir_name := args[0]

		dir_name, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

	_, err := Gen_candidate_list(dir_name)
	if err != nil {
		log.Fatal(err)
	}

	err = Rename_by_list("info.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done.")
}
*/

func Gen_candidate_list(path_base string, list_name string) error {
	log.Println("generate file list in: ", path_base)

	list := rename_list{}
	list.Path_base = path_base

	files, err := os.ReadDir(path_base)
	if err != nil {
		log.Fatal(err)
		return err
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
		new_name := gen_new_name(path_base, old_name)
		list.V_file = append(list.V_file, rename_candidate{old_name, new_name})
	}

	json_str, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	write_rename_list_file(list_name, json_str)
	return nil
}

func gen_new_name(path_base string, old_name string) string {
	mi := mediainfo.New()
	if err := mi.Open(filepath.Join(path_base, old_name)); err != nil {
		log.Fatal(err)
	}

	en_date := mi.GetKind(mediainfo.StreamGeneral, 0, "Encoded_Date", mediainfo.InfoText)
	t, _ := time.Parse("UTC 2006-01-02 15:04:05", en_date)
	cap_date := t.Format("20060102_150405")

	name := strings.Split(old_name, ".")
	name[0] = cap_date
	new_name := strings.Join(name, ".")
	return new_name
}

func Rename_by_list(list_name string, flag_dry_run bool) error {
	log.Println("rename below files:")
	list := rename_list{}

	info := read_rename_list_file(list_name)
	err := json.Unmarshal(info, &list)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, file := range list.V_file {
		if flag_dry_run {
			fmt.Println("rename: ", filepath.Join(list.Path_base, file.Old_name), "-->", filepath.Join(list.Path_base, file.New_name))
		} else {
			err := os.Rename(filepath.Join(list.Path_base, file.Old_name), filepath.Join(list.Path_base, file.New_name))
			if err != nil {
				log.Fatal(err)
			}
		}
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
