package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClosed    bool
	Rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {

	fmt.Println("Валедирую Имя!")
	if name == "" {
		fmt.Println("Имя не прошло валидацию!")
		return User{}
	}
	fmt.Println("Валедирую Возраст!")
	if age <= 0 || age >= 150 {
		fmt.Println("Возраст не прошол валидацию!")
		return User{}
	}
	fmt.Println("Валедирую Номер!")
	if phoneNumber == "" {
		fmt.Println("Номер не прошол валидацию!")
		return User{}
	}
	fmt.Println("Валедирую Рейтинг!")
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Рейтинг не прошол валидацию!")
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClosed:    isClose,
		Rating:      rating,
	}

}

func (u *User) ChangeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) ChangeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) ChangePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAcccaunt() {
	u.IsClosed = true
}

func (u *User) OpenAccaunt() {
	u.IsClosed = false
}

func (u *User) RatingUP(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
	}
}

func (u *User) RatingDOWN(rating float64) {
	if u.Rating-rating >= 0.0 {
		u.Rating -= rating
	}
}

func (u User) Greeting() {
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут", u.Name)
	fmt.Println("Мой рейтинг", u.Rating)
	fmt.Println("")

}

func (u User) Goodby() {
	fmt.Println("Всем всем пока!")
	fmt.Println("Меня звали", u.Name)
	fmt.Println("Мой рейтинг был", u.Rating)
	fmt.Println("")

}

func (u *User) RaingUP(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("Я добавил рейтинг")
	} else {
		fmt.Println("Я не прошел валидацию")

	}
}

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
