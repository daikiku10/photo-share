package domain

type Id string

// go:generate accessor -type=PhotoShare
type PhotoShare struct {
	id    Id
	title string
}
