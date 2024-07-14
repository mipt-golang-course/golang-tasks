package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	calendar := make(map[int]int)
	for _, guest := range guests {
		for day := guest.CheckInDate; day <= guest.CheckOutDate; day++ {
			if day == guest.CheckOutDate {
				calendar[day] = calendar[day]
			} else {
				calendar[day]++
			}
		}
	}

	load := []Load{}
	for day := range calendar {
		if day == 0 || calendar[day] != calendar[day-1] {
			load = append(load, Load{day, calendar[day]})
		}
	}

	sort.SliceStable(load, func(p, q int) bool {
		return load[p].StartDate < load[q].StartDate
	})

	return load
}
