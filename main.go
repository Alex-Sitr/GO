package main

import (
	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userAray := [3]User{
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

	pp.Println("до")
	pp.Println("----------")
	for index, user := range userAray {
		pp.Println(index, user)
	}
	pp.Println("")

	for index, user := range userAray {
		if user.Premium {
			userAray[index].Rating += 1
		}
	}
	pp.Println("после")
	pp.Println("----------")
	for i := 0; i < len(userAray); i++ {
		pp.Println(userAray[i])
	}
	pp.Println("")

}
