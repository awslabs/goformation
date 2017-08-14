package goformation

// Template represents an AWS CloudFormation template
type Template struct {
	Resources  map[string]interface{}
	Parameters map[string]interface{}
	Outputs    map[string]interface{}
}
