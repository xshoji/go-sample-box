package sports

//go:generate go run golang.org/x/tools/cmd/stringer -type=Sports
type Sports int

const (
	Null Sports = iota
	Baseball
	Swimming
	Soccer
	Karate
)
