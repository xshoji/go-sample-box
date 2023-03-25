package sports

//go:generate go install golang.org/x/tools/cmd/stringer@latest
//go:generate stringer -type=Sports
type Sports int

const (
	Null Sports = iota
	Baseball
	Swimming
	Soccer
	Karate
)
