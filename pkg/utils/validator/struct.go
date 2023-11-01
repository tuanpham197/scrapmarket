package validator

type Error struct {
	Field string
	Tag   string
	Value string
	Error string
}
