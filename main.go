package main

import (
	"archive/zip"
	"fmt"
	"log"
)

func main() {
	testFile := "test.zip"

	r, err := zip.OpenReader(testFile)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	for _, f := range r.File {

		fmt.Println(f.Name)
	}

}
