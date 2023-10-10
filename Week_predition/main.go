package main

import (
	"errors"
	"fmt"
)

func main() {
	var a string
	fmt.Println("Какой сегодня день?")
	fmt.Scan(&a)

	day, err := prediction(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(day)
}

func prediction(day string) (string, error) {
	switch day {
	case "понедельник":
		return "C началом рабочей недели! Сил вам, ведь впереди много дней без выходных", nil
	case "вторник":
		return "2ой рабочий день, понимаю вас, вам очень тяжело", nil
	case "среда":
		return "3ий день страданий, вы ещё живы?)", nil
	case "четверг":
		return "4ый день, конец рабочей недели не за горами)", nil
	case "пятница":
		return "Уже пятница?) Последний рывок перед выходными. Вечером уже можете чиллить хоть до рассвета.", nil
	case "суббота":
		return "Ура, чилл день. Приятного вам отдыха!)", nil
	case "воскресенье":
		return "Последний выходной, отдохните от души, а то завтра уже на работу.", nil
	default:
		return "Упс...., что то пошло не так! Невалидный день недели.", errors.New("invalid day of the week")
	}

}
