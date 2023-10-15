package controllers

import (
	"github.com/iliyasali2107/archiver/internal/controllers/archive_info"
)

type Controller struct {
	ArchiveInfoCtrl *archive_info.ArchiveInfoCtrl
}

func NewController(svc archive_info.ArchiveInfoSvc) *Controller {
	return &Controller{
		ArchiveInfoCtrl: archive_info.NewArchiveInfoCtrl(svc),
	}
}
