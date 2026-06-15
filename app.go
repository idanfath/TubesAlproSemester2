// disini simpan konfigurasi, data awal aplikasi, dll

package main

type AppConfig struct {
	title        string
	width        int
	topMargin    int
	bottomMargin int
}

var App = AppConfig{
	title:        "Mindflow",
	width:        120,
	topMargin:    2,
	bottomMargin: 5,
}

var SIGKILL bool = false

func exit() {
	SIGKILL = true
	return
}
