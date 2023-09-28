package src

import (
	"fmt"
	"log/slog"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Filename     string
	Size         int64
	Header       string
	Path         string
	OriginalName string
}

func (f *File) Upload(file multipart.File) error {
	f.Path = "uploads/"

	newFileName := time.Now().UnixNano()
	dst, err := os.Create(fmt.Sprintf("uploads/%d%s", newFileName, filepath.Ext(f.OriginalName)))
	if err != nil {
		slog.Error("Error Creating File",
			"error", err,
		)
	}

	defer dst.Close()
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
