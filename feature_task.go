package main

import "fmt"

type Task struct {
	id          int
	title       string
	priority    int // 1, 2, 3
	duration    int // menit, estimasi
	isCompleted bool
}

var tasks = []Task{
	{id: 1, title: "Belajar Go", priority: 1, duration: 120, isCompleted: false},
	{id: 2, title: "Olahraga", priority: 2, duration: 60, isCompleted: true},
	{id: 3, title: "Membaca Buku", priority: 3, duration: 90, isCompleted: false},
}

var task_priority = [3]string{"Tinggi", "Sedang", "Rendah"}

func getLabelPriorityTask(priority int) string {
	if priority < 1 || priority > 3 {
		return "Tidak Diketahui"
	}
	return task_priority[priority-1]
}

func TablePriorityTask() Table {
	var table Table
	table.header = []string{"Prioritas", "Label"}
	table.maxLengths = []int{10, 20}
	var i int
	for i = 1; i <= 3; i++ {
		table.rows = append(table.rows, []string{
			toString(i),
			getLabelPriorityTask(i),
		})
	}
	return table
}

func TaskTable() Table {
	var table Table
	table.header = []string{"ID", "Deskripsi", "Prioritas", "Durasi", "Status"}
	table.maxLengths = []int{5, 40, 10, 20, 20}
	var i int
	for i = 0; i < len(tasks); i++ {
		table.rows = append(table.rows, []string{
			toString(tasks[i].id),
			tasks[i].title,
			getLabelPriorityTask(tasks[i].priority),
			minuteToTime(tasks[i].duration),
			stringswitch(tasks[i].isCompleted, "Selesai", "Belum Selesai"),
		})
	}
	return table
}

func insertTask(task Task) {
	task.id = getMaxTaskID() + 1
	tasks = append(tasks, task)
}

func deleteTask(id int) {
	var i int
	var newTasks []Task
	for i = 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			continue
		}
		newTasks = append(newTasks, tasks[i])
	}
	tasks = newTasks
}

func updateTask(id int, newTask Task) {
	var index = findTaskIndex(id)
	if index == -1 {
		return
	}
	tasks[index] = newTask
}

func getMaxTaskID() int {
	var max_index int
	var i int
	for i = 0; i < len(tasks); i++ {
		if tasks[i].id > max_index {
			max_index = tasks[i].id
		}
	}
	return max_index
}

// --

func appOptionTaskView() {
	var id int
	inputNumber(InputNumber{
		Prompt: "Masukkan ID catatan tugas yang ingin dilihat: ",
		onSubmit: func(input float64) {
			id = int(input)
		},
	})
	if !v(findTaskIndex(id) != -1, "ID tidak ditemukan!") {
		return
	}
	Temp.id = id
	toPage("TaskView")
}

func appOptionTaskInsert() {
	var task Task
	input(Input{
		Prompt: "Judul tugas: ",
		onSubmit: func(input string) {
			task.title = input
		},
	})
	if !v(task.title != "", "Judul tugas tidak boleh kosong!") {
		return
	}

	printTable(TablePriorityTask())
	inputNumber(InputNumber{
		Prompt: "Prioritas tugas (1-3): ",
		onSubmit: func(input float64) {
			task.priority = int(input)
		},
	})
	if !v(task.priority >= 1 && task.priority <= 3, "Prioritas harus antara 1 dan 3!") {
		return
	}

	inputNumber(InputNumber{
		Prompt: "Durasi tugas (menit): ",
		onSubmit: func(input float64) {
			task.duration = int(input)
		},
	})
	if !v(task.duration >= 0, "Durasi harus bernilai positif!") {
		return
	}

	input(Input{
		Prompt: "Tandai tugas sebagai selesai? (y/n): (n) ",
		onSubmit: func(input string) {
			if input == "y" || input == "Y" {
				task.isCompleted = true
			}
		},
	})

	alert("Task berhasil ditambahkan!")
	insertTask(task)
}

func appOptionTaskUpdate() {
	var id int
	var temp string
	if App.currentPage == "TaskView" && Temp.id != -1 {
		id = Temp.id
	} else {
		inputNumber(InputNumber{
			Prompt: "Masukkan ID catatan tugas yang ingin diubah: ",
			onSubmit: func(input float64) {
				id = int(input)
			},
		})
	}
	var idx = findTaskIndex(id)
	if !v(idx != -1, "ID tidak ditemukan!") {
		return
	}

	var task = tasks[idx]

	input(Input{
		Prompt: "Judul tugas: (" + task.title + ") ",
		onSubmit: func(input string) {
			task.title = input
		},
	})
	if task.title == "" {
		task.title = tasks[idx].title
	}

	printTable(TablePriorityTask())
	input(Input{
		Prompt: "Prioritas tugas baru (1-3): (" + getLabelPriorityTask(task.priority) + ") ",
		onSubmit: func(input string) {
			if input == "" {
				task.priority = tasks[idx].priority
				return
			}
			task.priority = toInt(input)
		},
	})

	if !v(task.priority >= 1 && task.priority <= 3, "Prioritas harus antara 1 dan 3!") {
		return
	}

	input(Input{
		Prompt: "Durasi tugas baru (menit): (" + toString(task.duration) + ") ",
		onSubmit: func(input string) {
			temp = input

		},
	})
	if temp == "" {
		task.duration = tasks[idx].duration
	} else {
		if !v(isNumStr(temp, false), "Durasi harus berupa angka!") {
			return
		}
		if !v(toInt(temp) >= 0, "Durasi harus bernilai positif!") {
			return
		}
		task.duration = toInt(temp)
	}

	input(Input{
		Prompt: "Tandai tugas sebagai selesai? (y/n): (" + stringswitch(task.isCompleted, "y", "n") + ") ",
		onSubmit: func(input string) {
			if lower(input) == "y" {
				task.isCompleted = true
			} else if lower(input) == "n" {
				task.isCompleted = false
			} else {
				task.isCompleted = tasks[idx].isCompleted
			}
		},
	})

	updateTask(id, task)
	alert("Tugas berhasil diubah!")
}

func appOptionTaskDelete() {
	var id int
	if App.currentPage == "TaskView" && Temp.id != -1 {
		id = Temp.id
	} else {
		inputNumber(InputNumber{
			Prompt: "Masukkan ID catatan tugas yang ingin dihapus: ",
			onSubmit: func(input float64) {
				id = int(input)
			},
		})
	}
	if !v(findTaskIndex(id) != -1, "ID tidak ditemukan!") {
		return
	}
	deleteTask(id)
	alert("Tugas berhasil dihapus!")
	if App.currentPage == "TaskView" {
		toPage("Task")
	}
}

func appOptionToggleTaskStatus() {
	var id int
	if App.currentPage != "TaskView" || Temp.id == -1 {
		inputNumber(InputNumber{
			Prompt: "Masukkan ID catatan tugas yang ingin diubah statusnya: ",
			onSubmit: func(input float64) {
				id = int(input)
			},
		})
	} else {
		id = Temp.id
	}
	var index = findTaskIndex(id)
	if !v(index != -1, "ID tidak ditemukan!") {
		return
	}
	tasks[index].isCompleted = !tasks[index].isCompleted
	updateTask(Temp.id, tasks[index])
	alert("Status tugas berhasil diperbarui!")
	toPage("TaskView") // refresh page
}

func appDynamicTaskView() []string {
	var index = findTaskIndex(Temp.id)
	if index == -1 {
		return []string{"ID tidak ditemukan!"}
	}
	var task = tasks[index]
	return []string{
		fmt.Sprintf("%d | Prioritas %s | Estimasi %s | %s", task.id, getLabelPriorityTask(task.priority), minuteToTime(task.duration), stringswitch(task.isCompleted, "Selesai", "Belum Selesai")),
		task.title,
	}
}

//--

func findTaskIndex(id int) int {
	var i int
	for i = 0; i < len(tasks); i++ {
		if tasks[i].id == id {
			return i
		}
	}
	return -1
}
