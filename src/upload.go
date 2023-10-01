package src

import (
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type AllowFileType string

const (
	MP4  AllowFileType = "video/mp4"
	OGG  AllowFileType = "video/ogg"
	WEBM AllowFileType = "video/webm"
)

type File struct {
	Filename     string
	Size         int64
	Header       string
	Path         string
	OriginalName string
	Processed    bool `default:"false"`
	ProcessedDir string
}

func (f *File) Upload(file multipart.File) error {
	f.Path = "uploads/"
	newFileName := time.Now().UnixNano()
	f.Filename = fmt.Sprint(newFileName)
	f.Path = fmt.Sprintf("%s%d%s", f.Path, newFileName, filepath.Ext(f.OriginalName))
	dst, err := os.Create(fmt.Sprintf("uploads/%d%s", newFileName, filepath.Ext(f.OriginalName)))
	if err != nil {
		slog.Error("Error Creating File",
			"error", err,
		)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return nil
	}
	return nil
}

func ValidateDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		slog.Error("Error Creating Uploads Directory",
			"path", dir,
		)
		return err
	}
	return nil
}

func ValidateFileType(file multipart.File, w http.ResponseWriter) {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filetype := http.DetectContentType(buff)
	switch filetype {
	case string(MP4):
		break
	case string(OGG):
		break
	case string(WEBM):
		break
	default:
		http.Error(w, "File type not supported", http.StatusInternalServerError)
	}
}
