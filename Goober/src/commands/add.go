package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jaqen/goober/utils"
)

func Add(files []string) {

	if len(files) == 1 && files[0] == "." {
		files = utils.GetAllfilesInDirectory(".")
	}
	indexPath := ".goobers/index"
	indexFile, err := os.OpenFile(indexPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening index file:", err)
		return
	}
	defer indexFile.Close()
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}
		fileHash := utils.GenerateHash((string(content)))
		indexFile.WriteString(fileHash)
		fmt.Printf("Added the file: %s\n", file)
	}

}
