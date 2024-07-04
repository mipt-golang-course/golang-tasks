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
	loads := ComputeAllDaysLoad(guests)
	return ComputeResultLoad(loads)
}

func ComputeAllDaysLoad(guests []Guest) []Load {
	loads := make([]Load, 0)
	for _, guest := range guests {
		if guest.CheckOutDate >= len(loads) {
			for i := len(loads); i <= guest.CheckOutDate; i++ {
				loads = append(loads, Load{StartDate: i, GuestCount: 0})
			}
		}
		for i := guest.CheckInDate; i < guest.CheckOutDate; i++ {
			loads[i].GuestCount++
		}
	}
	return loads
}

func ComputeResultLoad(loads []Load) []Load {
	result := make([]Load, 0)
	currentLoad := 0
	for _, load := range loads {
		if load.GuestCount != currentLoad {
			result = append(result, load)
			currentLoad = load.GuestCount
		}
	}
	return result
}
