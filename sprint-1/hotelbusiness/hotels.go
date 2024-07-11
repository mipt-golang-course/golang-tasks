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
		
		if result[entry.CheckOutDate] != 0{

		} else {
			result[entry.CheckOutDate] = 0
		}
	}

	var i = 0

	load := make([]Load, 0)

	/*for k, _:= range result{
		if k !=0{
			if result[k] == result[k-1]{
				delete(result, k)
			}
		}
	}*/

	var previous = -1

	for k, v := range result{
		var test = Load{k,v}
		//fmt.Print(k,v,previous)
		if v != previous{
			load = append(load,test)
		}
		previous = v
		i++
	}

	return load
}