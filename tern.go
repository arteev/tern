package tern


//Op - Ternary operator. If the condition cond is true, it returns t, otherwise f
func Op(cond bool, t, f interface{}) interface{} {
	if cond {
		return t
	}
	return f
}

//Decode compares value to each search value one by one. If value is equal to a search, then returns the corresponding result.
//If no match is found, then returns default. If default is omitted, then returns nil.
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

//Nvl returns first value if it is not nil,otherwise default
func Nvl(value interface{},defvalue interface{}) interface{} {
	if value==nil {
		return defvalue
	}
	return value
}