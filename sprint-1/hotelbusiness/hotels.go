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

	type MyMap map[int] int

	var result = make(MyMap)

	for _, entry:= range guests{
		//fmt.Print(entry.CheckInDate)
		//fmt.Print(entry.CheckOutDate)
		for iter:= entry.CheckInDate; iter < entry.CheckOutDate ;{
			result[iter]++
			iter++
		}
	}

	var i = 0

	load := make([]Load, 0)

	for k, v := range result{
		var test = Load{k,v}
		load = append(load,test)
		i++
	}

	return load
}