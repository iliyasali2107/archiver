package dto

import (
	"bytes"
	"mime/multipart"

	"github.com/iliyasali2107/archiver/internal/models"
)

type ArchiveInfoRequest struct {
	FileHeader *multipart.FileHeader
}

type ArchiveInfoResponse struct {
	FileName    string        `json:"filename"`
	ArchiveSize float64       `json:"archive_size"`
	TotalSize   float64       `json:"total_size"`
	TotalFiles  float64       `json:"total_files"`
	Files       []models.File `json:"files"`
}

type ArchiveCompressRequest struct {
	Files []*multipart.FileHeader
}

type ArchiveCompressResponse struct {
	Buffer *bytes.Buffer
}

type SendMailRequest struct {
	ReceiverEmails []string
	FileHeader     *multipart.FileHeader
}

type SendMailResponse struct{}
