package mindflow

type AppConfig struct {
	Width int
}
type Option struct {
	Name string
}
type Options []Option

var App = AppConfig{
	Width: 150,
}
