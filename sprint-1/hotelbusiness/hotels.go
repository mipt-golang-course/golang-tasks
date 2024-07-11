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

	load := make([]Load,0)

	var previous = -1

	for iter:=0; iter < len(result) + 1; iter++{
		var test = Load{iter,result[iter]}
		if result[iter] != previous{
			load = append(load,test)
		}
		previous = result[iter]
		i++
	}

	load = load [1:]

	return load
}