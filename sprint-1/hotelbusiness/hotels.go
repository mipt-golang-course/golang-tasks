//go:build !solution

package hotelbusiness

import (
	"cmp"
	"slices"
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
	// MaxFunc rejects eating empty list
	if len(guests) == 0 {
		return []Load{}
	}

	lastDate := slices.MaxFunc(guests, func(a, b Guest) int {
		return cmp.Compare(a.CheckOutDate, b.CheckOutDate)
	}).CheckOutDate

	dates := make([]int, lastDate+1)
	for _, guest := range guests {
		dates[guest.CheckInDate] += 1
		dates[guest.CheckOutDate] -= 1
	}

	var loading []Load
	nowLoaded := 0

	for day, guestsChange := range dates {
		if guestsChange != 0 {
			nowLoaded += guestsChange
			loading = append(loading, Load{StartDate: day, GuestCount: nowLoaded})
		}
	}

	return loading
}
