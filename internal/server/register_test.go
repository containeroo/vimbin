package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("Register a handler without specified methods", func(t *testing.T) {
		// Clear existing handlers
		Handlers = nil

		path := "/example"
		description := "Example handler description"
		handlerFunc := func(http.ResponseWriter, *http.Request) {}

		Register(path, handlerFunc, description)

		// Check if the handler is registered correctly
		assert.Len(t, Handlers, 1)
		assert.Equal(t, path, Handlers[0].Path)
		assert.Equal(t, description, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Empty(t, Handlers[0].Methods)
	})

	t.Run("Register a handler with specified methods", func(t *testing.T) {
		// Clear existing handlers
		Handlers = nil

		path := "/example"
		description := "Example handler description"
		handlerFunc := func(http.ResponseWriter, *http.Request) {}
		methods := []string{"GET", "POST"}

		Register(path, handlerFunc, description, methods...)

		// Check if the handler is registered correctly
		assert.Len(t, Handlers, 1)
		assert.Equal(t, path, Handlers[0].Path)
		assert.Equal(t, description, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Equal(t, methods, Handlers[0].Methods)
	})

	t.Run("Register multiple handlers", func(t *testing.T) {
		// Clear existing handlers
		Handlers = nil

		path1 := "/example1"
		description1 := "Example handler 1 description"
		handlerFunc1 := func(http.ResponseWriter, *http.Request) {}
		methods1 := []string{"GET"}

		path2 := "/example2"
		description2 := "Example handler 2 description"
		handlerFunc2 := func(http.ResponseWriter, *http.Request) {}
		methods2 := []string{"POST"}

		Register(path1, handlerFunc1, description1, methods1...)
		Register(path2, handlerFunc2, description2, methods2...)

		// Check if both handlers are registered correctly
		assert.Len(t, Handlers, 2)

		assert.Equal(t, path1, Handlers[0].Path)
		assert.Equal(t, description1, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc1).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Equal(t, methods1, Handlers[0].Methods)

		assert.Equal(t, path2, Handlers[1].Path)
		assert.Equal(t, description2, Handlers[1].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc2).Pointer(), reflect.ValueOf(Handlers[1].Handler).Pointer())
		assert.Equal(t, methods2, Handlers[1].Methods)
	})
}
