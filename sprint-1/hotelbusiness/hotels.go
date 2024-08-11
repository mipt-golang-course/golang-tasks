package hotelbusiness

import (
	"sort"
)

// Guest представляет данные о заезде и выезде гостя.
type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

// Load представляет данные о загрузке курорта на конкретную дату.
type Load struct {
	StartDate  int
	GuestCount int
}

// ComputeLoad вычисляет загрузку курорта по дням на основе данных о заездах и выездах гостей.
func ComputeLoad(guests []Guest) []Load {
	// Создаем карту для подсчета изменений количества гостей на конкретные даты
	changes := make(map[int]int)

	// Заполняем карту: увеличиваем количество гостей на дату заезда и уменьшаем на дату выезда
	for _, guest := range guests {
		changes[guest.CheckInDate]++
		changes[guest.CheckOutDate]--
	}

	// Извлекаем и сортируем все уникальные даты
	dates := make([]int, 0, len(changes))
	for date := range changes {
		dates = append(dates, date)
	}
	sort.Ints(dates)

	// Идем по отсортированным датам, подсчитывая текущую загрузку курорта
	var loads []Load
	previousGuests := 0
	currentGuests := 0
	for _, date := range dates {
		currentGuests += changes[date]
		if currentGuests != previousGuests { // Добавляем запись о загрузке только если произошли изменения
			loads = append(loads, Load{
				StartDate:  date,
				GuestCount: currentGuests,
			})
		}
		previousGuests = currentGuests
	}

	return loads
}
