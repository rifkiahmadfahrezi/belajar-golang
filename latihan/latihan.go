package latihan

import (
	"fmt"
	"strconv"
)

// Soal latihan dari chatGPT

//Soal 1: Menghitung Bilangan Prima
// Buatlah program yang meminta pengguna untuk memasukkan sebuah angka dan menentukan apakah angka tersebut adalah bilangan prima atau bukan.

func IsPrimeNumber() {
	var angka int

	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&angka)
	if angka <= 1 {
		fmt.Println(angka, "bukan bilangan prima")
		return
	} else if angka > 1 {
		for i := 2; i <= angka; i++ {
			if angka%i == 0 {
				fmt.Println(strconv.Itoa(angka) + " = " + "true")
				return
			} else {
				fmt.Println(strconv.Itoa(angka) + " = " + "false")
				return
			}
		}
	} else {
		fmt.Println("Invalid input")
	}
}

// simple kalkulator
func Kalkulator() {

	var angka1 int
	var angka2 int
	var operator string

	fmt.Print("Angka ke 1 :")
	fmt.Scan(&angka1)
	fmt.Print("Angka ke 2 :")
	fmt.Scan(&angka2)
	// input operaror
	fmt.Print("Masukkan operator [+ - * / %] :")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("Hasil penjumlahan %v + %v = %v", angka1, angka2, (angka1 + angka2))
		return
	case "-":
		fmt.Printf("Hasil penjumlahan %v - %v = %v", angka1, angka2, (angka1 - angka2))
		return
	case "/":
		fmt.Printf("Hasil penjumlahan %v / %v = %v", angka1, angka2, (angka1 / angka2))
		return
	case "*":
		fmt.Printf("Hasil penjumlahan %v * %v = %v", angka1, angka2, (angka1 * angka2))
		return
	case "%":
		fmt.Printf("Hasil penjumlahan %v %% %v = %v", angka1, angka2, (angka1 % angka2))
		return
	default:
		fmt.Println("Invalid operator")
	}

}

func Faktorial() {
	var angka int

	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	var hasil int = 1

	for i := 1; i <= angka; i++ {
		hasil *= i
	}

	defer fmt.Printf("\n Hasil Faktorial %v adalah %v", angka, hasil)
}
