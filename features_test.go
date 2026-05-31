package main

import (
	"testing"
)

func TestValidateMoodDate(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"2024-06-01", true},
		{"2024-02-29", true},  // leap year
		{"2025-02-29", false}, // non-leap year
		{"24-06-01", false},
		{"2024/06/01", false},
		{"abcd-ef-gh", false},
		{"2024-13-01", false}, // invalid month
		{"2024-06-32", false}, // invalid day
		{"", false},
	}

	for _, tt := range tests {
		actual := validateMoodDate(tt.input)
		if actual != tt.expected {
			t.Errorf("validateMoodDate(%q) = %v; want %v", tt.input, actual, tt.expected)
		}
	}
}

func TestMoodCRUD(t *testing.T) {
	// Backup original state
	originalMoods := moods
	defer func() { moods = originalMoods }()

	// Reset moods to a known state
	moods = []Mood{
		{id: 1, description: "Hari yang menyenangkan!", score: 8, date: "2024-06-01"},
		{id: 2, description: "Agak stres", score: 2, date: "2024-06-02"},
	}

	// 1. Test findMoodIndex
	if findMoodIndex(1) != 0 {
		t.Errorf("Expected index 0 for mood ID 1, got %d", findMoodIndex(1))
	}
	if findMoodIndex(99) != -1 {
		t.Errorf("Expected index -1 for non-existent mood ID 99, got %d", findMoodIndex(99))
	}

	// 2. Test getMaxMoodID
	if getMaxMoodID() != 2 {
		t.Errorf("Expected max mood ID to be 2, got %d", getMaxMoodID())
	}

	// 3. Test insertMood
	newMood := Mood{description: "Merasa sangat bahagia!", score: 9, date: "2024-06-03"}
	insertMood(newMood)
	if len(moods) != 3 {
		t.Fatalf("Expected 3 moods after insertion, got %d", len(moods))
	}
	inserted := moods[2]
	if inserted.id != 3 || inserted.description != "Merasa sangat bahagia!" || inserted.score != 9 || inserted.date != "2024-06-03" {
		t.Errorf("Inserted mood values are incorrect: %+v", inserted)
	}

	// 4. Test updateMood
	updatedMood := Mood{id: 2, description: "Sedikit lebih baik", score: 4, date: "2024-06-02"}
	updateMood(2, updatedMood)
	if moods[1].description != "Sedikit lebih baik" || moods[1].score != 4 {
		t.Errorf("Mood update failed: %+v", moods[1])
	}

	// 5. Test deleteMood
	deleteMood(1)
	if len(moods) != 2 {
		t.Fatalf("Expected 2 moods after deletion, got %d", len(moods))
	}
	if findMoodIndex(1) != -1 {
		t.Error("Deleted mood still exists in index")
	}
}

func TestTaskCRUD(t *testing.T) {
	// Backup original state
	originalTasks := tasks
	defer func() { tasks = originalTasks }()

	// Reset tasks to a known state
	tasks = []Task{
		{id: 1, title: "Belajar Go", priority: 1, duration: 120, isCompleted: false},
		{id: 2, title: "Olahraga", priority: 2, duration: 60, isCompleted: true},
	}

	// 1. Test findTaskIndex
	if findTaskIndex(2) != 1 {
		t.Errorf("Expected index 1 for task ID 2, got %d", findTaskIndex(2))
	}

	// 2. Test getMaxTaskID
	if getMaxTaskID() != 2 {
		t.Errorf("Expected max task ID 2, got %d", getMaxTaskID())
	}

	// 3. Test insertTask
	newTask := Task{title: "Membaca Buku", priority: 3, duration: 90, isCompleted: false}
	insertTask(newTask)
	if len(tasks) != 3 {
		t.Fatalf("Expected 3 tasks after insertion, got %d", len(tasks))
	}
	inserted := tasks[2]
	if inserted.id != 3 || inserted.title != "Membaca Buku" || inserted.priority != 3 || inserted.duration != 90 {
		t.Errorf("Inserted task values are incorrect: %+v", inserted)
	}

	// 4. Test updateTask
	updatedTask := Task{id: 1, title: "Belajar Go Advanced", priority: 1, duration: 180, isCompleted: true}
	updateTask(1, updatedTask)
	if tasks[0].title != "Belajar Go Advanced" || tasks[0].priority != 1 || !tasks[0].isCompleted || tasks[0].duration != 180 {
		t.Errorf("Task update failed: %+v", tasks[0])
	}

	// 5. Test deleteTask
	deleteTask(2)
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks after deletion, got %d", len(tasks))
	}
	if findTaskIndex(2) != -1 {
		t.Error("Deleted task still exists in index")
	}
}

func TestGetLabelPriorityTask(t *testing.T) {
	tests := []struct {
		priority int
		expected string
	}{
		{1, "Tinggi"},
		{2, "Sedang"},
		{3, "Rendah"},
		{0, "Tidak Diketahui"},
		{4, "Tidak Diketahui"},
		{-1, "Tidak Diketahui"},
	}

	for _, tt := range tests {
		actual := getLabelPriorityTask(tt.priority)
		if actual != tt.expected {
			t.Errorf("getLabelPriorityTask(%d) = %q; want %q", tt.priority, actual, tt.expected)
		}
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		y, m, d  int
		expected bool
	}{
		{2024, 6, 1, true},
		{2024, 2, 29, true},
		{2025, 2, 29, false},
		{2024, 2, 30, false},
		{2024, 13, 1, false},
		{2024, 0, 1, false},
		{-1, 6, 1, false},
		{2024, 6, 31, false}, // June only has 30 days
		{2024, 7, 31, true},  // July has 31 days
	}

	for _, tt := range tests {
		actual := validateDate(tt.y, tt.m, tt.d)
		if actual != tt.expected {
			t.Errorf("validateDate(%d, %d, %d) = %v; want %v", tt.y, tt.m, tt.d, actual, tt.expected)
		}
	}
}

func TestMinuteToTime(t *testing.T) {
	tests := []struct {
		minutes  int
		expected string
	}{
		{0, "0 Menit"},
		{45, "45 Menit"},
		{60, "1 Jam "}, // note trailing space from simple implementation
		{120, "2 Jam "},
		{75, "1 Jam 15 Menit"},
		{135, "2 Jam 15 Menit"},
	}

	for _, tt := range tests {
		actual := minuteToTime(tt.minutes)
		if actual != tt.expected {
			t.Errorf("minuteToTime(%d) = %q; want %q", tt.minutes, actual, tt.expected)
		}
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"0", 0},
		{"-456", -456},
		{"12a34", 0}, // Invalid digits
		{"", 0},
	}

	for _, tt := range tests {
		actual := toInt(tt.input)
		if actual != tt.expected {
			t.Errorf("toInt(%q) = %d; want %d", tt.input, actual, tt.expected)
		}
	}
}

func TestIsNumStr(t *testing.T) {
	tests := []struct {
		input          string
		expectNegative bool
		expected       bool
	}{
		{"123", false, true},
		{"123", true, true},
		{"-123", true, true},
		{"-123", false, false},
		{"-", true, false},
		{"12a3", false, false},
		{"", false, false},
	}

	for _, tt := range tests {
		actual := isNumStr(tt.input, tt.expectNegative)
		if actual != tt.expected {
			t.Errorf("isNumStr(%q, %v) = %v; want %v", tt.input, tt.expectNegative, actual, tt.expected)
		}
	}
}

func TestStringSwitch(t *testing.T) {
	if stringswitch(true, "yes", "no") != "yes" {
		t.Error("stringswitch(true) failed")
	}
	if stringswitch(false, "yes", "no") != "no" {
		t.Error("stringswitch(false) failed")
	}
}

func TestLowerAndUpper(t *testing.T) {
	if lower("Hello WORLD!") != "hello world!" {
		t.Errorf("lower() failed: got %q", lower("Hello WORLD!"))
	}
	if upper("Hello world!") != "HELLO WORLD!" {
		t.Errorf("upper() failed: got %q", upper("Hello world!"))
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input     string
		maxLen    int
		expected  string
	}{
		{"testing", 0, ""},
		{"testing", -1, "testing"},
		{"testing", 10, "testing"},
		{"testing", 3, "tes"},
		{"testing", 5, "te..."},
	}

	for _, tt := range tests {
		actual := truncate(tt.input, tt.maxLen)
		if actual != tt.expected {
			t.Errorf("truncate(%q, %d) = %q; want %q", tt.input, tt.maxLen, actual, tt.expected)
		}
	}
}
