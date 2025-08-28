package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Явное
	var age int
	age = 0
	var y int = 42
	//Неявное
	var z = 42 // Тип x будет int. Как докажем? fmt.Println(fmt.Sprintf("%T", z))
	fmt.Println(age, y, z)

	//basicTypes()
	//interfaceType()
}

func basicTypes() {
	//https://go.dev/tour/basics/11
	//Простые типы

	// Логический тип
	var a bool = true
	var b bool = false
	fmt.Printf("bool: %v, %v\n", a, b)

	// Строковый тип
	var s string = "Hello, 世界"
	fmt.Printf("string: %s\n", s)

	// Целочисленные типы со знаком
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = -2147483648
	var i64 int64 = 9223372036854775807
	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n", i, i8, i16, i32, i64)

	// Целочисленные типы без знака
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var uptr uintptr = 0x12345678
	fmt.Printf("uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d, uintptr: %v\n", u, u8, u16, u32, u64, uptr)

	// byte — псевдоним для uint8
	var b8 byte = 255
	fmt.Printf("byte (uint8): %d\n", b8)

	// rune — псевдоним для int32, представляет Unicode-символ
	var r rune = '世'
	fmt.Printf("rune (int32): %c (%d)\n", r, r)

	// Дробные числа с плавающей точкой
	var f32 float32 = 3.1415926535
	var f64 float64 = 3.14159265358979323846
	fmt.Printf("float32: %.10f, float64: %.15f\n", f32, f64)

	// Комплексные числа
	var c64 complex64 = complex(1.0, 2.0) // real: 1.0, imag: 2.0
	var c128 complex128 = complex(3.0, 4.0)
	fmt.Printf("complex64: %v, complex128: %v\n", c64, c128)

	// Краткое объявление (с использованием :=)
	shortBool := true
	shortString := "Go"
	shortInt := 100
	shortFloat := 2.71828
	shortComplex := 1 + 2i
	fmt.Printf("Short declarations: %v, %s, %d, %f, %v\n", shortBool, shortString, shortInt, shortFloat, shortComplex)
}
func interfaceType() {
	var x interface{} = []int{1, 2, 3}
	xType := fmt.Sprintf("%T", x)
	fmt.Println(xType)

	//caseTypes()
	//reflectionTypes()

}

func reflectionTypes() {
	// Рефлексия для получения информации о типах
	var x interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(x)
	xValue := reflect.ValueOf(x)
	fmt.Println(xType, xValue)
}

func caseTypes() {
	// Переключение типов
	var x interface{} = 1.6
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	default:
		fmt.Println("unknown")
	}
}
