package main

import (
	"fmt"

	modulasi "github.com/jaringankantor/Pertemuan03LatihanModule"
)

func main() {
	a := 3
	b := 5
	name := "anwar"

	//function
	modulasi.About()

	//function return value
	fmt.Println("a=", a, ", b=", b, ". a+b=", modulasi.Tambah(a, b))

	//function dengan multiple return value
	var keliling, luas = modulasi.Persegi(a, b)
	fmt.Println("Diketahui persegi sisi a=", a, ", b=", b, ". Maka keliling=", keliling, ", luas=", luas)

	//function anonymous
	isAktif := func(name string) bool {
		switch name {
		case "anwar":
			return true
		case "budi":
			return true
		default:
			return false
		}
	}
	modulasi.User(name, isAktif)

	//struct
	modulasi.TampilData
}
