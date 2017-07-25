package canvas

// NullRune is a rune that can be null.
type NullRune struct {
	value rune
	valid bool
}

// Set sets the value to the given rune.
func (r *NullRune) Set(value rune) {
	r.value, r.valid = value, true
}

// Clear sets the value to null
func (r *NullRune) Clear() {
	r.valid = false
}

// Value returns the value.
// If the second return value is false, the rune should be treated a null.
func (r *NullRune) Value() (rune, bool) {
	return r.value, r.valid
}
