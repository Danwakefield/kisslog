package kisslog

import "fmt"

// Attrs is a key-value type used to pass structured logging data to Logger instances.
type Attrs map[string]interface{}

func (a *Attrs) Pretty() string {
	if a == nil || len(*a) == 0 {
		return ""
	}

	var result string
	for key, val := range *a {
		result = fmt.Sprintf("%s %s=%v", result, key, val)
	}

	return result
}

// SplitAttrs checks if the last item passed in v is an Attrs instance,
// if so it returns it separately. If not, v is returned as-is with a nil Attrs.
func splitAttrs(v ...interface{}) ([]interface{}, *Attrs) {
	if len(v) == 0 {
		return v, nil
	}

	attrs, ok := v[len(v)-1].(Attrs)

	if !ok {
		return v, nil
	}

	v = v[:len(v)-1]
	return v, &attrs
}
