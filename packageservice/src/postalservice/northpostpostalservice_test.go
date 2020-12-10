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

	assert.Equal(t, result, "failure")
	assert.Equal(t, err, errors.New("Must specify origin and destination"))
}

func TestNorthPolePostServiceSendDestinationFailure(t *testing.T) {
	northPolePost := NorthPolePostService{
		Location: "Santa's Workhouse",
	}

	origin := "Toronto, Canada"
	destination := "San Jose, USA"

	result, err := northPolePost.Send(origin, destination)

	assert.Equal(t, result, "failure")
	assert.Equal(t, err, errors.New("We only send letters to the north pole!"))
}
