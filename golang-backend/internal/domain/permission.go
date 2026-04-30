package domain

type Permission uint64

const (
	Create Permission = 1 << iota
	Update
	Delete
	View
)
