package main

import "github.com/k0kubun/pp"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userAray := []User{
		User{
			Name:    "Вася",
			Rating:  5.5,
			Premium: true,
		},
		User{
			Name:    "Вася",
			Rating:  4.5,
			Premium: false,
		},
		User{
			Name:    "Вася",
			Rating:  7.5,
			Premium: true,
		},
	}

	userArray := append(
		userAray,
		User{
			Name:    "Виталик",
			Rating:  4.0,
			Premium: true,
		},
	)

	pp.Println("len:", len(userArray))
	pp.Println("cap:", cap(userArray))

	intSlice := make([]int, 0, 10)
}

// type Slice struct {
// 	len int,
// 	cap int,
// 	ptr *[]type,
// }
