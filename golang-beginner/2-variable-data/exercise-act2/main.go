package main

import "fmt"

// Dari lecture-act2 yang diberikan
// Cobalah untuk mempraktikan kemampuan kalian untuk membuat kalkulator sederhana menggunakan if else dan operator

// Demo:
// Calculator
// - Input: Nilai Pertama, Nilai Kedua, Pilih Operator
// - Output: Result Operation

// Contoh:
// Input:
// - Masukkan nilai pertama: 10
// - Masukkan nilai pertama: 5
// 1: Addition (+)
// 2: Subtraction (-)
// 3: Multiplication (*)
// 4: Division (/)
// - Enter choice: 2
// Output:
// - Result: 5

func main() {
	var num1 int
	var num2 int
	fmt.Printf("Masukkan nilai pertama: ")
	fmt.Scan(&num1)
	fmt.Printf("Masukkan nilai kedua: ")
	fmt.Scan(&num2)

	var choice int = 0
	var result int = 0
	fmt.Println("1: Addition (+)")
	fmt.Println("2: Subtraction (-)")
	fmt.Println("3: Multiplication (*)")
	fmt.Println("4: Division (/)")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	//TODO: answer here
	if choice == 1 {
		result = num1 + num2
		fmt.Println("Result: ", result)
	} else if choice == 2 {
		result = num1 - num2
		fmt.Println("Result: ", result)
	} else if choice == 3 {
		result = num1 * num2
		fmt.Println("Result: ", result)
	} else if choice == 4 {
		result = num1 / num2
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("Invalid choice")
	}
}
