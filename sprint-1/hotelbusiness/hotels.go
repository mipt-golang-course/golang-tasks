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
	if len(guests) == 0 {
		return nil
	}

	startDate := guests[0].CheckInDate
	endDate := guests[0].CheckOutDate

	for _, val := range guests {
		if val.CheckInDate < startDate {
			startDate = val.CheckInDate
		}
		if val.CheckOutDate > endDate {
			endDate = val.CheckOutDate
		}
	}

	var load []Load
	dateCntPrev := 0
	for date := startDate; date <= endDate; date++ {
		dateCnt := 0
		for _, val := range guests {
			if val.CheckInDate <= date && date < val.CheckOutDate {
				dateCnt++
			}
		}
		if dateCnt != dateCntPrev {
			load = append(load, Load{date, dateCnt})
		}
		dateCntPrev = dateCnt
	}
	return load
}
