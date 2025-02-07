package model

import (
	"testing"
)

func TestQueryOrderNone(t *testing.T) {
	testCases := []struct {
		desc   string
		input  queryOrder
		output bool
	}{
		{
			desc:   "returns true if none enum value",
			input:  QueryOrderNone,
			output: true,
		},
		{
			desc:   "returns false if rating asc enum value",
			input:  QueryOrderRatingAsc,
			output: false,
		},
		{
			desc:   "returns false if rating desc enum value",
			input:  QueryOrderRatingDesc,
			output: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input.None() != tC.output {
				t.Errorf("expected %v, but got %v", tC.output, !tC.output)
			}
		})
	}
}

func TestQueryOrderLowestRating(t *testing.T) {
	testCases := []struct {
		desc   string
		input  queryOrder
		output bool
	}{
		{
			desc:   "returns false if none enum value",
			input:  QueryOrderNone,
			output: false,
		},
		{
			desc:   "returns true if rating asc enum value",
			input:  QueryOrderRatingAsc,
			output: true,
		},
		{
			desc:   "returns false if rating desc enum value",
			input:  QueryOrderRatingDesc,
			output: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input.LowestRating() != tC.output {
				t.Errorf("expected %v, but got %v", tC.output, !tC.output)
			}
		})
	}
}

func TestQueryOrderHighestRating(t *testing.T) {
	testCases := []struct {
		desc   string
		input  queryOrder
		output bool
	}{
		{
			desc:   "returns false if none enum value",
			input:  QueryOrderNone,
			output: false,
		},
		{
			desc:   "returns false if rating asc enum value",
			input:  QueryOrderRatingAsc,
			output: false,
		},
		{
			desc:   "returns true if rating desc enum value",
			input:  QueryOrderRatingDesc,
			output: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input.HighestRating() != tC.output {
				t.Errorf("expected %v, but got %v", tC.output, !tC.output)
			}
		})
	}
}

func TestQueryOptionsListOnly(t *testing.T) {
	testCases := []struct {
		desc   string
		input  QueryOptions
		output bool
	}{
		{
			desc:   "returns true if name is empty",
			input:  QueryOptions{Name: ""},
			output: true,
		},
		{
			desc:   "returns false if name is not empty",
			input:  QueryOptions{Name: "n"},
			output: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input.ListOnly() != tC.output {
				t.Errorf("expected %v, but got %v", tC.output, !tC.output)
			}
		})
	}
}
