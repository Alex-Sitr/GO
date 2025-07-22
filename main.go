package main

import "fmt"

func main() {
	user := NewUser(
		"Вася",
		500,
		"79118707090",
		false,
		6.0,
	)

	fmt.Println("User:", user)
	// user.RatingUP(2.5)
	// fmt.Println("Рейтинг после:", user.Rating)

}
