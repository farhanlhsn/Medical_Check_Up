package main

import "fmt"

const NMAX int = 100
type Patient struct {
	nama string
	asal string
	umur int
}

type Package struct {
	nama string
	harga int
	hasil string
}


func main() {
	var pasien [NMAX]Patient
	var paket [NMAX]Package
	var pilihan int
	var nPasien int

	for {
		fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
		fmt.Printf("1. Penambahan Data\n2. Pengubahan Data\n3. Penghapusan Data\n4. Laporan Pemasukan\n5. Pencarian Pasien\n6. Tampilkan Data Pasien Terurut\n7. Keluar\n")
		fmt.Printf("\nPilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
			case 1:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Println("Berapa pasien yang ingin ditambahkan? ")
				var penambahan int
				fmt.Scan(&penambahan)
				fmt.Println("Ingin menambahkan pasien dari indeks ke berapa? ")
				var idx int
				fmt.Scan(&idx)
				fmt.Printf("\nSilahkan masukkan data pasien sesuai urutan 'nama, umur, asal, jenis paket, harga paket, hasil MCU'\n")
				insert(&pasien, &paket, &nPasien, penambahan, idx)
				
			case 2:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Println("Ingin merubah data pasien atas nama siapa? ")
				var nama string
				fmt.Scan(&nama)
				edit(&pasien, &paket, nPasien, nama)
			case 3:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Println("Ingin menghapus data pasien atas nama siapa? ")
				var nama string
				fmt.Scan(&nama)
				delete(&pasien, &paket, &nPasien, nama)
			case 4:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Println("Ingin menampilkan laporan pemasukan MCU pada periode apa? ")
				var periode int
				fmt.Scan(&periode)
			case 5:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Printf("Ingin mencari daftar pasien berdasarkan apa?\n1. Berdasarkan paket MCU\n2. Berdasarkan Periode\n3. Berdasarkan nama pasien\n")
				var pilihan int
				fmt.Scan(&pilihan)
			case 6:
				fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
				fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Waktu MCU\n2. Paket MCU\n3. Tampilkan Semua\n")
				fmt.Scan(&pilihan)
				switch pilihan {
				case 3:
					for i := 0; i < nPasien; i++ {
						fmt.Println(pasien[i].nama, pasien[i].umur, pasien[i].asal, paket[i].nama, paket[i].harga, paket[i].hasil)
					}
				}
			case 7:
				break
			default:
				fmt.Printf("Menu tidak tersedia\n")
		}
		if pilihan == 7 {
            break
        }
	}
}

func insert(pasien *[NMAX]Patient, paket *[NMAX]Package, nPasien *int, penambahan, idx int) {
	*nPasien+=penambahan

	for i:=idx; i<*nPasien; i++ {
		fmt.Scanln(&pasien[i].nama, &pasien[i].umur, &pasien[i].asal, &paket[i].nama, &paket[i].harga, &paket[i].hasil)
	}
}

func edit(pasien *[NMAX]Patient, paket *[NMAX]Package, n int, x string) {
	var idx int = -1
	var find bool = false
	var i int = 0
	for i<n && !find{
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

func delete(pasien *[NMAX]Patient, paket *[NMAX]Package, nPasien *int, nama string) {
	var idx int = -1
	var find bool = false
	var i int = 0
	for i < *nPasien && !find {
		if pasien[i].nama == nama {
			idx = i
			find = true
		}
		i++
	}

	if find {
		for i := idx; i < *nPasien-1; i++ {
			pasien[i] = pasien[i+1]
			paket[i] = paket[i+1]
		}
		*nPasien--
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
