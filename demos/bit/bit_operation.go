package main

type Flags int8

const (
	READ Flags = 1 << iota
	WRITE
	EXECUTE
	FILE
	DIRECTORY
)

func IsRead(flags Flags) bool {
	return flags&READ == READ
}
func IsWrite(flags Flags) bool {
	return flags&WRITE == WRITE
}
func IsExecute(flags Flags) bool {
	return flags&EXECUTE == EXECUTE
}
func IsFile(flags Flags) bool {
	return flags&FILE == FILE
}
func IsDirectory(flags Flags) bool {
	return flags&DIRECTORY == DIRECTORY
}
