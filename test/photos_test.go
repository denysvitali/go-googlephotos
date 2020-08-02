package test

import (
	photos "github.com/denysvitali/go-googlephotos"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPhotos(t *testing.T){
	s, err := photos.New(http.DefaultClient)
	assert.Nil(t, err)
	assert.NotNil(t, s)
}