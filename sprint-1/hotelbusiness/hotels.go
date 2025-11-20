//go:build !solution

package hotelbusiness

import (
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
	loadByDays := getLoadByDays(guests)
	days := getDays(guests)

	load := []Load{}

	prevLoadDay := 0
	for _, day := range days {
		loadByDay := loadByDays[day] + prevLoadDay

		if loadByDay != prevLoadDay {
			load = append(load, Load{
				StartDate:  day,
				GuestCount: loadByDay,
			})
		}
		prevLoadDay = loadByDay
	}
	return load
}

func getLoadByDays(guests []Guest) map[int]int {
	loadMap := make(map[int]int)
	for _, guest := range guests {
		loadMap[guest.CheckInDate]++
		loadMap[guest.CheckOutDate]--
	}
	return loadMap
}

func getDays(guests []Guest) []int {
	setDays := make(map[int]struct{})
	for _, guest := range guests {
		setDays[guest.CheckInDate] = struct{}{}
		setDays[guest.CheckOutDate] = struct{}{}
	}

	days := make([]int, 0, len(setDays))

	for day := range setDays {
		days = append(days, day)
	}
	slices.Sort(days)
	return days
}
