package archive_compress

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/dto"
	"github.com/iliyasali2107/archiver/internal/helpers/random"
	"github.com/iliyasali2107/archiver/internal/services/archive"
)

type ArchiveCompressCtrl struct {
	svc ArchiveCompressSvc
}

type ArchiveCompressSvc interface {
	Compress(dto.ArchiveCompressRequest) (dto.ArchiveCompressResponse, error)
	Clear(filename string) error
}

func NewArchiveCompressCtrl(svc ArchiveCompressSvc) *ArchiveCompressCtrl {
	return &ArchiveCompressCtrl{
		svc: svc,
	}
}

const (
	formKey = "files[]"
)

func (acc *ArchiveCompressCtrl) Compress(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get multipart form"})
		return
	}
	files, ok := form.File[formKey]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there is no files in request"})
		return
	}

	req := dto.ArchiveCompressRequest{Files: files}
	res, err := acc.svc.Compress(req)
	if err != nil {
		if err == archive.ErrNotAllowedMIMEType {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you send file with MIMEType that is not allowed"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	buf := res.Buffer

	disposition := fmt.Sprintf("attachment; filename=%s", random.RandomZipFileName())
	c.Writer.Header().Set("Content-Type", "application/zip")
	c.Writer.Header().Set("Content-Disposition", disposition)

	_, err = io.Copy(c.Writer, buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
