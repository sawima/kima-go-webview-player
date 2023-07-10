package main

import "C"
import "github.com/webview/webview"

func main() {

	url := "http://localhost:8888/"

	debug := true
	w := webview.New(debug)
	window := w.Window()

	defer w.Destroy()

	C.gtk_window_fullscreen((*C.GtkWindow)(window))

	w.SetTitle("Vending Machine")
	w.SetSize(480, 800, webview.HintNone)
	w.Navigate(url)
	w.Run()
}
