package cloudformation

// Parameter enables you to input custom values to your template each time you create or update a stack
type Parameter struct {
	AllowedPattern string `json:"AllowedPattern,omitempty"`
	AllowedValues []interface{} `json:"AllowedValues,omitempty"`
	ConstraintDescription string `json:"ConstraintDescription,omitempty"`
	Default interface{} `json:"Default,omitempty"`
	Description string `json:"Description,omitempty"`
	MaxLength int `json:"MaxLength,omitempty"`
	MinLength int `json:"MinLength,omitempty"`
	MaxValue int `json:"MaxValue,omitempty"`
	MinValue int `json:"MinValue,omitempty"`
	NoEcho bool `json:"NoEcho,omitempty"`
	Type Type `json:"Type"`
}

// Type is the data type for the parameter
type Type = string

const (
	// StringType is the type of literal strings
	StringType = Type("String")
	// NummberType is the type of integers or floats
	NumberType = Type("Number")
	// ListNumberType is the type of array of integers or floats that are separated by commas
	ListNumberType = Type("List<Number>")
	// CommaDelimitedListType is the type of arrays of literal strings that are separated by commas
	CommaDelimitedListType = Type("CommaDelimitedList")

	// TODO Add AWS-Specific Parameter Types
	// Ref https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html#aws-specific-parameter-types
)
