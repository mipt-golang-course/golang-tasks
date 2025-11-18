//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int // 1 - cout
	CheckOutDate int // 1
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	loadMap := getLoadMap(&guests)
	sortedKeysLoadMap := getSortedKeysLoadMap(&loadMap)

	load := []Load{}
	for _, day := range sortedKeysLoadMap {
		load = append(load, Load{
			StartDate:  day,
			GuestCount: loadMap[day],
		})
	}
	return load
}

func getSortedKeysLoadMap(loadMap *map[int]int) []int {
	keys := make([]int, 0, len(*loadMap))
	for k := range *loadMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func getLoadMap(guests *[]Guest) map[int]int {

	loadCheckIn := getLoadCheckIn(guests)
	loadCheckOut := getLoadCheckOut(guests)
	days := getDays(guests)

	loadMap := make(map[int]int)

	prevLoadDay := 0
	for _, day := range days {
		checkInByDay := loadCheckIn[day]
		checkOutByDay := loadCheckOut[day]
		loadByDay := checkInByDay - checkOutByDay + prevLoadDay

		if loadByDay != prevLoadDay {
			loadMap[day] = loadByDay
		}

		prevLoadDay = loadByDay

	}
	return loadMap
}

func getLoadCheckIn(guests *[]Guest) map[int]int {
	loadMap := make(map[int]int)
	for _, guest := range *guests {
		loadMap[guest.CheckInDate] = loadMap[guest.CheckInDate] + 1
	}
	return loadMap
}

func getLoadCheckOut(guests *[]Guest) map[int]int {
	loadMap := make(map[int]int)
	for _, guest := range *guests {
		loadMap[guest.CheckOutDate] = loadMap[guest.CheckOutDate] + 1
	}
	return loadMap
}

func getDays(guests *[]Guest) []int {
	setDays := make(map[int]bool)
	for _, guest := range *guests {
		setDays[guest.CheckInDate] = true
		setDays[guest.CheckOutDate] = true
	}
	days := make([]int, 0, len(setDays))

	for day := range setDays {
		days = append(days, day)
	}
	sort.Ints(days)
	return days
}
