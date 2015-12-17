package tern

//Op - Ternary operator
func Op(cond bool, t, f interface{}) interface{} {
	if cond {
		return t
	}
	return f
}
