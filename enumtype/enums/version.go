package enums

//go:generate go run golang.org/x/tools/cmd/stringer -type=Version
type Version int

const (
	VersionNull Version = iota
	Version0000
	Version2017
	Version2018
	Version2019
)
