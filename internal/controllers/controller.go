package controllers

import (
	"github.com/iliyasali2107/archiver/internal/controllers/archive_compress"
	"github.com/iliyasali2107/archiver/internal/controllers/archive_info"
	"github.com/iliyasali2107/archiver/internal/controllers/mail_sender"
)

type Controller struct {
	ArchiveInfoCtrl    *archive_info.ArchiveInfoCtrl
	ArchiveCompresCtrl *archive_compress.ArchiveCompressCtrl
	MailSenderCtrl     *mail_sender.MailSenderCtrl
}

func NewController(aiSvc archive_info.ArchiveInfoSvc, acSvc archive_compress.ArchiveCompressSvc, msSvc mail_sender.MailSenderdSvc) *Controller {
	return &Controller{
		ArchiveInfoCtrl:    archive_info.NewArchiveInfoCtrl(aiSvc),
		ArchiveCompresCtrl: archive_compress.NewArchiveCompressCtrl(acSvc),
		MailSenderCtrl:     mail_sender.NewMailSenderCtrl(msSvc),
	}
}
