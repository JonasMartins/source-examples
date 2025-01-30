package main

import (
	"fmt"
	"net/http"
	"project/pkg/utils"
	"project/src/swagger/internal/factories"
)

func Run() {
	mux := http.NewServeMux()
	srv := factories.BuildSrv(mux)

	fmt.Printf("Server started at :%s\n", utils.DEFAULT_PORT_STR)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
