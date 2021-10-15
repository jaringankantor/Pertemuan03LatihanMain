package main

import (
	"fmt"

	modulasi "github.com/jaringankantor/Pertemuan03LatihanModule"
)

//interface struct

type PersegiInterface interface {
	Keliling() int
	Luas() int
}

type InputPersegi struct {
	a, b int
}

func (in InputPersegi) Keliling() int {
	return (in.a + in.b) * 2
}

func (in InputPersegi) Luas() int {
	return in.a * in.b
}

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
	modulasi.TampilData()

	//struct module
	welcome := modulasi.Pegawai{Nama: "Anwar"}
	welcome.SelamatDatang("Budi")

	//Anonymous struct
	modulasi.AnonymousDataOrtu()

	//interface struct

	persegi := InputPersegi{
		a: 23,
		b: 45,
	}

	fmt.Println("Jika a=", persegi.a, ", b=", persegi.b, ". Maka keliling=", persegi.Keliling())
	fmt.Println("Jika a=", persegi.a, ", b=", persegi.b, ". Maka luas=", persegi.Luas())
}
