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

func TestCanadaPostSendFailure(t *testing.T) {
	canadaPost := CanadaPostService{
		Location: "Toronto",
	}

	origin := ""
	destination := "San Jose, California"

	result, err := canadaPost.Send(origin, destination)

	assert.Equal(t, result, "")
	assert.Equal(t, err, errors.New("Invalid parameters"))
}

func TestCanadaPostSendToNorthPoleFailure(t *testing.T) {
	canadaPost := CanadaPostService{
		Location: "Toronto",
	}

	origin := "Toronto"
	destination := "North Pole"

	result, err := canadaPost.Send(origin, destination)

	assert.Equal(t, result, "")
	assert.Equal(t, err, errors.New("Sorry, we don't ship to the North Pole"))
}
