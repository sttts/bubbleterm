package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/taigrr/bubbleterm/emulator"
)

func main() {
	// Create a new emulator
	emu, err := emulator.New(80, 24)
	if err != nil {
		log.Fatal(err)
	}
	defer emu.Close()

	// Start the user's shell so you can try any command
	cmd := exec.Command(defaultShell())
	err = emu.StartCommand(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// Wait a moment for output
	time.Sleep(1 * time.Second)

	// Get the screen
	frame := emu.GetScreen()

	fmt.Println("Terminal output:")
	for i, row := range frame.Rows {
		fmt.Printf("%2d: %s\n", i, row)
	}
	fmt.Println("resizing!")

	emu.Resize(100, 40)
	// Wait a moment for output after resizing
	time.Sleep(1 * time.Second)
	fmt.Println("Terminal output after resizing:")
	frame = emu.GetScreen()
	for i, row := range frame.Rows {
		fmt.Printf("%2d: %s\n", i, row)
	}
}

func defaultShell() string {
	if shell := os.Getenv("SHELL"); shell != "" {
		return shell
	}
	if runtime.GOOS == "windows" {
		if shell := os.Getenv("COMSPEC"); shell != "" {
			return shell
		}
		return "cmd.exe"
	}
	return "/bin/sh"
}
