package texts

import (
	"fmt"
	"reflect"
	"time"
)

// Удаление элемента по значению
func RemoveByValue(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Удаление элемента string по индексу
func RemoveByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// Удаление элемента bool по индексу
func RemoveByIndexBool(slice []bool, index int) []bool {
	return append(slice[:index], slice[index+1:]...)
}

// Удаление элемента time по индексу
func RemoveByIndexTime(slice []time.Time, index int) []time.Time {
	return append(slice[:index], slice[index+1:]...)
}

// Печать элементов слайсов по одному
func PrintArraysElements(arr1, arr2, arr3, arr4, arr5 interface{}) {
	// Получаем reflect.Value для каждого массива
	v1 := reflect.ValueOf(arr1)
	v2 := reflect.ValueOf(arr2)
	v3 := reflect.ValueOf(arr3)
	v4 := reflect.ValueOf(arr4)
	v5 := reflect.ValueOf(arr5)

	// Проверяем, что длины массивов равны
	length := v1.Len()
	if v2.Len() != length || v3.Len() != length ||
		v4.Len() != length || v5.Len() != length {
		fmt.Println("Длины массивов должны быть одинаковыми")
		return
	}

	// Выводим элементы по одному
	for i := 0; i < length; i++ {
		fmt.Printf("%v", v1.Index(i).Interface())
		fmt.Println("")
		fmt.Printf("%v", v2.Index(i).Interface())
		fmt.Println("")
		fmt.Printf("Дата создания: %v", v3.Index(i).Interface())
		fmt.Println("")
		fmt.Printf("Выполнено? %v", v4.Index(i).Interface())
		fmt.Println("")
		fmt.Printf("Дата выполнения: %v\n", v5.Index(i).Interface())
		fmt.Println("")

	}
}
