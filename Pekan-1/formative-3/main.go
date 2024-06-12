package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjangPersegiPanjangInt, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjangInt, _ = strconv.Atoi(lebarPersegiPanjang)

	var alasSegitigaInt, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitigaInt,_ = strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
	var luasSegitiga int = (alasSegitigaInt * tinggiSegitigaInt) / 2

	fmt.Println("Luas Persegi Panjang :", luasPersegiPanjang)
	fmt.Println("keliling Persegi Panjang :", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga :", luasSegitiga)

	//Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indexNilaiJohn string
	var indexNilaiDoe string

	if nilaiJohn >= 80 {
		indexNilaiJohn = "A"
	}else if nilaiJohn >= 70 && nilaiJohn <80 {
		indexNilaiJohn = "B"
	}else if nilaiJohn >= 60 && nilaiJohn <70 {
		indexNilaiJohn = "C"
	}else if nilaiJohn >= 50 && nilaiJohn <60 {
		indexNilaiJohn = "D"
	}else if nilaiJohn < 50 {
		indexNilaiJohn = "C"
	}

	if nilaiDoe >= 80 {
		indexNilaiDoe = "A"
	}else if nilaiDoe >= 70 && nilaiDoe <80 {
		indexNilaiDoe = "B"
	}else if nilaiDoe >= 60 && nilaiDoe <70 {
		indexNilaiDoe = "C"
	}else if nilaiDoe >= 50 && nilaiDoe <60 {
		indexNilaiDoe = "D"
	}else if nilaiDoe < 50 {
		indexNilaiDoe = "C"
	}

	fmt.Println("Index Nilai Jhon: ", indexNilaiJohn)
	fmt.Println("Index Nilai Doe: ", indexNilaiDoe)

	//soal 3
	var tanggal = 18;
	var bulan = 9; 
	var tahun = 1998;

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		fmt.Println("Tidak ada nama bulannya")
	}

	fmt.Println(tanggal, namaBulan, tahun)

	//soal 4
	switch{
	case tahun >= 1944 && tahun <= 1964:
		fmt.Println("Baby boomer")
	case tahun >= 1965 && tahun <= 1979:
		fmt.Println("Generasi X")
	case tahun >= 1980 && tahun <= 1994:
		fmt.Println("Generasi Y")
	case tahun >= 1995 && tahun <= 2015:
		fmt.Println("Generasi Z")
	default:
		fmt.Println("Tidak tahu ya")
	}


}