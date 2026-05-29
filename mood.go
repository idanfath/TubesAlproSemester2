package main

import "fmt"

// mood adalah struktur data untuk menyimpan satu catatan mood
type mood struct {
	ID        int
	Tanggal   string
	Skor      int
	Deskripsi string
}

var moodlist [100]mood
var moodcount int = 0
var modid int = 1

// tambahMood: minta data dari user, lalu simpan ke moodlist
func tambahMood() {
	var m mood

	if moodcount >= 100 {
		fmt.Println("Penyimpanan mood penuh (maks 100 data)!")
		return
	}

	m.ID = modid
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&m.Tanggal)

	fmt.Print("Masukkan skor mood (1-10): ")
	fmt.Scan(&m.Skor)

	for m.Skor < 1 || m.Skor > 10 {
		fmt.Println("Skor mood harus antara 1 dan 10! Silakan coba lagi.")
		fmt.Print("silahkan Masukkan skor mood (1-10): ")
		fmt.Scan(&m.Skor)
	}

	fmt.Print("Deskripsi: ")
	fmt.Scan(&m.Deskripsi)

	moodlist[moodcount] = m
	moodcount++
	modid++

	fmt.Println("Catatan mood berhasil ditambahkan!")
}

// tampilMood: cetak semua catatan mood yang sudah ada dalam bentuk tabel
func tampilMood() {
	var i int

	if moodcount == 0 {
		fmt.Println("Belum ada catatan mood!")
		return
	}

	fmt.Println("\n+----+------------+------+--------------------------+")
	fmt.Println("| ID |  Tanggal   | Skor | Deskripsi                |")
	fmt.Println("+----+------------+------+--------------------------+")
	for i = 0; i < moodcount; i++ {
		m := moodlist[i]
		fmt.Printf("| %-2d | %-10s |  %-3d | %-24s |\n",
			m.ID, m.Tanggal, m.Skor, m.Deskripsi)
	}
	fmt.Println("+----+------------+------+--------------------------+")
}

// ubahMood: cari catatan berdasarkan ID, lalu perbarui isinya
func ubahMood() {
	var i, idx int
	if moodcount == 0 {
		fmt.Println("Belum ada catatan mood!")
		return
	}

	tampilMood()

	var targetID int
	fmt.Print("Masukkan ID catatan mood yang ingin diubah: ")
	fmt.Scan(&targetID)

	idx = -1
	i = 0
	for i < moodcount && idx == -1 {
		if moodlist[i].ID == targetID {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	fmt.Printf("Data lama → Tanggal: %s | Skor: %d | Deskripsi: %s\n",
		moodlist[idx].Tanggal, moodlist[idx].Skor, moodlist[idx].Deskripsi)
	fmt.Println("Masukkan data baru:")

	fmt.Print("Tanggal baru (YYYY-MM-DD): ")
	fmt.Scan(&moodlist[idx].Tanggal)

	fmt.Print("Skor baru (1-10): ")
	fmt.Scan(&moodlist[idx].Skor)
	if moodlist[idx].Skor < 1 || moodlist[idx].Skor > 10 {
		fmt.Println("Skor mood harus antara 1 dan 10!")
		return
	}

	fmt.Print("Deskripsi baru: ")
	fmt.Scan(&moodlist[idx].Deskripsi)

	fmt.Println("Catatan mood berhasil diubah!")
}

// hapusMood: cari catatan berdasarkan ID, lalu hapus dengan cara geser elemen
func hapusMood() {
	var i, idx int
	if moodcount == 0 {
		fmt.Println("Belum ada catatan mood!")
		return
	}

	tampilMood()

	var targetID int
	fmt.Print("Masukkan ID catatan mood yang ingin dihapus: ")
	fmt.Scan(&targetID)

	idx = -1
	i = 0
	for i < moodcount && idx == -1 {
		if moodlist[i].ID == targetID {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	var konfirmasi string
	fmt.Printf("Yakin hapus mood [%s | Skor: %d]? (y/n): ",
		moodlist[idx].Tanggal, moodlist[idx].Skor)
	fmt.Scan(&konfirmasi)

	if konfirmasi != "y" && konfirmasi != "Y" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	for i := idx; i < moodcount-1; i++ {
		moodlist[i] = moodlist[i+1]
	}
	moodlist[moodcount-1] = mood{}
	moodcount--

	fmt.Println("Catatan mood berhasil dihapus!")
}

// menuMood: sub-menu khusus pengelolaan catatan mood
func menuMood() {
	var pilihan int
	for {
		fmt.Println("\n+---------------------------+")
		fmt.Println("|    MENU CATATAN MOOD      |")
		fmt.Println("+---------------------------+")
		fmt.Println("| 1. Tambah Catatan Mood    |")
		fmt.Println("| 2. Tampil Semua Mood      |")
		fmt.Println("| 3. Ubah Catatan Mood      |")
		fmt.Println("| 4. Hapus Catatan Mood     |")
		fmt.Println("| 0. Kembali ke Menu Utama  |")
		fmt.Println("+---------------------------+")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahMood()
		case 2:
			tampilMood()
		case 3:
			ubahMood()
		case 4:
			hapusMood()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
