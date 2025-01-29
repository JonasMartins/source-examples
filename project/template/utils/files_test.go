package utils

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {
	description := "Test if utils files are working as espected"
	defer func() {
		log.Printf("Test: %s\n", description)
		log.Println("Deferred tearing down.")
	}()

	t.Run("Get root path", func(t *testing.T) {
		path, err := GetRootDir()
		assert.Nil(t, err)
		assert.NotZero(t, len(*path))
	})

	t.Run("Get migrations path", func(t *testing.T) {
		path, err := GetFilePath(&[]string{"utils", "files.go"})
		assert.Nil(t, err)
		assert.NotZero(t, len(*path))
	})
}
