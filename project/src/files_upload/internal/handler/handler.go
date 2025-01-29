package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"project/pkg/utils"
	"strings"
)

type Handler struct{}

func (h *Handler) UploadSingleFile(w http.ResponseWriter, r *http.Request) {
	uploadDir, err := utils.GetFilePath(&[]string{"src", "files_upload", "uploads"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file extension
	if !strings.HasSuffix(header.Filename, ".txt") {
		http.Error(w, "Invalid file type, only .txt files are allowed", http.StatusBadRequest)
		return
	}

	// Save the file locally
	filePath := filepath.Join(*uploadDir, header.Filename)
	if err := saveFile(file, filePath); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
}

// uploadMultipleFiles handles multiple file uploads
func (h *Handler) UploadMultipleFiles(w http.ResponseWriter, r *http.Request) {
	uploadDir, err := utils.GetFilePath(&[]string{"src", "files_upload", "uploads"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	for _, header := range files {
		// Validate file extension
		if !strings.HasSuffix(header.Filename, ".txt") {
			http.Error(w, fmt.Sprintf("Invalid file type, only .txt files are allowed: %s", header.Filename), http.StatusBadRequest)
			return
		}

		// Open the file
		file, err := header.Open()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to open file: %s", header.Filename), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Save the file locally
		filePath := filepath.Join(*uploadDir, header.Filename)
		if err := saveFile(file, filePath); err != nil {
			http.Error(w, fmt.Sprintf("Failed to save file: %s", header.Filename), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "All files uploaded successfully")
}

// saveFile saves the uploaded file to the specified path
func saveFile(file io.Reader, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	return err
}
