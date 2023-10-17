package archive

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/iliyasali2107/archiver/internal/dto"
	"github.com/iliyasali2107/archiver/internal/helpers/mimetype"
	"github.com/iliyasali2107/archiver/internal/models"
)

type ArchiveService struct{}

func NewArchiveSvc() *ArchiveService {
	return &ArchiveService{}
}

// size is 512, because for mimetype check we need only first 512 bytes.
const buffSize = 512

var legalMIMETypes = []string{
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/xml",
	"image/jpeg",
	"image/png",
}

var ErrNotAllowedMIMEType = errors.New("not allowed MIMEType is used")

func (as *ArchiveService) GetArchiveInfo(req dto.ArchiveInfoRequest) (dto.ArchiveInfoResponse, error) {
	var res dto.ArchiveInfoResponse

	fileHeader := req.FileHeader
	mpFile, err := fileHeader.Open()
	if err != nil {
		return dto.ArchiveInfoResponse{}, nil
	}
	defer mpFile.Close()

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

func (as *ArchiveService) Compress(req dto.ArchiveCompressRequest) (dto.ArchiveCompressResponse, error) {
	files := req.Files

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return dto.ArchiveCompressResponse{}, err
		}
		defer file.Close()

		mimeCheckerBuffer := make([]byte, buffSize)
		_, err = file.Read(mimeCheckerBuffer)
		if err != nil {
			return dto.ArchiveCompressResponse{}, err
		}

		mimeType := http.DetectContentType(mimeCheckerBuffer)
		if !mimetype.Contains(legalMIMETypes, mimeType) {
			return dto.ArchiveCompressResponse{}, ErrNotAllowedMIMEType
		}

		fh := &zip.FileHeader{Name: fileHeader.Filename, Flags: 0x800, Method: zip.Deflate}
		destWriter, err := zipWriter.CreateHeader(fh)
		if err != nil {
			return dto.ArchiveCompressResponse{}, err
		}

		_, err = io.Copy(destWriter, file)
		if err != nil {
			return dto.ArchiveCompressResponse{}, err
		}

		zipWriter.Close()
	}

	return dto.ArchiveCompressResponse{
		Buffer: buf,
	}, nil
}

func (as *ArchiveService) Clear(filename string) error {
	return os.Remove(filename)
}
