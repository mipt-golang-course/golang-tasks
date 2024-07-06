//go:build !solution

package hotelbusiness

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	days := make(map[int]int)

	maxDay := 0

	for _, guest := range guests {
		days[guest.CheckInDate] += 1
		days[guest.CheckOutDate] -= 1
		if guest.CheckInDate > maxDay {
			maxDay = guest.CheckInDate
		}

		if guest.CheckOutDate > maxDay {
			maxDay = guest.CheckOutDate
		}
	}

	loads := make([]Load, 0, maxDay)

	lastLoad := 0
	for day := 0; day <= maxDay; day++ {
		val, ok := days[day]
		if ok && val != 0 {
			lastLoad += val
			loads = append(loads, Load{day, lastLoad})
		}
	}

	return loads
}
