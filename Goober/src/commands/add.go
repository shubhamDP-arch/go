package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jaqen/goober/src/utils"
)

func Add(files []string) {
	if len(files) == 1 && files[0] == "." {
		files = utils.GetAllfilesInDirectory(".")
	}

	indexPath := ".goobers/index"
	stagedFiles := readIndex(indexPath)

	indexFile, err := os.OpenFile(indexPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening index file:", err)
		return
	}
	defer indexFile.Close()

	for _, file := range files {
		if utils.IsDirectory(file) {
			continue
		}

		if _, exists := stagedFiles[file]; exists {
			fmt.Printf("File already staged: %s\n", file)
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		fileHash := utils.GenerateHash(string(content))
		_, err = indexFile.WriteString(fmt.Sprintf("%s %s\n", file, fileHash))
		if err != nil {
			fmt.Printf("Error writing to index file: %v\n", err)
			continue
		}

		fmt.Printf("Added the file: %s\n", file)
	}
}

func readIndex(indexPath string) map[string]string {
	stagedFiles := make(map[string]string)

	file, err := os.Open(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return stagedFiles
		}
		fmt.Println("Error reading index file:", err)
		return stagedFiles
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			stagedFiles[parts[0]] = parts[1] 
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning index file:", err)
	}

	return stagedFiles
}
