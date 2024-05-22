package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 1000

type Patient struct {
	nama string
	asal string
	umur int
}

type Package struct {
	nama    string
	harga   int
	hasil   string
	tanggal tgl
}

type tgl struct {
	bulan int
	tahun int
}

var pasien [NMAX]Patient
var paket [NMAX]Package

var nama string
var pilihan int
var nPasien int = 0

func main() {
	clearScreen()
	fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
	fmt.Printf("Silahkan login terlebih dahulu.\n")

	if !login() {
		os.Exit(0)
	}

	for {
		fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
		fmt.Printf("1. Penambahan Data\n2. Pengubahan Data\n3. Penghapusan Data\n4. Laporan Pemasukan\n5. Pencarian Pasien\n6. Tampilkan Data Pasien Terurut\n7. Keluar\n")
		fmt.Printf("\nPilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Berapa pasien yang ingin ditambahkan? ")
			var penambahan int
			fmt.Scan(&penambahan)
			fmt.Printf("\nSilahkan masukkan data pasien sesuai urutan 'nama, usia, asal, jenis paket, harga paket, hasil MCU'\n")
			insert(penambahan)
		case 2:

			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin merubah data pasien atas nama siapa? ")
			fmt.Scan(&nama)
			edit(nama)
			clearScreen()
		case 3:

			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin menghapus data pasien atas nama siapa? ")
			delete()
			clearScreen()
		case 4:

			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin menampilkan laporan pemasukan MCU pada periode apa? ")
			var periode int
			fmt.Scan(&periode)
		case 5:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Printf("Ingin mencari daftar pasien berdasarkan apa?\n1. Berdasarkan paket MCU\n2. Berdasarkan Periode\n3. Berdasarkan nama pasien\n")
			fmt.Scan(&pilihan)
		case 6:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Waktu MCU\n2. Paket MCU\n3. Tampilkan Semua\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 3:
				clearScreen()
				for i := 0; i < nPasien; i++ {
					fmt.Println(pasien[i].nama, pasien[i].umur, pasien[i].asal, paket[i].nama, paket[i].harga, paket[i].hasil)
				}
			}
		case 7:
			break
		default:
			clearScreen()
			fmt.Printf("Menu tidak tersedia\n")
		}
		if pilihan == 7 {
			clearScreen()
			fmt.Println("Terima kasih telah menggunakan aplikasi ini")
			break
		}
	}
}

func login() bool {
	var username, password string
	var attempts int

	for attempts < 3 {
		fmt.Println("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Println("Masukkan password: ")
		fmt.Scan(&password)

		if username == "pegawai" && password == "pegawai123" {
			clearScreen()
			fmt.Println("Login berhasil!")
			return true
		} else {
			fmt.Println("Username atau password salah. Sisa percobaan:", 3-attempts-1)
			attempts++
		}
	}
	clearScreen()
	fmt.Println("Anda telah salah memasukkan username dan password 3 kali. Silakan coba lagi nanti.")
	return false
}

func insert(penambahan int) {
	for i := nPasien; i < (nPasien + penambahan); i++ {
		fmt.Scanln(&pasien[i].nama, &pasien[i].umur, &pasien[i].asal, &paket[i].nama, &paket[i].harga, &paket[i].hasil)
	}
	nPasien += penambahan
}

func edit(x string) {
	var idx int = -1
	var find bool = false
	var i int = 0
	for i < idx && !find {
		if pasien[i].nama == x {
			idx = i
			find = true
		}
		i++
	}

	if find {
		fmt.Println("Masukkan ulang data yang benar:")
		fmt.Scanln(&pasien[idx].nama, &pasien[idx].umur, &pasien[idx].asal, &paket[idx].nama, &paket[idx].harga, &paket[idx].hasil)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func delete() {
	var nama string
	fmt.Scan(&nama)
	var idx int = -1
	var find bool = false
	var i int = 0
	for i < nPasien && !find {
		if pasien[i].nama == nama {
			idx = i
			find = true
		}
		i++
	}

	if find {
		for i := idx; i < nPasien-1; i++ {
			pasien[i] = pasien[i+1]
			paket[i] = paket[i+1]
		}
		nPasien--
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		//Do nothing for unsupported OS
	}
}
