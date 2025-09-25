package main

import (
	"fmt"
)

var data []string // slice untuk simpan data

func main() {
	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Lihat Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			tambahData()
		case 2:
			lihatData()
		case 3:
			updateData()
		case 4:
			hapusData()
		case 5:
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahData() {
	var nama string
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	data = append(data, nama)
	fmt.Println("Data berhasil ditambahkan!")
}

func lihatData() {
	fmt.Println("\n===== DATA =====")
	if len(data) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, v := range data {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func updateData() {
	lihatData()
	if len(data) == 0 {
		return
	}

	var index int
	fmt.Print("Masukkan nomor data yang mau diupdate: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(data) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	var namaBaru string
	fmt.Print("Masukkan nama baru: ")
	fmt.Scanln(&namaBaru)
	data[index-1] = namaBaru
	fmt.Println("Data berhasil diupdate!")
}

func hapusData() {
	lihatData()
	if len(data) == 0 {
		return
	}

	var index int
	fmt.Print("Masukkan nomor data yang mau dihapus: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(data) {
		fmt.Println("Nomor tidak valid!")
		return
	}

	data = append(data[:index-1], data[index:]...)
	fmt.Println("Data berhasil dihapus!")
}
