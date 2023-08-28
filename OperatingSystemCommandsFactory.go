package podcrashcollector

import "runtime"

func GetPlatformOperations() IOperatingSystemCommands {
	if runtime.GOOS == "windows" {
		return &WindowsCommands{}
	} else if runtime.GOOS == "linux" {
		return &LinuxCommands{}
	}
	panic("Unsupported OS platform. CoreDumpHandler currently supports Linux and Windows.")
}
