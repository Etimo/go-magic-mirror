package support

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const expectedBasePathName string = "go-magic-mirror"

func AttemptToFindBasePath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return ".", err
	}

	fmt.Println(path)
	//Note: filepath.SplitList does not split string correctly by file separator
	sep := string(filepath.Separator)
	directories := strings.Split(path, sep)
	workDirIndex := -1
	for i, dir := range directories {
		if strings.Contains(dir, "go-magic-mirror") {
			workDirIndex = i
		}
	}
	if workDirIndex < 1 {
		return "./", nil
	}
	workDir := filepath.Join(directories[:(workDirIndex + 1)]...)
	fmt.Println("Guessed workdir: " + workDir)
	return "/" + workDir, nil

}
