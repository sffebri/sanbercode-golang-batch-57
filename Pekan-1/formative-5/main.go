package main

import (
	"fmt"
	"strings"
)

func main() {

	//Soal 1
	panjang := 12
	lebar := 4
	tinggi :=8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) 

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) 

	//Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	//Soal 4
	var dataFilm = []map[string]string{}
	var tambahDataFilm = func (tambahData  ...string) []map[string]string{
		var dataTampung = map[string]string{}
		var genre, jam,tahun, judul string
		for i, data := range tambahData {
			if i == 2 {
				genre = data
			}else if i == 1 {
				jam = data
			}else if i == 3{
				tahun = data
			}else if i == 0 {
				judul = data
			}
			
		}
		dataTampung = map[string]string{
			"genre" : genre, "jam" : jam, "tahun" :tahun, "title" : judul,
		}
		dataFilm = append(dataFilm, dataTampung) 
		return dataFilm
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}


// Func Soal 1
func luasPersegiPanjang(p, l int) int{
	var val = p * l
	return val
}

func kelilingPersegiPanjang(p, l int) int {
	var val = 2 * (p + l)
	return val
}

func volumeBalok(p, l, t int)  int {
	var val = p * l * t
	return val
}

// Func Soal 2
func introduce(nama string, jenisKelamin string, pekerjaan string, umur string)  string{
	var val string
	if jenisKelamin == "perempuan" {
		val = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}else if jenisKelamin == "laki-laki" {
		val = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}else{
		val = "Jenis Kelamin error"
	}

	return val
}

//Func Soal 3

func buahFavorit(nama string, fruits ...string)  string{
	namaLower := strings.ToLower(nama) // mengubah nama menjadi huruf kecil
	val := "halo nama saya " + namaLower + " dan buah favorit saya adalah "
	for i, fruit := range fruits {
		if i == 0 {
			val = val + `"` + fruit + `"`
		}else{
			val = val +"," + `"` + fruit + `"`
		}
		
	}
	
	return val
}
