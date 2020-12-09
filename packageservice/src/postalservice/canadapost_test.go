package postalservice

import (
	"errors"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSendSuccess(t *testing.T) {
	canadaPost := CanadaPostService{
		Location: "Toronto",
	}

	origin := "Toronto, Canada"
	destination := "San Jose, California"

	result, err := canadaPost.Send(origin, destination)

	assert.Equal(t, result, "success")
	assert.Equal(t, err, nil)
}

func TestSendFailure(t *testing.T) {
	canadaPost := CanadaPostService{
		Location: "Toronto",
	}

	origin := ""
	destination := "San Jose, California"

	result, err := canadaPost.Send(origin, destination)

	assert.Equal(t, result, "")
	assert.Equal(t, err, errors.New("Invalid parameters"))
}
