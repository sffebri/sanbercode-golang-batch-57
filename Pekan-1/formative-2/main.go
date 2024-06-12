package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//Soal 1
	var kata1 string = "Bootcamp"
	var kata2 string = "Digital"
	var kata3 string = "Skill"
	var kata4 string = "Sanbercode"
	var kata5 string = "Golang"
	
	fmt.Println(kata1, kata2, kata3, kata4, kata5)

	//Soal 2
	halo := "Halo Dunia"
    var find = "Dunia"  
	var replaceWith = "Golang"

	var newText = strings.Replace(halo, find, replaceWith, -1)
	fmt.Println(newText)

	//Saol 3
	var kataPertama = "saya";
	var kataKedua = "senang";
	var kataKetiga = "belajar";
	var kataKeempat = "golang";

	var kata2find = "s"  
	var kata2replaceWith = "S"
	kataKedua = strings.Replace(kataKedua, kata2find, kata2replaceWith, 1)

	var kata3find = "r"  
	var kata3replaceWith = "R"
	kataKetiga = strings.Replace(kataKetiga, kata3find, kata3replaceWith, 1)
	kataKeempat = strings.ToTitle(kataKeempat)
	fmt.Println(kataPertama, kataKedua, kataKetiga, kataKeempat)

	//Soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";	

	var angkaPertamaNew, err1 = strconv.Atoi(angkaPertama)
	var angkaKeduaNew, err2 = strconv.Atoi(angkaKedua)
	var angkaKetigaNew, err3 = strconv.Atoi(angkaKetiga)
	var angkaKeempatNew, err4 = strconv.Atoi(angkaKeempat)

    if (err1 == nil && err2 == nil && err3 == nil && err4 == nil) {
		var hasil = angkaPertamaNew + angkaKeduaNew + angkaKetigaNew + angkaKeempatNew
        fmt.Println(hasil) 
    }

	//Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var findHalo = "halo"  
	var replaceWithHi = "Hi"

	var newKalimat = `"` + strings.Replace(kalimat, findHalo, replaceWithHi, -1)+ `"`
	fmt.Println(newKalimat,`- `,angka) 

}