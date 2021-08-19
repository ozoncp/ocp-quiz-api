package main

import (
	"bytes"
	"fmt"
	"os"
)

func FileReader(files []string) {
	fileReader := func(fname string) (string, error) {
		file, err := os.Open(fname)
		if err != nil {
			return "", err
		}

		defer file.Close()

		data := new(bytes.Buffer)
		if _, err = data.ReadFrom(file); err != nil {
			return "", err
		}

		return data.String(), nil
	}

	for _, file := range files {
		_, err := fileReader(file)
		if err != nil {
			fmt.Printf("fileReader error: %v", err)
		}
	}
}

func main() {

}
