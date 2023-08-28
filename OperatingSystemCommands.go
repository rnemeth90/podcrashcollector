package podcrashcollector

import "os/exec"

type WindowsCommands struct{}

func (c *WindowsCommands) ParseHostName(path string) string {
	out, _ := exec.Command("powershell.exe", "your command here").Output()
	return string(out)
}

type LinuxCommands struct{}

func (c *LinuxCommands) ParseHostName(path string) string {
	// Similar pattern, but with Linux commands
}
