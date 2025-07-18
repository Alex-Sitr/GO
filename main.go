package main

import (
	"fmt"
	"study/user"
)

func main() {
	u := user.NewUser("Данил", 50)
	fmt.Println("User Name:", u)
}
