package main

type AppConfig struct {
	Title     string
	Width     int
	TopMargin int
}

var App = AppConfig{
	Title:     "My CLI App",
	Width:     140,
	TopMargin: 10,
}
