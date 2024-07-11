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
	var result []Load
	var sliceDate []int
	mapDate := make(map[int][]int)
	var cnt int
	var prevCnt int
	for _, v := range guests {
		if !slices.Contains(sliceDate, v.CheckInDate) {
			sliceDate = append(sliceDate, v.CheckInDate)
		}
		if !slices.Contains(sliceDate, v.CheckOutDate) {
			sliceDate = append(sliceDate, v.CheckOutDate)

		}
		mapDate[v.CheckInDate] = append(mapDate[v.CheckInDate], 1)
		mapDate[v.CheckOutDate] = append(mapDate[v.CheckOutDate], -1)
	}
	slices.Sort(sliceDate)

	for _, v := range sliceDate {
		for _, gcv := range mapDate[v] {
			if gcv > 0 {
				cnt++
			} else if gcv < 0 {
				cnt--
			}
		}
		if cnt != prevCnt {
			result = append(result, Load{StartDate: v, GuestCount: cnt})
		}
		prevCnt = cnt
	}
	return result
}
