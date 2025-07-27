package helpers

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
)

func CreateArchive(files []string, archiveName string) ([]string, string) {
	downloadedFiles := make(map[string]string, len(files))
	failedFiles := []string{}
	
	// Формирование временных файлов для оригинальных
	for _, url := range files {
		fileID, _ := uuid.NewRandom()
		downloadedFiles[url] = fileID.String() + path.Ext(url)
	}
	
	// Создание архива
	zipName := archiveName + ".zip"
	zipFile, _ := os.Create(zipName)
	defer zipFile.Close()
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Запись файлов в архив
	for url, tempFile := range downloadedFiles {
		resp, err := http.Get(url)
		if err != nil {
			failedFiles = append(failedFiles, url)
			continue
		}
		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			failedFiles = append(failedFiles, url)
			continue
		}

		header := &zip.FileHeader{
			Name: tempFile,
			Method: zip.Deflate,
			Modified: time.Now(),
		}

		writer, _ := zipWriter.CreateHeader(header)
		_, err = writer.Write(content)
	}

	return failedFiles, zipName
}
