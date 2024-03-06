package files

import (
	"blockchain_learningtool/shared"
	"errors"
	"fmt"
	"os"
)

func CreateDirs() error {
	toCreateDirs := []string{"output/keys", "output/blocks", "output/walktrough"}
	for _, directoryPath := range toCreateDirs {
		if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
			err := os.MkdirAll(directoryPath, 0755)
			if err != nil {
				return errors.New(fmt.Sprintf("Error creating directory: %s", err))
			}
			if shared.Myparameters.Verbose {
				fmt.Printf("Directory %s created successfully!\n", directoryPath)
			}

		} else {
			if shared.Myparameters.Verbose {
				fmt.Printf("Directory %s already exists.\n", directoryPath)
			}

		}
	}

	return nil
}
