package main // объявляем  пакет, которому принадлежит текущий код

import "fmt" // уточняем, что будем работать с пакетом foomt для функций форматированного ввода и вывода

func main() { // объявляю функцию main - точку входа
	fmt.Println("go terms - понятия golang")
	fmt.Println(" ")
	fmt.Println("basic types / базовые типы")
	fmt.Println(" ")
	fmt.Println("bool - логические тип, константы true/false")
	fmt.Println(" ")
	fmt.Println("string - тип строки,набор строковых значений")
	fmt.Println("Строковое значение это (возможно, пустая) последовательность байтов")
	fmt.Println(" ")
	fmt.Println("Числовые типы:")
	fmt.Println("int - того же размера, что и uint")
	fmt.Println("	int8 - набор всех 8-битных целых чисел со знаком (от -128 до 127)")
	fmt.Println("	int16 - 16-битных (от -32768 до 32767)")
	fmt.Println("	int32 - 32-битных (от -2147483648 до 2147483647)")
	fmt.Println("	int64 - 64-битных (от -9223372036854775808 до 9223372036854775807")
	fmt.Println("uint - либо 32, либо 64 бит")
	fmt.Println("	uint8 - набор всех беззнаковых 8-битных целых чисел (от 0 до 255)")
	fmt.Println("	uint16 - 16-битных (от 0 до 65535)")
	fmt.Println("	uint32 - 32-битных (от 0 до 4294967295)")
	fmt.Println("	uint64 - 64-битных (от 0 до 18446744073709551615)")
	fmt.Println("uintptr - целое число без знака, достаточно большое, чтобы хранить неинтерпретированные биты значения указателя")
	fmt.Println("byte - псевдоним для uint8")
	fmt.Println("rune - севдоним для int32 represents a Unicode code point")
	fmt.Println("float32 - набор всех 32-битных чисел с плавающей запятой IEEE-754")
	fmt.Println("float64 - 64-битных")
	fmt.Println("complex64 - набор всех комплексных чисел с вещественной и мнимой частями float32")
	fmt.Println("complex128 - набор всех комплексных чисел с вещественной и мнимой частями float64")
}
