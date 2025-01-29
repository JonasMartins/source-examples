package main

import (
	"fmt"
	"log"
	"os"
	"project/pkg/utils"
	"project/src/files_upload/internal/factories"
)

func Run() {
	preareUploadDir()
	srv := factories.BuildSrv()
	// Start the server
	fmt.Printf("Server started at :%s\n", utils.DEFAULT_PORT_STR)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}

}

func preareUploadDir() {
	uploadDir, err := utils.GetFilePath(&[]string{"src", "files_upload", "uploads"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = os.MkdirAll(*uploadDir, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
	}

}
