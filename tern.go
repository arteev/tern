package tern


//Op - Ternary operator
func Op(cond bool, t, f interface{}) interface{} {
	if cond {
		return t
	}
	return f
}

func Decode(value interface{},val1,val2 interface{}, vals... interface{} ) interface{} {
	values := []interface{}{val1,val2,}
	values = append(values,vals...)
	l:=len(values)
	i:=0;
	for i<l {
		if value == values[i] {
			if i+1<=l {
				return values[i+1]
			}
		}
		if i+1==l {
			return values[i]
		}
		i+=2
	}
	return nil
}