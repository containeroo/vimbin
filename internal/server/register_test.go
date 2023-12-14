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

		Register(path, description, false, handlerFunc)

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

		Register(path, description, false, handlerFunc, methods...)

		// Check if the handler is registered correctly
		assert.Len(t, Handlers, 1)
		assert.Equal(t, path, Handlers[0].Path)
		assert.Equal(t, description, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Equal(t, methods, Handlers[0].Methods)
	})

	t.Run("Register a handler token set to true", func(t *testing.T) {
		// Clear existing handlers
		Handlers = nil

		path := "/example"
		description := "Example handler description"
		handlerFunc := func(http.ResponseWriter, *http.Request) {}

		Register(path, description, true, handlerFunc)

		// Check if the handler is registered correctly
		assert.Len(t, Handlers, 1)
		assert.Equal(t, path, Handlers[0].Path)
		assert.Equal(t, description, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Empty(t, Handlers[0].Methods)
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

		path3 := "/example3"
		description3 := "Example handler 3 description"
		handlerFunc3 := func(http.ResponseWriter, *http.Request) {}
		methods3 := []string{"GET", "POST"}

		Register(path1, description1, false, handlerFunc1, methods1...)
		Register(path2, description2, false, handlerFunc2, methods2...)
		Register(path3, description3, true, handlerFunc3, methods3...)

		// Check if both handlers are registered correctly
		assert.Len(t, Handlers, 3)

		assert.Equal(t, path1, Handlers[0].Path)
		assert.Equal(t, description1, Handlers[0].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc1).Pointer(), reflect.ValueOf(Handlers[0].Handler).Pointer())
		assert.Equal(t, methods1, Handlers[0].Methods)

		assert.Equal(t, path2, Handlers[1].Path)
		assert.Equal(t, description2, Handlers[1].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc2).Pointer(), reflect.ValueOf(Handlers[1].Handler).Pointer())
		assert.Equal(t, methods2, Handlers[1].Methods)

		assert.Equal(t, path3, Handlers[2].Path)
		assert.Equal(t, description3, Handlers[2].Description)
		assert.Equal(t, reflect.ValueOf(handlerFunc3).Pointer(), reflect.ValueOf(Handlers[2].Handler).Pointer())
		assert.Equal(t, methods3, Handlers[2].Methods)

	})
}
