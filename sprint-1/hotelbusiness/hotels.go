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
	if len(guests) == 0 {
		return []Load{}
	}

	changes := make(map[int]int)
	for _, v := range guests {
		changes[v.CheckInDate]++
		changes[v.CheckOutDate]--
	}

	var days []int
	for day := range changes {
		days = append(days, day)
	}
	sort.Ints(days)

	var result []Load
	count := 0
	prevValue := -1
	for _, day := range days {
		count += changes[day]

		if count != prevValue {
			result = append(result, Load{day, count})
			prevValue = count
		}
	}

	return result
}
