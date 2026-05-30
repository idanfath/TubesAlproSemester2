package main

import "fmt"

type task struct {
	id        int
	nama      string
	prioritas int
	durasi    int
	selesai   bool
}

var tasklist [100]task
var taskcount int = 0
var taskid int = 1

func labelPrioritas(p int) string {
	switch p {
	case 1:
		return "Tinggi"
	case 2:
		return "Sedang"
	case 3:
		return "Rendah"
	default:
		return "Tidak Diketahui"
	}
}

func labelSelesai(s bool) string {
	if s {
		return "Selesai"
	}
	return "Belum Selesai"
}

func tambahtask() {
	var t task

	if taskcount >= 100 {
		fmt.Println("Penyimpanan tugas penuh (maks 100 data)!")
		return
	}

	t.id = taskid
	t.selesai = false

	fmt.Print("Nama Tugas: ")
	fmt.Scan(&t.nama)

	fmt.Print("Prioritas (1=Tinggi, 2=Sedang, 3=Rendah): ")
	fmt.Scan(&t.prioritas)

	if t.prioritas < 1 || t.prioritas > 3 {
		fmt.Println("Prioritas harus antara 1 dan 3!")
		return
	}

	fmt.Print("Durasi (menit): ")
	fmt.Scan(&t.durasi)

	tasklist[taskcount] = t
	taskcount++
	taskid++

	fmt.Println("Tugas berhasil ditambahkan!")
}

func tampilTask() {
	if taskcount == 0 {
		fmt.Println("Belum ada tugas!")
		return
	}

	fmt.Println("\n+----+------------------+----------+--------+------------------+")
	fmt.Println("| ID | Nama Tugas       | Prioritas| Durasi | Status           |")
	fmt.Println("+----+------------------+----------+--------+------------------+")

	for i := 0; i < taskcount; i++ {
		t := tasklist[i]
		fmt.Printf("| %-2d | %-16s | %-8s | %4d m | %-16s |\n",
			t.id, t.nama, labelPrioritas(t.prioritas), t.durasi, labelSelesai(t.selesai))
	}

	fmt.Println("+----+------------------+----------+--------+------------------+")
	fmt.Printf("Total: %d tugas\n", taskcount)
}

func ubahtask() {
	var i, idx int

	if taskcount == 0 {
		fmt.Println("Belum ada tugas!")
		return
	}

	tampilTask()

	var targetID int
	fmt.Print("Masukkan ID tugas yang ingin diubah: ")
	fmt.Scan(&targetID)

	idx = -1
	i = 0
	for i < taskcount && idx == -1 {
		if tasklist[i].id == targetID {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	fmt.Printf("Data lama → Nama: %s | Prioritas: %s | Durasi: %d menit | Status: %s\n",
		tasklist[idx].nama, labelPrioritas(tasklist[idx].prioritas),
		tasklist[idx].durasi, labelSelesai(tasklist[idx].selesai))

	fmt.Println("Masukkan data baru:")

	fmt.Print("Nama Tugas baru: ")
	fmt.Scan(&tasklist[idx].nama)

	fmt.Print("Prioritas baru (1=Tinggi, 2=Sedang, 3=Rendah): ")
	fmt.Scan(&tasklist[idx].prioritas)

	if tasklist[idx].prioritas < 1 || tasklist[idx].prioritas > 3 {
		fmt.Println("Prioritas harus antara 1 dan 3!")
		return
	}

	fmt.Print("Durasi baru (menit): ")
	fmt.Scan(&tasklist[idx].durasi)

	var selesaiInput string
	fmt.Print("Apakah tugas sudah selesai? (y/n): ")
	fmt.Scan(&selesaiInput)
	tasklist[idx].selesai = (selesaiInput == "y" || selesaiInput == "Y")

	fmt.Println("Tugas berhasil diubah!")
}

func hapustask() {
	var i, idx int

	if taskcount == 0 {
		fmt.Println("Belum ada tugas!")
		return
	}

	tampilTask()

	var targetID int
	fmt.Print("Masukkan ID tugas yang ingin dihapus: ")
	fmt.Scan(&targetID)

	idx = -1
	i = 0
	for i < taskcount && idx == -1 {
		if tasklist[i].id == targetID {
			idx = i
		}
		i++
	}

	if idx == -1 {
		fmt.Println("ID tidak ditemukan!")
		return
	}

	var konfirmasi string
	fmt.Printf("Yakin hapus tugas [%s | Prioritas: %s]? (y/n): ",
		tasklist[idx].nama, labelPrioritas(tasklist[idx].prioritas))
	fmt.Scan(&konfirmasi)

	if konfirmasi != "y" && konfirmasi != "Y" {
		fmt.Println("Hapus tugas dibatalkan.")
		return
	}

	for i := idx; i < taskcount-1; i++ {
		tasklist[i] = tasklist[i+1]
	}
	taskcount--
	tasklist[taskcount] = task{}

	fmt.Println("Tugas berhasil dihapus!")
}
