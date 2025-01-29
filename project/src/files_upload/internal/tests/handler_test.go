package tests

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"project/pkg/utils"
	"project/src/files_upload/internal/factories"
	"testing"
)

func TestUpload(t *testing.T) {
	description := "Test if the update plan usecases are working correctly"
	defer func() {
		log.Printf("Test: %s\n", description)
		log.Println("Deferred tearing down.")
	}()
	srv := factories.BuildSrv()

	t.Run("Should check if upload single file is working", func(t *testing.T) {
		// Create a temporary .txt file
		fileContent := "This is a test file."
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", "testfile.txt")
		if err != nil {
			t.Fatal(err)
		}
		part.Write([]byte(fileContent))
		writer.Close()

		uploadDir, err := utils.GetFilePath(&[]string{"src", "files_upload", "uploads"})
		if err != nil {
			t.Errorf("%v", err)
		}

		// Create a request to the /upload/single endpoint
		req := httptest.NewRequest(http.MethodPost, "/upload/single", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Call the handler
		response := utils.ExecuteRequest(req, srv.Handler)
		utils.CheckResponseCode(t, http.StatusOK, response.Code)

		// Verify the file was saved
		filePath := filepath.Join(*uploadDir, "testfile.txt")
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("File was not saved: %s", filePath)
		}

		// Clean up
		os.Remove(filePath)
	})

	t.Run("Should check if upload multiple files is working", func(t *testing.T) {
		// Create a buffer and multipart writer
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		uploadDir, err := utils.GetFilePath(&[]string{"src", "files_upload", "uploads"})
		if err != nil {
			t.Errorf("%v", err)
		}

		// Add first file
		part1, err := writer.CreateFormFile("files", "file1.txt")
		if err != nil {
			t.Fatal(err)
		}
		part1.Write([]byte("First file content"))

		// Add second file
		part2, err := writer.CreateFormFile("files", "file2.txt")
		if err != nil {
			t.Fatal(err)
		}
		part2.Write([]byte("Second file content"))

		writer.Close()

		// Create a request to the /upload/multiple endpoint
		req := httptest.NewRequest(http.MethodPost, "/upload/multiple", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		response := utils.ExecuteRequest(req, srv.Handler)
		utils.CheckResponseCode(t, http.StatusOK, response.Code)

		// Verify the files were saved
		file1Path := filepath.Join(*uploadDir, "file1.txt")
		file2Path := filepath.Join(*uploadDir, "file2.txt")
		if _, err := os.Stat(file1Path); os.IsNotExist(err) {
			t.Errorf("File was not saved: %s", file1Path)
		}
		if _, err := os.Stat(file2Path); os.IsNotExist(err) {
			t.Errorf("File was not saved: %s", file2Path)
		}

		// Clean up
		os.Remove(file1Path)
		os.Remove(file2Path)
	})

}
