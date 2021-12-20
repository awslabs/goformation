package cdkmetadata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CDKMetadata AWS CloudFormation Resource (AWS::CDK::Metadata)
type CDKMetadata struct {
	Analytics string `json:"Analytics,omitempty"`

	Metadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *CDKMetadata) AWSCloudFormationType() string {
	return "AWS::CDK::Metadata"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r CDKMetadata) MarshalJSON() ([]byte, error) {
	type Properties CDKMetadata
	return json.Marshal(&struct {
		Type       string
		Properties Properties
		Metadata   map[string]interface{} `json:"Metadata,omitempty"`
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(r),
		Metadata:   r.Metadata,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *CDKMetadata) UnmarshalJSON(b []byte) error {
	type Properties CDKMetadata
	res := &struct {
		Type       string
		Properties *Properties
		Metadata   map[string]interface{}
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = CDKMetadata(*res.Properties)
	}
	return nil
}
