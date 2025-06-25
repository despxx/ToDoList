package main

import (
	"fmt"
	"time"
	"todolist/tasks"
	"todolist/texts"

	"github.com/k0kubun/pp"
)

// Меню команд!
func Menu() {

	// Слайсы всех данных
	nameArr := make([]string, 0)        // Названия задач
	textArr := make([]string, 0)        // Текст задач
	dateAddArr := make([]time.Time, 0)  // Даты создания
	isDoneArr := make([]bool, 0)        // Выполнена / не выполнена
	dateDoneArr := make([]time.Time, 0) // Даты выполнения
	eventArr := make([]string, 0)       // Все события

	for {
		pp.Print("Введите команду (Доступные команды можно узнать, введя команду help): ")

		answer := tasks.ReadString("")

		switch answer {
		case "help":
			pp.Println("=== Доступные команды ===")
			pp.Println("add - Добавить новую задачу")
			pp.Println("list - Список всех задач")
			pp.Println("done - Отметить задачу как выполненную")
			pp.Println("del - Удалить задачу")
			pp.Println("events - Список всех событий")
			pp.Println("exit - Выйти из приложения")
			pp.Println("=========================")
			// Добавление события
			eventArr = append(eventArr, "Вызвана команда -help-")

		case "add":
			NewTaskAdd(&nameArr, &textArr, &dateAddArr, &isDoneArr, &dateDoneArr)
			// Добавление события
			nameAdd := nameArr[len(nameArr)-1]
			allAdd := fmt.Sprintf("Добавлена задача %s", nameAdd)
			eventArr = append(eventArr, allAdd)

		case "list":
			pp.Println("== Список задач ==")
			texts.PrintArraysElements(nameArr, textArr, dateAddArr, isDoneArr, dateDoneArr)
			// Добавление события
			eventArr = append(eventArr, "Вызвана команда -list-")

		case "done":
			var doneAnswer string
			pp.Println("Какую задачу пометить выполненной? ")
			fmt.Scan(&doneAnswer)

			// Получаем индекс элемента и меняем значение на выполнено
			var indexDone int
			for i, name := range nameArr {
				if doneAnswer == name {
					indexDone = i
					isDoneArr[indexDone] = true
					// Меняем время выполнения задачи
					dateDoneArr[indexDone] = time.Now()
					// Добавление события
					allDone := fmt.Sprintf("Задача %s помечена выполненной", doneAnswer)
					eventArr = append(eventArr, allDone)
				}
			}

		case "del":
			var delAnswer string
			pp.Println("Какую задачу удалить? ")
			fmt.Scan(&delAnswer)

			// Получаем индекс элемента и удаляем значение
			var indexDel int
			for i, name := range nameArr {
				if delAnswer == name {
					indexDel = i
					nameArr = append(nameArr[:i], nameArr[i+1:]...)
					textArr = append(textArr[:i], textArr[i+1:]...)
					dateAddArr = append(dateAddArr[:i], dateAddArr[i+1:]...)
					isDoneArr = append(isDoneArr[:i], isDoneArr[i+1:]...)
					dateDoneArr = append(dateDoneArr[:i], dateDoneArr[i+1:]...)
					// Добавление события
					allDel := fmt.Sprintf("Задача %s удалена", delAnswer)
					eventArr = append(eventArr, allDel)
				}
			}

			// Удаляем тело задачи и дату создания по индексу
			texts.RemoveByIndex(textArr, indexDel)
			texts.RemoveByIndexTime(dateAddArr, indexDel)
			texts.RemoveByIndexBool(isDoneArr, indexDel)
			texts.RemoveByIndexTime(dateDoneArr, indexDel)

		case "events":
			// Добавление события
			eventArr = append(eventArr, "Вызвана команда -events-")
			pp.Println(eventArr)

		case "exit":
			return

		default:
			pp.Println("Неизвестная команда. Попробуйте еще раз")
			// Добавление события
			eventArr = append(eventArr, "Вызвана неизвестная команда")
		}

	}
}

// Создание новой задачи
func NewTaskAdd(
	nameArr *[]string,
	textArr *[]string,
	dateAddArr *[]time.Time,
	isDoneArr *[]bool,
	dateDoneArr *[]time.Time,
) {

	pp.Println("=== Введите данные новой задачи ===")

	// Название задачи
	name := tasks.ReadString("Название: ")
	*nameArr = append(*nameArr, name) // Добавляем в слайс названий задач

	// Текст и дата создания задачи
	text := tasks.ReadString("Текст задачи: ")
	dateAdd := time.Now()
	*textArr = append(*textArr, text)          // Добавляем в слайс задач
	*dateAddArr = append(*dateAddArr, dateAdd) // Добавляем в слайс дат создания заметок

	// Выполнено / не выполнено
	var isDone bool
	*isDoneArr = append(*isDoneArr, isDone) // Добавляем в слайс выполнено / не выполнено

	// Дата выполнения
	var dateDone time.Time
	*dateDoneArr = append(*dateDoneArr, dateDone)

	task, err := tasks.NewTask(name, text, dateAdd, isDone, dateDone)
	if err != nil {
		fmt.Println("Ошибка создания задачи:", err)
		return
	}

	pp.Printf("\n=== Создана задача ===\n")
	pp.Printf("Название: %s\nТекст задачи: %s\nДата создания: %s\nВыполнено? %s\n",
		task.Name, task.Text, task.DateAdd, task.IsDone)
	pp.Println("")
}

func main() {
	Menu()
}
