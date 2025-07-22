package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	// weather := map[int]int{
	// 	11: +3,
	// 	12: +6,
	// 	13: +9,
	// 	14: -4,
	// 	15: +1,
	// }

	// for key, _ := range weather {
	// 	weather[key] += 1
	// }

	// c, ok := weather[30]
	// if ok {
	// 	pp.Println("OKEY")
	// } else {
	// 	pp.Println("bad")
	// }
	// pp.Println(c)
	// pp.Println(weather)

	criminal := map[string]bool{
		"Вася":    true,
		"Петя":    false,
		"Антон":   false,
		"Вова":    false,
		"Алексей": true,
	}

	c, ok := criminal["Вася"]

	if !ok {
		pp.Println("Человека нет в базе")
		return
	}
	pp.Println("Человека найден в базе")
	if c {
		pp.Println("Человек судим")
	} else {
		pp.Println("Человек не судим")
	}
	// pp.Println(c, ok)
}

// type Slice struct {
// 	len int,
// 	cap int,
// 	ptr *[]type,
// }
