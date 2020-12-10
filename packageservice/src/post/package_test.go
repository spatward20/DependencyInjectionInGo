package post

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSendPackageSuccess(t *testing.T) {

	p := Package{
		Destination: "Canada",
		Origin:      "Canada",
	}

	result, err := p.PostPackage()

	assert.Equal(t, err, nil)
	assert.Equal(t, result, "success")
}
