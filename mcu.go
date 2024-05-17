package main
import "fmt"

const NMAX int = 1024

type paket struct {
	jenis string
	harga int
}

type pasien struct {
	nama string
	umur int
	asal string
}
type pkg [NMAX]paket
type pas [NMAX]pasien

func main() {
	fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
	fmt.Printf("1. Penambahan Data\n2. Pengubahan Data\n3. Penghapusan Data\n4. Laporan Pemasukan\n5. Pencarian Pasien\n6. Tampilkan Data Pasien Terurut\n")
	fmt.Printf("\nPilih Menu: ")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
		case 1:
            fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Berapa pasien yang ingin ditambahkan? ")
			var nPasien int
			fmt.Scan(&nPasien)
			
        case 2:
            fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin merubah data pasien atas nama siapa? ")
			var nama string
			fmt.Scan(&nama)
        case 3:
            fmt.Printf("===============================\n   Aplikasi Medical Check-Up   \n===============================\n")
			fmt.Println("Ingin menghapus data pasien atas nama siapa? ")
			var nama string
			fmt.Scan(&nama)
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
			fmt.Printf("Ingin menampilakan data pasien terurut berdasarkan apa?\n1. Waktu MCU\n2. Paket MCU\n")
			var nama string
			fmt.Scan(&nama)
		default:
            fmt.Printf("Menu tidak tersedia\n")
	}
}