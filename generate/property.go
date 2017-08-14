package main

// Property represents an AWS CloudFormation resource property
type Property struct {

	// Documentation - A link to the AWS CloudFormation User Guide that provides information about the property.
	Documentation string `json:"Documentation"`

	// DuplicatesAllowed - If the value of the Type field is List, indicates whether AWS CloudFormation allows duplicate values.
	// If the value is true, AWS CloudFormation ignores duplicate values. If the value is false,
	// AWS CloudFormation returns an error if you submit duplicate values.
	DuplicatesAllowed bool `json:"DuplicatesAllowed"`

	// ItemType - If the value of the Type field is List or Map, indicates the type of list or map if they contain
	// non-primitive types. Otherwise, this field is omitted. For lists or maps that contain primitive
	// types, the PrimitiveItemType property indicates the valid value type.
	//
	// A subproperty name is a valid item type. For example, if the type value is List and the item type
	//  value is PortMapping, you can specify a list of port mapping properties.
	ItemType string `json:"ItemType"`

	// PrimitiveItemType - If the value of the Type field is List or Map, indicates the type of list or map
	// if they contain primitive types. Otherwise, this field is omitted. For lists or maps that contain
	// non-primitive types, the ItemType property indicates the valid value type.
	// The valid primitive types for lists and maps are String, Long, Integer, Double, Boolean, or Timestamp.
	// For example, if the type value is List and the item type value is String, you can specify a list of strings
	// for the property. If the type value is Map and the item type value is Boolean, you can specify a string
	// to Boolean mapping for the property.
	PrimitiveItemType string `json:"PrimitiveItemType"`

	// PrimitiveType - For primitive values, the valid primitive type for the property. A primitive type is a
	// basic data type for resource property values.
	// The valid primitive types are String, Long, Integer, Double, Boolean, Timestamp or Json.
	// If valid values are a non-primitive type, this field is omitted and the Type field indicates the valid value type.
	PrimitiveType string `json:"PrimitiveType"`

	// Required indicates whether the property is required.
	Required bool `json:"Required"`

	// Type - For non-primitive types, valid values for the property. The valid types are a subproperty name,
	// List or Map. If valid values are a primitive type, this field is omitted and the PrimitiveType field
	// indicates the valid value type. A list is a comma-separated list of values. A map is a set of key-value pairs,
	// where the keys are always strings. The value type for lists and maps are indicated by the ItemType
	// or PrimitiveItemType field.
	Type string `json:"Type"`

	// UpdateType - During a stack update, the update behavior when you add, remove, or modify the property.
	// AWS CloudFormation replaces the resource when you change Immutable properties. AWS CloudFormation doesn't
	// replace the resource when you change mutable properties. Conditional updates can be mutable or immutable,
	// depending on, for example, which other properties you updated. For more information, see the relevant
	// resource type documentation.
	UpdateType string `json:"UpdateType"`
}

// IsPrimitive checks whether a property is a primitive type
func (p Property) IsPrimitive() bool {

	if p.PrimitiveType != "" {
		return true
	}

	if p.PrimitiveItemType != "" {
		return true
	}

	return false
}

// IsMap checks whether a property should be a map (map[string]...)
func (p Property) IsMap() bool {
	return p.Type == "Map"
}

// IsList checks whether a property should be a list ([]...)
func (p Property) IsList() bool {
	return p.Type == "List"
}

// PropertyType determins which Go type the property should be represented as
func (p Property) PropertyType() string {

	if p.PrimitiveType != "" {
		// This is a primitive type
		return getPrimitiveType(p.PrimitiveType)
	}

	if p.ItemType != "" {
		// This is a custom type for a list/map
		return p.ItemType
	}

	if p.PrimitiveItemType != "" {
		// This is a primitive for a list/map
		return getPrimitiveType(p.PrimitiveItemType)
	}

	// This must be a custom type
	return p.Type

}

// JSONSchemaPropertyType determins which JSON Schema type the property should be represented as
func (p Property) JSONSchemaPropertyType() string {

	if p.PrimitiveType != "" {
		// This is a primitive type
		return getJSONSchemaPrimitiveType(p.PrimitiveType)
	}

	switch p.Type {
	case "List":
	case "Map":
	default:
		return p.Type
	}

	// This must be a custom type
	return p.Type
}

func getPrimitiveType(pt string) string {
	switch pt {
	case "String":
		return "string"
	case "Long":
		return "int64"
	case "Integer":
		return "int64"
	case "Double":
		return "float64"
	case "Boolean":
		return "bool"
	case "Timestamp":
		return "time.Time"
	case "Json":
		return "object"
	default:
		return ""
	}
}

func getJSONSchemaPrimitiveType(pt string) string {
	switch pt {
	case "String":
		return "string"
	case "Long":
		return "number"
	case "Integer":
		return "number"
	case "Double":
		return "number"
	case "Boolean":
		return "boolean"
	case "Timestamp":
		return "string"
	case "Json":
		return "object"
	default:
		return ""
	}
}
