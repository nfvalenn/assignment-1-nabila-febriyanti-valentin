package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataBiodata = map[int]Biodata{
	1: {"Muhammad Nahrowi", "Surabaya", "Software Enginners", "Ingin memperdalam pemrograman dengan Golang"},
	2: {"Muhammad Dermawan", "Bandung", "Backend Developer", "Tertarik dengan Golang"},
	3: {"Ikhsan Kurniawan", "Jakarta", "Fullstack Developer", "Ingin menambah pengetahuan bahasa pemrograman go"},
	4: {"Ariq Athallah", "Bandung", "Data Scientist", "Ingin mempelajari bidang baru"},
	5: {"Nabila Febriyanti Valentin", "Yogyakarta", "Mahasiswa", "Ingin mempelajari golang lebih dalam"},
}

func tampilkanData(nomor int) {
	biodata, ok := dataBiodata[nomor]
	if !ok {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan.")
		return
	}

	fmt.Println("Nama:", biodata.Nama)
	fmt.Println("Alamat:", biodata.Alamat)
	fmt.Println("Pekerjaan:", biodata.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", biodata.Alasan)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	nomorAbsen := args[1]

	var nomor int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomor)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	tampilkanData(nomor)
}
