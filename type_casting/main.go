package main

import "fmt"

func main() {
	// Числа
	var a int = 100
	var b float64 = float64(a)
	fmt.Println("int → float64:", b)

	// Байты и строки
	s := "Go"
	bts := []byte(s)
	fmt.Println("string → []byte:", bts)
	fmt.Println("[]byte → string:", string(bts))

	// Утверждение типа
	var i interface{} = 42
	if num, ok := i.(int); ok {
		fmt.Println("interface{} → int:", num)
	}

	// Пользовательский тип
	type UserID int
	var id UserID = UserID(a)
	fmt.Println("int → UserID:", id)
}
