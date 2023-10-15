package main

import (
	"archive/zip"
	"fmt"
	"log"

	"github.com/iliyasali2107/archiver/internal/controllers/archive_info"
)

func main() {
	testFile := "test.zip"

	r, err := zip.OpenReader(testFile)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Close()

	res := archive_info.ArchiveInfoResponse{}

	res.

	for _, f := range r.File {

		fmt.Println(f.Name)
	}

}
