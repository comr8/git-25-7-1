package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Print("Введите данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Вы ввели следующие данные: %s\n", n)
}
