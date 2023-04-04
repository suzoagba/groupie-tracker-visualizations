package handler

import (
	"log"
	"os/exec"
	"runtime"
)

// Define a function called Open that takes a string argument path
func Open(path string) { // https://github.com/0x434D53/openinbrowser
	// Create an empty slice of strings to store the command arguments
	var args []string
	// Use a switch statement to choose the appropriate command for the OS
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path} // Use the 'open' command on macOS
	case "windows":
		args = []string{"cmd", "/c", "start", path} // Use 'cmd /c start' on Windows
	default:
		args = []string{"xdg-open", path} // Use 'xdg-open' on Linux and other Unix-like systems
	}

	// Create a new command using the chosen arguments and path
	cmd := exec.Command(args[0], args[1:]...)
	// Run the command and capture any errors
	err := cmd.Run()
	if err != nil {
		log.Printf("openinbrowser: %v\n", err) // Log any errors using the log package
	}
}
