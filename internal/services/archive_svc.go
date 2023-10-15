package services

import (
	"archive/zip"

	"github.com/iliyasali2107/archiver/internal/dto"
	"github.com/iliyasali2107/archiver/internal/helpers/mimetype"
	"github.com/iliyasali2107/archiver/internal/models"
)

type ArchiveService struct{}

// size is 512, because for mimetype check we need only first 512 bytes.
const buffSize = 512

func (as *ArchiveService) GetArchiveInfo(req dto.ArchiveInfoRequest) (dto.ArchiveInfoResponse, error) {
	var res dto.ArchiveInfoResponse

	fileHeader := req.FileHeader
	mpFile, err := fileHeader.Open()
	defer mpFile.Close()

	if err != nil {
		return dto.ArchiveInfoResponse{}, nil
	}

	zipReader, err := zip.NewReader(mpFile, fileHeader.Size)
	if err != nil {
		return dto.ArchiveInfoResponse{}, nil
	}

	var files []models.File
	var filesCount float64
	var totalSize float64

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue
		}

		file := models.File{}
		file.FilePath = f.Name
		file.Size = float64(f.FileInfo().Size())
		rc, err := f.Open()
		if err != nil {
			return dto.ArchiveInfoResponse{}, nil
		}

		buffer := make([]byte, buffSize)
		rc.Read(buffer)
		file.Mimetype = mimetype.GetMIMEtype(buffer)

		files = append(files, file)
		filesCount++
		totalSize += float64(file.Size)

	}

	res.FileName = fileHeader.Filename
	res.ArchiveSize = float64(fileHeader.Size)
	res.TotalFiles = filesCount
	res.TotalSize = totalSize
	res.Files = files

	return res, nil

}
