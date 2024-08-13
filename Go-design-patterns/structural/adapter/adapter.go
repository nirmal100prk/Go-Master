/*
Adapter is a structural design pattern, which allows incompatible objects
 to collaborate.
*/

package main

import "fmt"

type IComputer interface {
	ConnectToLightiningPort()
}

type Mac struct{}

func (m *Mac) ConnectToLightiningPort() {
	fmt.Println("Mac conected to port")
}

type Windows struct{}

func (w *Windows) ConnectToWindowsPort() {
	fmt.Println("Windows Connected to port")
}

type WindowsAdapter struct {
	Windows *Windows
}

func (w *WindowsAdapter) ConnectToLightiningPort() {
	fmt.Println("Windows using adapter")
	w.Windows.ConnectToWindowsPort()
}

func main() {
	var I IComputer
	I = &Mac{}

	I.ConnectToLightiningPort()
	I = &WindowsAdapter{}
	I.ConnectToLightiningPort()

}
