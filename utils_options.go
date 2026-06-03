package main

import "fmt"

// return options paling panjang
func longestOptionName(options Options) int {
	var max int = 0
	var i int
	var optionsAmount int = len(options)
	for i = 0; i < optionsAmount; i++ {
		if len(options[i].name) > max {
			max = len(options[i].name)
		}
	}
	return max
}

func buildOptions(options Options) Options {
	var result Options
	var i int
	var page = getPage(App.currentPage)
	if len(App.history) > 0 && !page.noBack {
		result = append(result, Option{
			name: "Back",
			action: func() {
				back()
			},
		})
	}
	if page.hasTable {
		if (Temp.table_page * App.pagination) < page.dataCount() {
			result = append(result,
				Option{name: "Halaman Berikutnya", action: func() {
					Temp.table_page++
				}})
		}
		if Temp.table_page > 1 {
			result = append(result,
				Option{name: "Halaman Sebelumnya", action: func() {
					Temp.table_page--
				}})
		}
		if page.dataCount() > App.pagination {
			// totalHalaman = total data / pagination, kalo ada sisa, tambah 1
			var totalHalaman = page.dataCount()/App.pagination + ifi(page.dataCount()%App.pagination == 0, 0, 1)
			result = append(result,
				Option{name: "Buka Halaman", action: func() {
					var page int
					inputNumber(InputNumber{
						prompt: fmt.Sprintf("Masukkan nomor halaman (1-%d): ", totalHalaman),
						onSubmit: func(f float64) {
							page = int(f)
						},
					})
					if page >= 1 && page <= totalHalaman {
						Temp.table_page = page
					} else {
						alert("Nomor halaman tidak valid!")
					}
				}},
			)
		}
	}
	for i = 0; i < len(options); i++ {
		result = append(result, options[i])
	}
	return result
}

func runOption(options Options, index int) {
	if index < 0 || index >= len(options) || options[index].action == nil {
		return
	}
	options[index].action()
}

func readOption(options Options) {
	var index int
	fmt.Print(cen("Select an option: "))
	fmt.Scan(&index)
	runOption(options, index)
}
