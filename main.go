package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Ввидите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}

		cmd := fields[0]

		if cmd == "выйти" {
			fmt.Println("До скорого!")
			return
		}

		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Вы хотите добавить: ", str)
		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Вы ажется хотите удалить: ", str)
		} else if cmd == "help" {
			fmt.Println("Команда: help")
			fmt.Println("-- эта команда выводит список доступных команд")
			fmt.Println("")
			fmt.Println("Команда: добавить {что нужно добавить}")
			fmt.Println("-- эта команда позволяет добавлять что-то")
			fmt.Println("")
			fmt.Println("Команда: удалить {что нужно удалить}")
			fmt.Println("-- эта команда позволяет удалять что-то")
			fmt.Println("")
		} else {
			fmt.Println("вы вели неизвестную команду:")
		}
	}
}
