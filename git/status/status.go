package status

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// GetStatus returns changed and deleted files
func GetStatus() ([]string, []string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return nil, nil, err
	}

	var changedFiles, deletedFiles []string
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 3 {
			status := line[:2]
			file := strings.TrimSpace(line[2:])
			switch status {
			case " D":
				deletedFiles = append(deletedFiles, file)
			case "??", " M", "MM", "AM":
				changedFiles = append(changedFiles, file)
			}
		}
	}
	return changedFiles, deletedFiles, nil
}

// GetStagedFiles returns files that are staged for commit
func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var stagedFiles []string
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 3 {
			status := line[:2]
			file := strings.TrimSpace(line[2:])
			if status == "A" {
				stagedFiles = append(stagedFiles, file)
			}
		}
	}
	return stagedFiles, nil
}
