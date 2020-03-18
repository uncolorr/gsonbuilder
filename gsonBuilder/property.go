package gsonBuilder

type property struct {
	name string
	serializedName string
	propertyType string
}

// Set name and set first character to lowercase
func(property* property) setNameWithFormat(name string)  {
	property.name = lowerCaseFirst(name)
}

// Set type and set first character to uppercase
func(property* property) setTypeWithFormat(propertyType string)  {
	property.propertyType = upperCaseFirst(propertyType)
}