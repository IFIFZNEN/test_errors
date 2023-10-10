package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var num int
	fmt.Print("Напишите ваш возраст: ")
	fmt.Scanln(&num)
	info, err := YoungOrOld(num)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print(info)
}

func YoungOrOld(age int) (string, error) {
	if age > 0 && age <= 12 {
		return "Вы ещё ребёнок", nil
	} else if age > 12 && age <= 21 {
		return "Вы подросток", nil
	} else if age > 21 && age <= 45 {
		return "Вы зрелый", nil
	} else if age > 45 && age <= 90 {
		return "Вы пожилой индивид", nil
	} else if age > 90 && age <= 120 {
		return "Вы долгожитель!", nil
	} else if age > 120 {
		return "ВЫ точно человек?", errors.New("are you human?")
	}
	return "Вы неверно написали свой возраст", errors.New("you misspelled your age")
}
