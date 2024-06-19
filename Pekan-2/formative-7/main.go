package main

import "fmt"

//Struct Soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

//Struct Soal 2
type segitiga struct{
  alas, tinggi int
}

type persegi struct{
  sisi int
}

type persegiPanjang struct{
  panjang, lebar int
}

//Struct Soal 3
type phone struct{
   name, brand string
   year int
   colors []string
}

//Struct Soal 4
type movie struct{
	title, genre string
	duration, year int
}

func main() {
	//Soal 1
	var nanas = buah{"Nanas", "Kuning", false, 9000}
	var jeruk = buah{"Jeruk", "Oranye", true, 8000}
	var semangka = buah{"Semangka", "Hijau & Merah", true, 10000}
	var pisang = buah{"Pisang", "Kuning", false, 5000}

	fmt.Print("Nama Buah:", nanas.nama, " Warna :", nanas.warna,)
	biji(nanas.adaBijinya)
	fmt.Println("Harga :", nanas.harga)

	fmt.Print("Nama Buah:", jeruk.nama, " Warna :", jeruk.warna,)
	biji(jeruk.adaBijinya)
	fmt.Println("Harga :", jeruk.harga)

	fmt.Print("Nama Buah:", semangka.nama, " Warna :", semangka.warna,)
	biji(semangka.adaBijinya)
	fmt.Println("Harga :", semangka.harga)

	fmt.Print("Nama Buah:", pisang.nama, " Warna :", pisang.warna,)
	biji(pisang.adaBijinya)
	fmt.Println("Harga :", pisang.harga)

	//Soal 2
	var segi3 = segitiga{3,4}
	segi3.luasSegitiga()

	var segi4 = persegi{4}
	segi4.luasPersegi()

	var segi4Panjang = persegiPanjang{3,4}
	segi4Panjang.luasPersegiPanjang()

	//Soal 3
	//var hp = phone{name : "Galaxy A32", brand: "Samsung", year: 2021, colors:{"Ungu"}}

	//Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		fmt.Println(i+1, "title :", item.title, "duration :", item.duration, "genre :", item.genre, "year :", item.year)
	}
}

// Func Soal 1
func biji(AdaBiji bool){
	if AdaBiji ==false  {
		fmt.Print(" Ada Bijinya : Tidak ")
	}else{
		fmt.Print(" Ada Bijinya : Ada ")
	}

}

//Func Soal 2
func (l segitiga) luasSegitiga() {
	var luas = (l.alas * l.tinggi)/2
	fmt.Println("Luas Segitiga : ",luas)
}

func (l persegi) luasPersegi() {
	var luas = l.sisi * l.sisi
	fmt.Println("Luas Persegi : ",luas)
}

func (l persegiPanjang) luasPersegiPanjang() {
	var luas = l.lebar * l.panjang
	fmt.Println("Luas Persegi Panjang : ",luas)
}

//Func Soal 4
func tambahDataFilm(judul string, durasi int, genre string, tahun int, film *[]movie)  {
	var dataTampung = movie{title: judul, duration: durasi, genre: genre, year: tahun}
	*film = append(*film, dataTampung)
}
