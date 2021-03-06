package main

import "fmt"

//variadic param
// variadic parameter adalah parameter yang bisa diberi banyak nilai
// contoh variadic parameter seperti dibawah ini

func main() {
	createSentence("halo")
	createSentence("halo", "selamat siang")
	createSentence("halo", "andi", "dan", "bagus")
	createSentence("mencoba", "variadic", "param", "pada", "go", "lang", "golang")
}

func createSentence(words ...string) {
	result := ""
	for _, word := range words {
		result += " " + word
	}
	fmt.Println("hasil membuat kalimat", result)
}