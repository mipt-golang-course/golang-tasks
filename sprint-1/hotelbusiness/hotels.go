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

func ComputeLoad(guests []Guest) []Load {
	type matrix map[int]int
	var result []Load
	var guestCount matrix
	guestCount = make(matrix)

	for _, gst := range guests {
		for i := gst.CheckInDate; i <= gst.CheckOutDate; i++ {
			guestCount[i]++
		}
		guestCount[gst.CheckOutDate]--
	}

	keys := make([]int, 0)
	for k, _ := range guestCount {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	currentState := 0
	for _, k := range keys {
		if guestCount[k] != currentState {
			result = append(result, Load{k, guestCount[k]})
			currentState = guestCount[k]
		}
	}

	return result
}
