package main

import "fmt"

func carimood() {
	var keyword string
	fmt.Print("Masukkan tanggal atau deskripsi mood yang dicari: ")
	fmt.Scan(&keyword)
	sequentialSearchMood(keyword)
}

// sequentialSearchMood: mencari mood berdasarkan tanggal ATAU deskripsi
func sequentialSearchMood(keyword string) {
	found := false

	fmt.Println("\n[Hasil Pencarian Mood — Sequential Search]")
	fmt.Println("+----+------------+------+--------------------------+")
	fmt.Println("| ID |  Tanggal   | Skor | Deskripsi                |")
	fmt.Println("+----+------------+------+--------------------------+")

	i := 0
	for i < moodcount {
		m := moodlist[i]
		if m.Tanggal == keyword || m.Deskripsi == keyword {
			fmt.Printf("| %-2d | %-10s |  %-3d | %-24s |\n",
				m.ID, m.Tanggal, m.Skor, m.Deskripsi)
			found = true
		}
		i++
	}

	fmt.Println("+----+------------+------+--------------------------+")
	if !found {
		fmt.Println("Tidak ada catatan mood yang cocok.")
	}
}

func caritask() {
	var keyword string
	fmt.Print("Masukkan nama tugas yang dicari: ")
	fmt.Scan(&keyword)

	// Langkah 1: Salin tasklist ke array sementara agar data asli tidak berubah
	var sortedTasks [100]task
	for i := 0; i < taskcount; i++ {
		sortedTasks[i] = tasklist[i]
	}

	// Langkah 2: Urutkan array sementara berdasarkan nama (Insertion Sort ascending)
	i := 1
	for i <= taskcount-1 {
		temp := sortedTasks[i]
		j := i
		for j > 0 && temp.nama < sortedTasks[j-1].nama {
			sortedTasks[j] = sortedTasks[j-1]
			j = j - 1
		}
		sortedTasks[j] = temp
		i = i + 1
	}

	binarySearchTask(keyword, sortedTasks, taskcount)
}

// binarySearchTask: mencari tugas berdasarkan nama menggunakan Binary Search
func binarySearchTask(keyword string, a [100]task, n int) {
	kr := 0
	kn := n - 1
	found := -1

	// Loop sampai rentang habis (kr > kn) atau data ditemukan (found != -1)
	for kr <= kn && found == -1 {
		med := (kr + kn) / 2

		if a[med].nama < keyword {

			kr = med + 1
		} else if a[med].nama > keyword {

			kn = med - 1
		} else {
			found = med
		}
	}

	fmt.Println("\n[Hasil Pencarian Tugas — Binary Search]")
	if found != -1 {
		t := a[found]
		fmt.Println("+----+------------------+----------+--------+------------------+")
		fmt.Println("| ID | Nama Tugas       | Prioritas| Durasi | Status           |")
		fmt.Println("+----+------------------+----------+--------+------------------+")
		fmt.Printf("| %-2d | %-16s | %-8s | %4d m | %-16s |\n",
			t.id, t.nama, labelPrioritas(t.prioritas), t.durasi, labelSelesai(t.selesai))
		fmt.Println("+----+------------------+----------+--------+------------------+")
	} else {
		fmt.Println("Tidak ada tugas dengan nama tersebut.")
	}
}
