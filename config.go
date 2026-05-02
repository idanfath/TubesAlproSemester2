package main

type AppConfig struct {
	Title     string
	Width     int
	TopMargin int
}
type Option struct {
	Name   string
	Action func()
}
type Options []Option

var App = AppConfig{
	Title:     "My CLI App",
	Width:     160,
	TopMargin: 10,
}
