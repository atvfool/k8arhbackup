package main

import (
	"bytes"
	"fmt"
	"io"
	"k8arh/backup/options"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	dir, dirErr := os.UserHomeDir()

	var (
		configPath string
		origConfig []byte
	)

	if dirErr == nil {
		configPath = filepath.Join(dir, "k8arhbackup", "settings.config")
		var err error
		origConfig, err = os.ReadFile(configPath)
		if err != nil && !os.IsNotExist(err) {
			// The user has a config file but we couldn't read it
			// Report the error instead of ingnoring their config
			log.Fatal(err)
		} else if err != nil && os.IsNotExist(err) {
			// The config file does not exist
			// Report the error
			log.Fatal(err)
		}
	}

	// Use and perhap make changes to the config
	config := bytes.Clone(origConfig)
	configAsString := string(config)

	fmt.Println("-----------------------------------------------")
	fmt.Println("Load Config")
	fmt.Println("-----------------------------------------------")

	opts := options.NewFromString(configAsString)
	o := opts.GetValueByKey("overwrite")

	if len(o) > 1 {
		log.Fatal("overwrite option can only be defined once")
	}

	overwritebuff := o[0]
	overwritebuff = strings.ToLower(overwritebuff)

	overwrite := overwritebuff == "true"
	sources := opts.GetValueByKey("source")
	dests := opts.GetValueByKey("dest")

	fmt.Printf("Overwrite option: %t\n", overwrite)
	fmt.Printf("sources: \n")
	for _, i := range sources {
		fmt.Println(i)
	}
	fmt.Printf("dests: \n")

	for _, i := range dests {
		fmt.Println(i)
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("End Load Config")
	fmt.Println("-----------------------------------------------")

	fmt.Println("-----------------------------------------------")
	fmt.Println("Backup Files")
	fmt.Println("-----------------------------------------------")

	for _, i := range sources {
		fmt.Printf("Source File %q\n", i)
		srcFile, err := os.Open(i)
		ErrorCheck(err)
		defer srcFile.Close()
		_, fileName := filepath.Split(srcFile.Name())
		fmt.Printf("Filename %q \n", fileName)

		// Copy to each destination
		for _, j := range dests {
			fmt.Printf("dest folder %q\n", j)
			newPath := path.Join(j, fileName)
			_, err := os.Stat(newPath)
			fileExists := os.IsExist(err)
			removeExistingFile := fileExists && overwrite

			if fileExists && !overwrite {
				log.Printf("skipping, destinaton file %q already exists and overwrite is set to %t\n", fileName, overwrite)
			} else {
				// copy to a temporary file to prevent messing up an existing file
				tempFile, err := os.CreateTemp(j, "tempFile*")
				ErrorCheck(err)
				defer tempFile.Close() //os.Remove(tempFile.Name()) // Clean up

				fmt.Printf("temp file %q\n", tempFile.Name())
				fmt.Println("copying file")
				_, err = io.Copy(tempFile, srcFile) // Copy File
				ErrorCheck(err)

				// Close temp file
				if err := tempFile.Close(); err != nil {
					log.Fatal(err)
				}
				ErrorCheck(err)

				// Move temp to dest file
				if removeExistingFile {
					fmt.Printf("remove existing file %q\n", newPath)
					err := os.Remove(newPath)
					ErrorCheck(err)
				}
				fmt.Printf("rename temp file %q to %q\n", tempFile.Name(), newPath)
				err = os.Rename(tempFile.Name(), newPath)

				ErrorCheck(err)
			}
		}
	}
	fmt.Println("----------------------------")
	fmt.Println("Complete")
	fmt.Println("----------------------------")

}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
