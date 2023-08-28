package podcrashcollector

type IOperatingSystemCommands interface {
	ParseServiceName(path string) string
	ParseCreationDateTime(path string) string
	ParseHostName(path string) string
}

type ICloudStorageUploader interface {
	UploadFile(osCommands IOperatingSystemCommands, path string, shouldOverwrite bool) bool
}
