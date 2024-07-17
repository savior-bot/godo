package desktop_gui

import (
	//	"image/color"
	//   "fmt"
//	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	//	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)


func Desktop() {
	myApp := app.New()
	main_window := myApp.NewWindow("Box Layout")
   
   // my first window goes here 
   main_content := widget.NewLabel(" main")
   profile := widget.NewButton("profile", func() {}) 
   folders := container.NewAppTabs(
      container.NewTabItem("tab1", widget.NewLabel("new tabe")),
      container.NewTabItem("tab2", widget.NewLabel("second tabe")),
   )
   side_bar := container.NewVBox(
      profile, 
      folders,
   )
   content := container.NewHSplit(side_bar, main_content)
   content.Offset = .25
   main_window.SetContent(content)
   main_window.Resize(fyne.NewSize(9999,9999))
   main_window.SetMaster()
	main_window.ShowAndRun()
}
