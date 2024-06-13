package main

import "fmt"

func main() {

	//Soal 1
	for i := 1; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else{
			if i % 3 == 0{
				fmt.Println(i, "- I Love Coding")
			}else{
				fmt.Println(i, "- Santai")
			}
		}
		
	}

	//Soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
		
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	//soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}

	for i, sayuran := range sayuran{
		fmt.Print(i+1, ". ", sayuran)
		fmt.Println("")
	}

	//Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volume = 1

	for key, val := range satuan{
		volume = volume * val
		fmt.Println(key, " = ", val)
	}
	fmt.Println("Volume Balok = ", volume)

}