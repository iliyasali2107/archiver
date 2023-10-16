package archive_compress

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/dto"
	"github.com/iliyasali2107/archiver/internal/helpers/random"
)

type ArchiveCompressCtrl struct {
	svc ArchiveCompressSvc
}

type ArchiveCompressSvc interface {
	Compress(dto.ArchiveCompressRequest) (dto.ArchiveCompressResponse, error)
}

func NewArchiveCompressCtrl(svc ArchiveCompressSvc) *ArchiveCompressCtrl {
	return &ArchiveCompressCtrl{
		svc: svc,
	}
}

const formKey = "files[]"
const zipStoragePath = "archives/"

func (acc *ArchiveCompressCtrl) Compress(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is no files in request"})
		return
	}
	files, ok := form.File[formKey]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is no files in request"})
		return
	}

	randZipName := random.RandomZipFileName()
	filePath := zipStoragePath + randZipName
	zipFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		defer file.Close()

		// f := &zip.FileHeader{Name: fileHeader.Filename, Method: zip.Deflate}
		createdDest, err := zipWriter.Create(fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		_, err = io.Copy(createdDest, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

	}

	fileInfo, err := zipFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	disposition := fmt.Sprintf("attachment; filename=%s", fileInfo.Name())
	c.Writer.Header().Set("Content-Type", "application/zip")
	c.Writer.Header().Set("Content-Disposition", disposition)
	// c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()-int64(100000)))

	c.File(filePath)
	// c.FileAttachment(filePath, fileInfo.Name())

}
