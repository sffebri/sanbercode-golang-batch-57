package main

import "fmt"

func main() {

	//Soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	change(&luasLigkaran, &kelilingLingkaran, 7)
	fmt.Println("Luas Lingkaran  :", luasLigkaran)
	fmt.Println("Keliling Lingkaran  :", kelilingLingkaran)

	//Soal 2
	var sentence string 
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"


	//Soal 3
	var buah = []string{}
	var listBuah = []string{
		"Jeruk",
		"Semangka",
		"Mangga",
		"Strawberry",
		"Durian",
		"Manggis",
		"Alpukat",
	}

	masukanBuah(&buah, listBuah)
	for i, item := range buah {
		fmt.Println(i+1, ".", item)
	}

	//Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Println(i+1, item)
	}

}

// Func Soal 1
func change(luas *float64, keliling *float64, jariJari float64) {
	var phi = 3.14
	*luas = phi * jariJari * jariJari
	*keliling = 2 * phi * jariJari
}

//Func Soal 2
func introduce(kalimat *string, nama string, jenisKelamin string, pekerjaan string, umur string)  {
	if jenisKelamin == "perempuan" {
		*kalimat = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}else if jenisKelamin == "laki-laki" {
		*kalimat = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}else{
		*kalimat = "Jenis Kelamin error"
	}
}

//Func Soal 3
func masukanBuah(buahBuahan *[]string, namaBuah []string)  {
	*buahBuahan = namaBuah
}

//Func Soal 4
func tambahDataFilm(judul string, durasi string, genre string, tahun string, film *[]map[string]string)  {
	var dataTampung = map[string]string{}
	dataTampung = map[string]string{
		"genre" : genre, "duration" : durasi, "year" :tahun, "title" : judul,
	}
	*film = append(*film, dataTampung)
}