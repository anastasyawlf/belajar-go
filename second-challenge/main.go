package main

import "fmt"

func main() {
	// Input kalimat
	kalimat := "Golang for Women"

	// Pecah kalimat menjadi kata-kata
	var kataKata []string
	kata := ""
	for _, char := range kalimat {
		if char != ' ' {
			kata += string(char)
		} else {
			kataKata = append(kataKata, kata)
			kata = ""
		}
	}
	if kata != "" {
		kataKata = append(kataKata, kata)
	}

	// Looping untuk mencetak setiap huruf dari setiap kata
	for _, kata := range kataKata {
		for _, huruf := range kata {
			fmt.Println(string(huruf))
		}
		fmt.Println() // Baris kosong untuk memisahkan setiap kata
	}

	// Hitung kemunculan setiap huruf dengan mapping
	hitungan := make(map[rune]int)
	for _, kata := range kataKata {
		for _, huruf := range kata {
			hitungan[huruf]++
		}
	}

	// Cetak hasil hitungan
	fmt.Println("Map hasil hitungan:")
	for huruf, jumlah := range hitungan {
		fmt.Printf("%c:%d ", huruf, jumlah)
	}
}
