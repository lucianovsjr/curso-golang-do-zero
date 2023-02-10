package main

import "fmt"

/*
	Utilizada quando você precisa tratar muitas condições.
*/

func dayOfWeek(day int) string {
	// Switch no Go não tem o break, ele já faz isso por padrão
	switch day {
	case 1:
		return "Mon"
	case 2:
		return "Tue"
	case 3:
		return "Wed"
	case 4:
		return "Thu"
	case 5:
		return "Fri"
	case 6:
		return "Sat"
	default:
		return "Invalid day"
	}
}

func dayOfWeek2(day int) string {
	var dayName string

	// Útil quando quer criar condições de variáveis diferentes
	switch {
	case day == 1:
		dayName = "Mon"
		fallthrough // Joga para a próxima condição
	case day == 2:
		dayName = "Tue"
	case day == 3:
		dayName = "Wed"
	case day == 4:
		dayName = "Thu"
	case day == 5:
		dayName = "Fri"
	case day == 6:
		dayName = "Sat"
	case day == 7:
		dayName = "Sun"
	case day > 7:
		fallthrough
	default:
		dayName = "Invalid day"
	}

	return dayName
}

func main() {
	day1 := dayOfWeek(1)
	fmt.Println(day1)

	day2 := dayOfWeek(10)
	fmt.Println(day2)

	day3 := dayOfWeek2(3)
	fmt.Println(day3)

	day4 := dayOfWeek2(1)
	fmt.Println(day4)

	day5 := dayOfWeek2(8)
	fmt.Println(day5)
}
