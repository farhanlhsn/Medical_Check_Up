package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const NMAX int = 1000

type Patient struct {
	namapasien string
	asal string
	umur int
}

type Package struct {
	namapaket string
	harga   int
	tanggal int
}

type rekaphasil struct {
	tanggal int
	namapaket string
	namapasien string
}

var pasien[NMAX]Patient
var rekap[NMAX]rekaphasil
var paket[NMAX]Package

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
			fmt.Printf("\nSilahkan masukkan data pasien sesuai urutan yang diminta\n")
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
			fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Periode MCU\n2. Paket MCU\n3. Nama (Terurut dari A-Z)\n")
			fmt.Scan(&pilihan)
			switch pilihan {
			case 1:
				clearScreen()
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				//sortPeriode()
			case 2:
				clearScreen()
                fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
                fmt.Println("Masukkan Nama Paket")
				var namapaket string
				fmt.Scanln(&namapaket)
                //sortPaket(namapaket)
			case 3:
				clearScreen()
                fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
                sortNama()
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
	fmt.Println("Nama Pasien, Asal, Umur, Nama Paket, Tahun Registrasi")
	for i := nPasien; i < (nPasien + penambahan); i++ {
		fmt.Scanln(&pasien[i].namapasien, &pasien[i].asal,&pasien[i].umur,&paket[i].namapaket,&paket[i].tanggal)
		rekap[i].tanggal = paket[i].tanggal
		rekap[i].namapaket = paket[i].namapaket
		rekap[i].namapasien = pasien[i].namapasien
		rekap[i].tanggal = paket[i].tanggal
	}
	nPasien += penambahan
}

func edit(x string) {
	var find bool = false
	var i int = 0
	for i < nPasien && !find {
		if pasien[i].namapasien == x {
			find = true
			i=i-1
		}
		i++
	}

	if find {
		fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
		fmt.Println("Ingin merubah data apa?")
		fmt.Printf("1. Nama Pasien\n2. Asal Pasien\n3. Umur Pasien\n4. Nama Paket\n5. Tahun Registrasi\n")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
            fmt.Println("Masukkan nama pasien baru: ")
            fmt.Scan(&pasien[i].namapasien)
			rekap[i].namapasien = pasien[i].namapasien
        case 2:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Masukkan asal pasien baru: ")
			fmt.Scan(&pasien[i].asal)
		case 3:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
            fmt.Println("Masukkan umur pasien baru: ")
            fmt.Scan(&pasien[i].umur)
        case 4:
			fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Masukkan nama paket baru: ")
			fmt.Scan(&paket[i].namapaket)
			rekap[i].namapaket = paket[i].namapaket
		case 5: 
		    fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
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

	if find {
		for i := idx; i < nPasien-1; i++ {
			pasien[i] = pasien[i+1]
			paket[i] = paket[i+1]
			rekap[i] = rekap[i+1]
		}
		nPasien--
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
	for i := 0; i < nPasien; i++ {
		if rekap[i].tanggal == periode {
            total += paket[i].harga
        }
    }
	clearScreen()
	fmt.Println("Total pemasukan pada tahun", periode, "adalah", total)
}

func cetak() {
	clearScreen()
    fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
    fmt.Println("Nama Pasien, Asal, Umur, Nama Paket, Tahun Registrasi")
    for i := 0; i < nPasien; i++ {
        fmt.Println(pasien[i].namapasien, pasien[i].asal, pasien[i].umur, paket[i].namapaket, paket[i].tanggal)
    }
}

/* func sortPaket(paket string) {
	var j int
	var tempPaket Package
	var tempPasien Patient
	var tempRekap rekaphasil
	var pass int = 1
	for pass <= nPasien {
		j = pass - 1
		for j >= 0 && paket[j].namapaket >= tempPaket[pass].namapasien {
			paket[j+1] = paket[j]
			pasien[j+1] = pasien[j]
			rekap[j+1] = rekap[j]
			j--
		}
		paket[j+1] = tempPaket[pass]
		pasien[j+1] = tempPasien[pass]
		rekap[j+1] = tempRekap[pass]
		pass++
	}
} */

func sortNama() {
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

		cetak()
	}
}


