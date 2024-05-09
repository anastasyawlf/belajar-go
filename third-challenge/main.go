package main

import (
	"fmt"
	"os"
)

// Struct untuk merepresentasikan data teman di kelas
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman di kelas
var temanKelas = []Teman{
	{"Alfina", "Bogor", "Mahasiswi", "Ingin mempelajari bahasa Go untuk meningkatkan skill pemrograman"},
	{"Angela", "Madiun", "Karyawan BUMN", "Tertarik mempelajari pemrograman melalui bahasa Go"},
	{"Fasaya", "Makassar", "Web Developer", "Mendapat rekomendasi dari teman tentang kegunaan Go dalam pengembangan web"},
	{"Kartika", "Bekasi", "Mobile Developer", "Selain belajar Go, juga mencari teman ambis dan ingin tahu bagaimana cara mengajar"},
	{"Najya", "Solo", "Mahasiswi", "Ingin mendalami backend menggunakan bahasa Go"},
}

// Function untuk menampilkan data teman berdasarkan nama
func DataTeman(nama string) {
	found := false
	for _, teman := range temanKelas {
		if teman.Nama == nama {
			found = true
			fmt.Println("Nama:", teman.Nama)
			fmt.Println("Alamat:", teman.Alamat)
			fmt.Println("Pekerjaan:", teman.Pekerjaan)
			fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
			break
		}
	}
	if !found {
		fmt.Println("Data dengan nama", nama, "tidak tersedia.")
	}
}

func main() {
	args := os.Args[1:] // Mengambil argumen dari CLI, mengabaikan argumen pertama (setelah nama program)
	if len(args) < 1 {  // Apabila tidak ada argumen
		fmt.Println("Gunakan kode berikut: go run main.go <Nama>")
		fmt.Println("Contoh: go run main.go Angela")
		return
	}

	nama := args[0]
	DataTeman(nama)
}
