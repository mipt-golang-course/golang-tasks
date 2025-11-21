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
	days := getDays(loadByDays)

	load := []Load{}

	prevLoadDay := 0
	for _, day := range days {
		loadByDay := loadByDays[day] + prevLoadDay

		if loadByDay != prevLoadDay {
			load = append(load, Load{
				StartDate:  day,
				GuestCount: loadByDay,
			})
			prevLoadDay = loadByDay
		}

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

func getDays(loadByDays map[int]int) []int {
	days := make([]int, 0, len(loadByDays))
	for day := range loadByDays {
		days = append(days, day)
	}
	slices.Sort(days)
	return days
}
