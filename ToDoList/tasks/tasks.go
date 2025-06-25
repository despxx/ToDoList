package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

// Структура задачи
type Task struct {
	Name     string
	Text     string
	DateAdd  time.Time // время получения задачи
	IsDone   bool
	DateDone time.Time
}

// Конструктор новой задачи с валидацией
func NewTask(
	name string,
	text string,
	dateAdd time.Time,
	isDone bool,
	dateDone time.Time,
) (*Task, error) {
	// Проверки на ошибки ввода
	// Проверка названия
	nameSlice := strings.Fields(name)
	if len(nameSlice) != 1 {
		return nil, pp.Errorf("Название должно быть из одного слова")
	}

	// isDoneStr := ReadString(isDone)

	return &Task{
		Name:     strings.TrimSpace(name),
		Text:     strings.TrimSpace(text),
		DateAdd:  time.Now(),
		IsDone:   isDone,
		DateDone: dateDone,
	}, nil
}

// Чтение строки с пробелами
func ReadString(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}
