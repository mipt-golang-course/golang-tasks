//go:build !solution

package hotelbusiness

import (
"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest, ) []Load {

	var load []Load

	dates := make(map[int]int)
	var keys []int

	for _,g := range guests {

		if g == (Guest{}) {
			continue
		}
		dates[g.CheckInDate]++
		dates[g.CheckOutDate]--
	}


	for k := range dates {
		keys = append(keys,k)
	}

	sort.Ints(keys)

	var g int
	for _,k := range keys {

		gOld := g
		var l Load
		l.StartDate = k

		g = g + dates[k]
		if g == gOld {
			continue
		}

		l.GuestCount = g

		load = append(load, l)
	}

	return load
}
