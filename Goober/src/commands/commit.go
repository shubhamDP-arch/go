package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/jaqen/goober/src/utils"
)

func Commit(message string) {
	index := ".goobers/index"

	if _, err := os.Stat(index); os.IsNotExist(err) {
		fmt.Println("Files are not staged Commiting the cwd")
	}

	stagedFiles, err := ioutil.ReadFile(index)
	if err != nil {
		fmt.Println("Error bro")
	}

	CreateCommitObject(string(stagedFiles), message)
}

func CreateCommitObject(stagedFiles string, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	commitHash := utils.GenerateHash(stagedFiles+ message+ timestamp)
	commitPath := fmt.Sprintf(".goobers/objects/%s\n", commitHash)
	file, err := os.Create(commitPath)
	if err != nil {
		fmt.Printf("Problem commiting")
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("commit %s\n%s\n%s", commitHash, message, stagedFiles))
	fmt.Printf("Committed with hash: %s\n", commitHash)
}
