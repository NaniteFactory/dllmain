package main

import "C"

import (
	"unsafe"

	"github.com/nanitefactory/winmb"
)

//export Test
func Test() {
	winmb.MessageBoxPlain("export Test", "export Test")
}

// OnProcessAttach is an async callback (hook).
//export OnProcessAttach
func OnProcessAttach(
	hinstDLL unsafe.Pointer, // handle to DLL module
	fdwReason uint32, // reason for calling function
	lpReserved unsafe.Pointer, // reserved
) {
	winmb.MessageBoxPlain("OnProcessAttach", "OnProcessAttach")
}

const title = "TITLE"

var version = "undefined"

func main() {
	// nothing really. xD
}
