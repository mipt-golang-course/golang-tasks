//go:build !solution

package hotelbusiness

import (
	"slices"
	"maps"
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
	if len(guests) == 0 {
		return []Load{}
	}

	// Записываем изменение гостей по дням
	delta := make(map[int]int)
	for _, guest := range guests {
		delta[guest.CheckInDate] += 1
		delta[guest.CheckOutDate] -= 1
	}

	// Создаём сортированный список дней
	dates := slices.Collect(maps.Keys(delta))

	slices.Sort(dates)

	curGuests := 0
	result := make([]Load, 0)
	for _, day := range dates {
		curGuests += delta[day]
		if (len(result) == 0) || (curGuests != result[len(result)-1].GuestCount) {
			result = append(result, Load{day, curGuests})
		}
	}

	return result
}
