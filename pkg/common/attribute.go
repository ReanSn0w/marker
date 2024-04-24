package common

const (
	EmptyValue Value = ""
)

type (
	attributes map[Key]Value
	Key        string
	Value      string
)
