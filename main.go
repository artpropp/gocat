package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("At least one parameter with a filename is required")
		os.Exit(1)
	}

	for _, fileName := range os.Args[1:] {
		f, err := os.Open(fileName)
		if err != nil {
			log.Printf("Error while trying to open file: %s\n%s\n", fileName, err)
			f.Close()
			continue
		}

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Error copying file %s to stdout\n%s", fileName, err)
		}
		f.Close()
	}

}