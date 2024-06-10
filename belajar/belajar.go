package belajar

import (
	"fmt"
	"strconv"
)

type Siswa struct {
	Nama    string
	Usia    uint8
	Jurusan string
}

// function dengan parameter
func Tambah(a int, b int) int {
	return a + b
}

// function dengan parameter tanpa batas (varidic function)
// atau (rest parameter di JS)
func TambahBanyak(nilai ...int) (hasil int) {
	for _, n := range nilai {
		hasil += n
	}

	return hasil
}

func TampilkanSiswa(s Siswa) string {
	return "Nama : " + s.Nama + ", usia :" + strconv.FormatUint(uint64(s.Usia), 10) + ", Jurusan : " + s.Jurusan
}

// defer, panic & recover
func SomeMessage() {
	fmt.Println("Ini SomeMessage")
	if message := recover(); message != nil {
		fmt.Println("Recover : ", message)
	}
}
func ErrorTest(val bool) {
	defer SomeMessage()
	if val {
		panic("Terjadi error")
	}
}

// struct method
type Pegawai struct {
	Nama string
	Usia uint8
}

func (p Pegawai) Bertugas(kerjaan string) {
	fmt.Println("Pegawai", p.Nama, "bekerja sebagai", kerjaan)
}
