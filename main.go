package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(formatted)
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	clock := widget.NewLabel("")
	updateTime(clock)
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
	fmt.Println("exited")

}
