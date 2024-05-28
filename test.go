package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const NMAX = 1000

type Patient struct {
	namapasien string
	asal       string
	umur       int
}

type Package struct {
	namapaket string
	tanggal   int
}

type rekaphasil struct {
	tanggal    int
	namapaket  string
	namapasien string
}

var pasien [NMAX]Patient
var rekap [NMAX]rekaphasil
var paket [NMAX]Package

var nPasien int

func main() {
	var penambahan int
	var pilihan int

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
			fmt.Scan(&penambahan)
			fmt.Printf("\nSilahkan masukkan data pasien sesuai urutan yang diminta\n")
			insert(penambahan)
		case 2:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin merubah data pasien atas nama siapa? ")
			var nama string
			fmt.Scan(&nama)
			edit(nama)
		case 3:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin menghapus data pasien atas nama siapa? ")
			delete()
		case 4:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin menampilkan laporan pemasukan MCU pada periode apa? ")
			var periode int
			fmt.Scan(&periode)
			laporanPemasukan(periode)
		case 5:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Printf("Ingin mencari daftar pasien berdasarkan apa?\n1. Berdasarkan paket MCU\n2. Berdasarkan Periode\n3. Berdasarkan nama pasien\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			/* 				case 1:
			                    clearScreen()
								fmt.Println("Masukkan Nama Paket")
								var nama string
								fmt.Scan(&nama)
								searchPaket(&nama)
			                case 2:
								fmt.Println("Masukkan Periode")
								var periode int
								fmt.Scanln(&periode)
								searchPeriode(&periode)
			                    clearScreen() */
			/*  case 3:
								fmt.Println("Masukkan Nama Pasien")
								var nama string
								fmt.Scan(&nama)
								searchNama(&nama)
			                    clearScreen() */
			}
		case 6:
			clearScreen()
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Periode MCU\n2. Paket MCU\n3. Nama\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				clearScreen()
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				sortPeriode()
			case 2:
				clearScreen()
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				sortPaket()
			case 3:
				clearScreen()
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				sortNama()
			}
		case 7:

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
	/* fitur login yang akan meminta user untuk memasukkan kredensial username dan password.
	Apabila kredensial yang dimasukkan benar, maka user akan berhasil login. Jika salah,
	user akan diberikan kesempatan 3 kali untuk mencoba. */
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
	// procedure yang digunakan untuk memasukkan data pasien sesuai dengan jumlah yang diinginkan
	fmt.Println("Nama Pasien, Asal, Umur, Nama Paket, Tahun Registrasi")
	for i := nPasien; i < (nPasien + penambahan); i++ {
		fmt.Scanln(&pasien[i].namapasien, &pasien[i].asal, &pasien[i].umur, &paket[i].namapaket, &paket[i].tanggal)
		rekap[i].tanggal = paket[i].tanggal
		rekap[i].namapaket = paket[i].namapaket
		rekap[i].namapasien = pasien[i].namapasien
		rekap[i].tanggal = paket[i].tanggal
	}
	nPasien += penambahan
}

func edit(x string) {
	var pilihan int
	var find bool = false
	var i int = 0
	for i < nPasien && !find {
		if pasien[i].namapasien == x {
			find = true
		} else {
			i++
		}
	}

	if find {
		fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
		fmt.Println("Ingin merubah data apa?")
		fmt.Printf("1. Nama Pasien\n2. Asal Pasien\n3. Umur Pasien\n4. Nama Paket\n5. Tahun Registrasi\n")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
			fmt.Println("Masukkan nama pasien baru: ")
			fmt.Scan(&pasien[i].namapasien)
			rekap[i].namapasien = pasien[i].namapasien
		case 2:
			fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
			fmt.Println("Masukkan asal pasien baru: ")
			fmt.Scan(&pasien[i].asal)
		case 3:
			fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
			fmt.Println("Masukkan umur pasien baru: ")
			fmt.Scan(&pasien[i].umur)
		case 4:
			fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
			fmt.Println("Masukkan nama paket baru: ")
			fmt.Scan(&paket[i].namapaket)
			rekap[i].namapaket = paket[i].namapaket
		case 5:
			fmt.Printf("===============================\n Aplikasi Medical Check-Up \n===============================\n")
			fmt.Println("Masukkan tahun registrasi baru: ")
			fmt.Scan(&paket[i].tanggal)
			rekap[i].tanggal = paket[i].tanggal
		default:
			fmt.Println("Menu tidak tersedia")
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func delete() {
	/* procedure yang digunakan untuk menghapus data pasien berdasarkan nama pasien yang diinginkan. Untuk mencari nama pasien,
	   digunakan sequential search */
	var nama string
	fmt.Scan(&nama)
	var idx int = -1
	var find bool = false
	var i int = 0
	for i < nPasien && !find {
		if pasien[i].namapasien == nama {
			idx = i
			find = true
		}
		i++
	}
	clearScreen()
	if find {
		for i := idx; i < nPasien-1; i++ {
			pasien[i] = pasien[i+1]
			paket[i] = paket[i+1]
			rekap[i] = rekap[i+1]
		}
		nPasien--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/k", "cls")
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

func laporanPemasukan(periode int) {
	var total int
	fmt.Printf("Data pasien yang melakukan MCU pada periode %d:\n", periode)
	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Registrasi")
	for i := 0; i < nPasien; i++ {
		if rekap[i].tanggal == periode {
			fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", rekap[i].namapasien, pasien[i].asal, pasien[i].umur, rekap[i].namapaket, rekap[i].tanggal)
			switch rekap[i].namapaket {
			case "A":
				total += 50000
			case "B":
				total += 75000
			case "C":
				total += 100000
			}
		}
	}
	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Total pemasukan pada tahun", periode, "adalah", total)
}

func cetak() {
	clearScreen()
	fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
	fmt.Printf("%-20s %-10s %-5s %-15s %-20s\n", "Nama Pasien", "Asal", "Umur", "Nama Paket", "Tahun Registrasi")
	for i := 0; i < nPasien; i++ {
		fmt.Printf("%-20s %-10s %-5d %-15s %-20d\n", pasien[i].namapasien, pasien[i].asal, pasien[i].umur, paket[i].namapaket, paket[i].tanggal)
	}
}

func sortPaket() {
	//mengurutkan data pasien menurut Paket MCU yang dipilih dengan konsep Insertion Sort
	var j int
	var tempPaket Package
	var tempPasien Patient
	var tempRekap rekaphasil
	var pass int = 1
	for pass <= nPasien-1 {
		j = pass
		for j > 0 && (paket[j].namapaket < paket[j-1].namapaket || (paket[j].namapaket == paket[j-1].namapaket && strings.ToLower(pasien[j].namapasien) < strings.ToLower(pasien[j-1].namapasien))) {
			tempPaket = paket[j]
			tempPasien = pasien[j]
			tempRekap = rekap[j]

			paket[j] = paket[j-1]
			pasien[j] = pasien[j-1]
			rekap[j] = rekap[j-1]

			paket[j-1] = tempPaket
			pasien[j-1] = tempPasien
			rekap[j-1] = tempRekap

			j--
		}
		pass++
	}
	cetak()
}

func sortNama() {
	//mengurutkan data pasien menurut nama pasien dengan konsep Selection Sort
	var i, j, minIdx int
	var tempPasien Patient
	var tempPaket Package
	var tempRekap rekaphasil
	for i = 0; i < nPasien-1; i++ {
		minIdx = i
		for j = i + 1; j < nPasien; j++ {
			if strings.ToLower(pasien[j].namapasien) < strings.ToLower(pasien[minIdx].namapasien) {
				minIdx = j
			}
		}
		tempPasien = pasien[minIdx]
		pasien[minIdx] = pasien[i]
		pasien[i] = tempPasien

		tempPaket = paket[minIdx]
		paket[minIdx] = paket[i]
		paket[i] = tempPaket

		tempRekap = rekap[minIdx]
		rekap[minIdx] = rekap[i]
		rekap[i] = tempRekap
	}
	cetak()
}

func sortPeriode() {
	//mengurutkan data pasien menurut periode MCU dengan konsep Selection Sort
	var i, j, minIdx int
	var tempPasien Patient
	var tempPaket Package
	var tempRekap rekaphasil
	for i = 0; i < nPasien-1; i++ {
		minIdx = i
		for j = i + 1; j < nPasien; j++ {
			if paket[j].tanggal < paket[minIdx].tanggal || (paket[j].tanggal == paket[minIdx].tanggal && strings.ToLower(pasien[j].namapasien) < strings.ToLower(pasien[minIdx].namapasien)) {
				minIdx = j
			}
		}
		tempPasien = pasien[minIdx]
		pasien[minIdx] = pasien[i]
		pasien[i] = tempPasien

		tempPaket = paket[minIdx]
		paket[minIdx] = paket[i]
		paket[i] = tempPaket

		tempRekap = rekap[minIdx]
		rekap[minIdx] = rekap[i]
		rekap[i] = tempRekap
	}

	cetak()
}
