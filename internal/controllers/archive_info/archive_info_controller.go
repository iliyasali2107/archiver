package archive_info

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iliyasali2107/archiver/internal/dto"
)

type ArchiveInfoCtrl struct {
	svc ArchiveInfoSvc
}

type ArchiveInfoSvc interface {
	GetArchiveInfo(dto.ArchiveInfoRequest) (dto.ArchiveInfoResponse, error)
}

func NewArchiveInfoCtrl(svc ArchiveInfoSvc) *ArchiveInfoCtrl {
	return &ArchiveInfoCtrl{
		svc: svc,
	}
}

const formKey = "file"

func (aic *ArchiveInfoCtrl) GetArchiveInfo(c *gin.Context) {
	file, err := c.FormFile(formKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "your request has no file"})
	}

	res, err := aic.svc.GetArchiveInfo(dto.ArchiveInfoRequest{FileHeader: file})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something unexpected occured"})
	}

	c.JSON(http.StatusOK, res)

}
