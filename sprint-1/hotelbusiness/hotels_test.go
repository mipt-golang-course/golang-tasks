package hotelbusiness_test

import (
	"testing"

	hotel "github.com/mipt-golang-course/golang-tasks/sprint-1/hotelbusiness"
	"github.com/stretchr/testify/require"
)

func TestComputeLoad_basic(t *testing.T) {
	for _, tc := range []struct {
		title  string
		guests []hotel.Guest
		result []hotel.Load
	}{
		{
			title:  "empty input",
			guests: []hotel.Guest{},
			result: []hotel.Load{},
		},
		{
			title:  "one hotel.Guest",
			guests: []hotel.Guest{{1, 2}},
			result: []hotel.Load{{1, 1}, {2, 0}},
		},
		{
			title:  "two guests, one-by-one, without any gaps",
			guests: []hotel.Guest{{1, 2}, {2, 3}},
			result: []hotel.Load{{1, 1}, {3, 0}},
		},
		{
			title:  "two guests, one-by-one, with a gap",
			guests: []hotel.Guest{{1, 2}, {3, 4}},
			result: []hotel.Load{{1, 1}, {2, 0}, {3, 1}, {4, 0}},
		},
		{
			title:  "two guests, together",
			guests: []hotel.Guest{{1, 2}, {1, 2}},
			result: []hotel.Load{{1, 2}, {2, 0}},
		},
		{
			title:  "overlapping",
			guests: []hotel.Guest{{1, 3}, {3, 5}, {2, 4}},
			result: []hotel.Load{{1, 1}, {2, 2}, {4, 1}, {5, 0}},
		},
		{
			title:  "stairs",
			guests: []hotel.Guest{{1, 6}, {2, 5}, {3, 4}},
			result: []hotel.Load{{1, 1}, {2, 2}, {3, 3}, {4, 2}, {5, 1}, {6, 0}},
		},
		{
			title:  "starting late",
			guests: []hotel.Guest{{3, 7}, {5, 7}},
			result: []hotel.Load{{3, 1}, {5, 2}, {7, 0}},
		},
		{
			title:  "unordered",
			guests: []hotel.Guest{{4, 7}, {2, 4}, {2, 3}},
			result: []hotel.Load{{2, 2}, {3, 1}, {7, 0}},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			if len(tc.result) == 0 {
				require.Empty(t, hotel.ComputeLoad(tc.guests))
			} else {
				require.Equal(t, tc.result, hotel.ComputeLoad(tc.guests))
			}
		})
	}
}

func TestComputeLoad_stress1(t *testing.T) {
	n := 1000000
	g := make([]hotel.Guest, 0, 1000000)
	for i := 0; i < n; i++ {
		g = append(g, hotel.Guest{1, 2})
	}
	l := hotel.ComputeLoad(g)
	require.Equal(t, []hotel.Load{{1, n}, {2, 0}}, l)
}

func TestComputeLoad_stress2(t *testing.T) {
	n := 1000000
	g := make([]hotel.Guest, 0, 1000000)
	for i := 0; i < n; i++ {
		g = append(g, hotel.Guest{i, i + 1})
	}
	l := hotel.ComputeLoad(g)
	require.Equal(t, []hotel.Load{{0, 1}, {n, 0}}, l)
}
