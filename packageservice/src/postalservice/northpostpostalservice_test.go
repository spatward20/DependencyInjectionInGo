package postalservice

import (
	"errors"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestNorthPolePostServiceSendSuccess(t *testing.T) {
	northPolePost := NorthPolePostService{
		Location: "Santa's Workhouse",
	}

	origin := "Toronto, Canada"
	destination := "North Point, North Pole"

	result, err := northPolePost.Send(origin, destination)

	assert.Equal(t, result, "success")
	assert.Equal(t, err, nil)
}

func TestNorthPolePostServiceSendFailure(t *testing.T) {
	northPolePost := NorthPolePostService{
		Location: "Santa's Workhouse",
	}

	origin := ""
	destination := "North Pole"

	result, err := northPolePost.Send(origin, destination)

	assert.Equal(t, result, "")
	assert.Equal(t, err, errors.New("Invalid parameters"))
}
