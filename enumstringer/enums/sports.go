package sports

//go:generate stringer -type=Sports
type Sports int

const (
	_ Sports = iota
	Baseball
	Swimming
	Soccer
	Karate
)
