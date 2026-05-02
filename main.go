package main

var SigKill bool = false

func main() {
	for {
		render()
		if SigKill {
			render()
			break
		}
	}
}

func exitApp() {
	SigKill = true
	toPage("Exit")
}
