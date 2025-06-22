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

	start_date := guests[0].CheckInDate
	end_date := guests[0].CheckOutDate
	for _, val := range guests {
		if val.CheckInDate < start_date {
			start_date = val.CheckInDate
		}
		if val.CheckOutDate > end_date {
			end_date = val.CheckOutDate
		}
	}

	var load []Load
	date_cnt_prev := 0
	for date := start_date; date <= end_date; date++ {
		date_cnt := 0
		for _, val := range guests {
			if val.CheckInDate <= date && date < val.CheckOutDate {
				date_cnt++
			}
		}
		if date_cnt != date_cnt_prev {
			load = append(load, Load{date, date_cnt})
		}
		date_cnt_prev = date_cnt
	}
	return load
}
