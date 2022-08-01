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

type renameCandidate struct {
	OldName string `json:"video_file"`
	NewName string `json:"new_name"`
}

type candidateList struct {
	PathBase      string            `json:"path_base"`
	CandidateList []renameCandidate `json:"candidate_list"`
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

// GenCandidateList to generate the list file under given path
func GenCandidateList(basePath string, listFileName string) error {
	log.Println("generate file list in: ", basePath)

	list := candidateList{}
	list.PathBase = basePath

	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToLower(file.Name()), ".mp4") {
			continue
		}
		if strings.Contains(file.Name(), "_new") {
			continue
		}

		currentName := file.Name()
		newName := genNewName(basePath, currentName)
		list.CandidateList = append(list.CandidateList, renameCandidate{currentName, newName})
	}

	jsonString, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Fatal(err)
		return err
	}

	writeListtoFile(listFileName, jsonString)
	return nil
}

func genNewName(path string, fileName string) string {
	mi := mediainfo.New()
	if err := mi.Open(filepath.Join(path, fileName)); err != nil {
		log.Fatal(err)
	}

	encodedDate := mi.GetKind(mediainfo.StreamGeneral, 0, "Encoded_Date", mediainfo.InfoText)
	t, _ := time.Parse("UTC 2006-01-02 15:04:05", encodedDate)
	captureDate := t.Format("20060102_150405")

	name := strings.Split(fileName, ".")
	name[0] = captureDate
	return strings.Join(name, ".")
}

/*
RenamebyList reads list file to comple the rename.
If the flagDryRun is true, it prints only.
*/
func RenamebyList(listFileName string, flagDryRun bool) error {
	log.Println("rename below files:")
	list := candidateList{}

	info := readListFile(listFileName)
	err := json.Unmarshal(info, &list)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, file := range list.CandidateList {
		fmt.Println("rename: ", filepath.Join(list.PathBase, file.OldName), "-->", filepath.Join(list.PathBase, file.NewName))
		if !flagDryRun {
			err := os.Rename(filepath.Join(list.PathBase, file.OldName), filepath.Join(list.PathBase, file.NewName))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}

func readListFile(listFileName string) []byte {
	info, err := ioutil.ReadFile(listFileName)
	if err != nil {
		log.Fatal(err)
	}

	return info
}

func writeListtoFile(listFileName string, listContent []byte) error {
	filePtr, err := os.Create(listFileName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer filePtr.Close()

	_, err = filePtr.Write(listContent)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
