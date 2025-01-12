package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jaqen/goober/src/utils"
)

func Commit(message string) {
	index := ".goobers/index"
	if _, err := os.Stat(index); os.IsNotExist(err) {
		fmt.Println("Files are not staged. Committing the cwd")
		commitAllFiles(message)
		return
	}

	stagedFiles, err := ioutil.ReadFile(index)
	if err != nil {
		fmt.Println("Error reading the index file:", err)
		return
	}

	latestCommitHash := getLatestCommitHash()
	if latestCommitHash != "" {

		latestCommitContent := getCommitContent(latestCommitHash)

		stagedContentHash := getStagedFilesHash(string(stagedFiles))
		latestCommitContentHash := getStagedFilesHash(latestCommitContent)

		if stagedContentHash == latestCommitContentHash {
			fmt.Println("No changes detected, commit skipped.")
			return
		}
	}

	CreateCommitObject(string(stagedFiles), message)
}

func commitAllFiles(message string) {
	files := utils.GetAllfilesInDirectory(".")
	var commitContent string

	for _, file := range files {
		if utils.IsDirectory(file) {
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		fileHash := utils.GenerateHash(string(content))
		commitContent += fmt.Sprintf("%s %s\n", file, fileHash)
	}

	CreateCommitObject(commitContent, message)
}

func CreateCommitObject(stagedFiles string, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	commitHash := utils.GenerateHash(stagedFiles + message + timestamp)
	commitPath := fmt.Sprintf(".goobers/objects/%s", commitHash)

	// Create commit file
	file, err := os.Create(commitPath)
	if err != nil {
		fmt.Printf("Error creating commit: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("commit %s\n%s\n%s", commitHash, message, stagedFiles))
	fmt.Printf("Committed with hash: %s\n", commitHash)
}

func getLatestCommitHash() string {
	files, err := ioutil.ReadDir(".goobers/objects/")
	if err != nil {
		fmt.Println("Error reading .goobers/objects:", err)
		return ""
	}

	if len(files) == 0 {

		return ""
	}

	latestCommitFile := files[len(files)-1]
	return latestCommitFile.Name()
}

func getCommitContent(commitHash string) string {
	commitPath := fmt.Sprintf(".goobers/objects/%s", commitHash)
	content, err := ioutil.ReadFile(commitPath)
	if err != nil {
		fmt.Printf("Error reading commit content for hash %s: %v\n", commitHash, err)
		return ""
	}
	return string(content)
}

func getStagedFilesHash(commitContent string) string {

	commitContent = normalizeContent(commitContent)

	lines := strings.Split(commitContent, "\n")

	if len(lines) < 3 {
		return "" 
	}

	return strings.Join(lines[2:], "\n")
}

func normalizeContent(content string) string {

	content = strings.TrimSpace(content)

	content = strings.Join(strings.Fields(content), " ")
	return content
}

func CompareHash(commit1, commit2 string) bool {
	commit1Hash := getStagedFilesHash(commit1)
	commit2Hash := getStagedFilesHash(commit2)

	return commit1Hash != commit2Hash
}
