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

func ComputeLoad(guests []Guest) (res []Load) {
	if len(guests) == 0 {
		return nil
	}

	startDate := guests[0].CheckInDate
	endDate := guests[0].CheckOutDate
	changesInGuests := make(map[int]int)
	for _, guest := range guests {
		changesInGuests[guest.CheckInDate]++
		changesInGuests[guest.CheckOutDate]--

		startDate = min(startDate, guest.CheckInDate)
		endDate = max(endDate, guest.CheckOutDate)
	}

	currentGuests := 0
	for date := startDate; date <= endDate; date++ {
		if changesInGuests[date] != 0 {
			currentGuests += changesInGuests[date]
			res = append(res, Load{date, currentGuests})
		}
	}
	return res
}
