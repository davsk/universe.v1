package main

import (
    "github.com/andlabs/ui"
)

func main() {
    err := ui.Main(func() {
        name := ui.NewEntry()
        button := ui.NewButton("Greet")
        greeting := ui.NewLabel("")
		checkbox := ui.NewCheckbox("Checkbox")
        box := ui.NewVerticalBox()
        box.Append(ui.NewLabel("Enter your name:"), false)
        box.Append(name, false)
        box.Append(button, false)
        box.Append(greeting, false)
		box.Append(checkbox, false)
        window := ui.NewWindow("Universe", 200, 100, false)
        window.SetChild(box)
        button.OnClicked(func(*ui.Button) {
            greeting.SetText("Hello, " + name.Text() + "!")
        })
        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })
        window.Show()
    })
    if err != nil {
        panic(err)
    }
}