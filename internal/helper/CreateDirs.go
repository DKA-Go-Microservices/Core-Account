package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

func Dirs() {
	// Get the path of the currently executing program
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// Get the directory of the executable
	execDir := filepath.Dir(execPath)

	// Define the folders to create
	folders := []string{"assets", "config"}

	for _, folder := range folders {
		folderPath := filepath.Join(execDir, folder)
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			// Create the folder if it does not exist
			err := os.Mkdir(folderPath, 0755)
			if err != nil {
				return
			}
		}
	}
}
