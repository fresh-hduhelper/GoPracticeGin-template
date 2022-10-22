package main

import (
	"log"
	"syscall"
)

func main() {
	dll, err := syscall.LoadDLL("hash.dll")
	if err != nil {
		log.Fatal(err)
	}
	defer dll.Release()

	proc := dll.MustFindProc("getModifiedChaptersIndexes")

	proc.Call()
}

type CheckItem struct {
	Title string
}
