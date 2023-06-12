package main

import "fmt"

type PC interface {
	InsertToLightningPort()
}

type MAC struct{}

func (m *MAC) InsertToLightningPort() {
	fmt.Println("Inserted to lightning port")
}

// incompatible class
type Windows struct{}

func (w *Windows) InsertToUSBPort() {
	fmt.Println("Inserted to USB Port")
}

// windows adapter class
type WindowsAdapter struct {
	win *Windows
}

func (w *WindowsAdapter) InsertToLightningPort() {
	fmt.Println("Adapted converting lightning port to USB port")
	w.win.InsertToUSBPort()
}

func main() {
	mac := MAC{}
	mac.InsertToLightningPort()

	winAdapter := WindowsAdapter{}
	winAdapter.InsertToLightningPort()
}
