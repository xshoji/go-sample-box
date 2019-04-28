package sports

//go:generate stringer -type=Sports
type Sports int

const (
	Null Sports = iota
	Baseball
	Swimming
	Soccer
	Karate
)
