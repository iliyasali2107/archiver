package controllers

import (
	"github.com/iliyasali2107/archiver/internal/controllers/archive_compress"
	"github.com/iliyasali2107/archiver/internal/controllers/archive_info"
)

type Controller struct {
	ArchiveInfoCtrl    *archive_info.ArchiveInfoCtrl
	ArchiveCompresCtrl *archive_compress.ArchiveCompressCtrl
}

func NewController(aiSvc archive_info.ArchiveInfoSvc, acSvc archive_compress.ArchiveCompressSvc) *Controller {
	return &Controller{
		ArchiveInfoCtrl:    archive_info.NewArchiveInfoCtrl(aiSvc),
		ArchiveCompresCtrl: archive_compress.NewArchiveCompressCtrl(acSvc),
	}
}
