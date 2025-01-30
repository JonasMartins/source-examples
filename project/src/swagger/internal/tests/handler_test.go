package tests

import (
	"log"
	"net/http"
	"project/pkg/utils"
	"project/src/swagger/internal/factories"
	"testing"
)

func TestSwagger(t *testing.T) {
	description := "Test if the swagger handler is working correctly"
	defer func() {
		log.Printf("Test: %s\n", description)
		log.Println("Deferred tearing down.")
	}()

	mux := http.NewServeMux()
	srv := factories.BuildSrv(mux)

	t.Run("Should call the health endpoint correctely", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/health", nil)
		if err != nil {
			t.Errorf("%v", err)
		}
		response := utils.ExecuteRequest(req, srv.Handler)
		utils.CheckResponseCode(t, http.StatusOK, response.Code)
	})
}
