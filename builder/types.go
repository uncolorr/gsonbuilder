package builder

import "math"

// Determines the type of json value
// Return name of primitive type as string
func TypeOf(v interface{}) string {
	switch v.(type) {
	case int:
		return "Int"
	case float64:
		_, b := math.Modf(v.(float64))
		if b == 0 {
			return "Long"
		}
		return "Double"
	case bool:
		return "Boolean"
	case string:
		return "String"
	default:
		return "Any"
	}
}
