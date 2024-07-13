package desktop_gui

import (
	"image/color"
   "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func Desktop() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")
   class_switch := true 

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
   content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)
   button := widget.NewButton("switch horizontal/vertical align", func() {
      class_switch = !class_switch
      fmt.Println(class_switch)
       
   })
   button2 := widget.NewButton("switch horizontal/vertical align", func() {
      class_switch = !class_switch
      fmt.Println(class_switch)
       
   })

   button_content := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, button2, layout.NewSpacer())

   if class_switch {
      myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered, button_content))
   }
   myWindow.SetContent(container.New(layout.NewHBoxLayout(), content, centered, button_content))
      
   /*
   switch class_switch{
   case true:
      myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered, button_content))
   case false:
      myWindow.SetContent(container.New(layout.NewHBoxLayout(), content, centered, button_content))
   }
   */
   myWindow.Resize(fyne.NewSize(480,360))
	myWindow.ShowAndRun()
}
