package gsonBuilder

import (
	"bytes"
	"encoding/json"
	"reflect"
)

type GsonClassBuilder struct {

}

// Parse json data and generate classes as string
// It returns all generated classes as one string or error if json format is invalid

func (builder *GsonClassBuilder) Parse(jsonData string) (string, error){
	m, err := builder.convertToMap([]byte(jsonData))
	if err != nil {
		return "", err
	}
	var classes []class

	builder.parseMapData(m, "Root", &classes)

	buf := bytes.Buffer{}
	for _ , class := range classes {
		gsonClass := builder.generateGsonClass(class)
		buf.WriteString(gsonClass)
	}
	result := buf.String()
	return result, nil
}

// Parse every json value that is not a primitive type
// It returns all generated classes as one string or error if json format is invalid
func (builder *GsonClassBuilder) parseMapData(m map[string]interface{}, className string, classes* []class){
	var class class
	class.setNameWithFormat(className)
	var properties []property

	for k, v := range m {
		var property property
		if reflect.TypeOf(v) == reflect.TypeOf(make(map[string]interface{})){
			property.setTypeWithFormat(k)
			property.serializedName = k
			property.setNameWithFormat(k)
			properties = append(properties, property)
			builder.parseMapData(v.(map[string]interface{}), k, classes)
			continue
		}

		t := TypeOf(v)

		property.setNameWithFormat(k)
		property.setTypeWithFormat(t)
		property.serializedName = k
		properties = append(properties, property)
	}

	class.properties = properties

	*classes = append(*classes, class)
}

// Convert class struct to Gson class
// Return Gson class as string
func (builder *GsonClassBuilder) generateGsonClass(class class) string  {
	buf := bytes.Buffer{}
	buf.WriteString("data")
	buf.WriteString(" ")
	buf.WriteString("class ")
	buf.WriteString(class.name)
	buf.WriteString("(\n")
	for _ , property := range class.properties {
		buf.WriteString(builder.generateGsonProperty(property))
	}
	buf.WriteString(")\n")
	result := buf.String()
	return result
}

// Convert property struct to Gson property for Gson class
// Return Gson property as string
 func (*GsonClassBuilder) generateGsonProperty(property property) string {
	buf := bytes.Buffer{}
	buf.WriteString("\t")
	buf.WriteString("@SerializedName(\"")
	buf.WriteString(property.serializedName)
	buf.WriteString("\")")
	buf.WriteString("\n")
	buf.WriteString("\t")
	buf.WriteString("val ")
	buf.WriteString(property.name)
	buf.WriteString(": ")
	buf.WriteString(property.propertyType)
	buf.WriteString("\n")
	result := buf.String()
	return result
}

// Convert json data to map and check is json format valid
// Return json as map or error if json format invalid
func (*GsonClassBuilder) convertToMap(data []byte) (map[string]interface{}, error){
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
