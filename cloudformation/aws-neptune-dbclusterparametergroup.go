package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSNeptuneDBClusterParameterGroup AWS CloudFormation Resource (AWS::Neptune::DBClusterParameterGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html
type AWSNeptuneDBClusterParameterGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html#cfn-neptune-dbclusterparametergroup-description
	Description *Value `json:"Description,omitempty"`

	// Family AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html#cfn-neptune-dbclusterparametergroup-family
	Family *Value `json:"Family,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html#cfn-neptune-dbclusterparametergroup-name
	Name *Value `json:"Name,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html#cfn-neptune-dbclusterparametergroup-parameters
	Parameters interface{} `json:"Parameters,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html#cfn-neptune-dbclusterparametergroup-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSNeptuneDBClusterParameterGroup) AWSCloudFormationType() string {
	return "AWS::Neptune::DBClusterParameterGroup"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSNeptuneDBClusterParameterGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSNeptuneDBClusterParameterGroup
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSNeptuneDBClusterParameterGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSNeptuneDBClusterParameterGroup
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSNeptuneDBClusterParameterGroup(*res.Properties)
	}

	return nil
}

// GetAllAWSNeptuneDBClusterParameterGroupResources retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
func (t *Template) GetAllAWSNeptuneDBClusterParameterGroupResources() map[string]AWSNeptuneDBClusterParameterGroup {
	results := map[string]AWSNeptuneDBClusterParameterGroup{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSNeptuneDBClusterParameterGroup:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Neptune::DBClusterParameterGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSNeptuneDBClusterParameterGroup{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSNeptuneDBClusterParameterGroupWithName retrieves all AWSNeptuneDBClusterParameterGroup items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSNeptuneDBClusterParameterGroupWithName(name string) (AWSNeptuneDBClusterParameterGroup, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSNeptuneDBClusterParameterGroup:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Neptune::DBClusterParameterGroup" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSNeptuneDBClusterParameterGroup{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSNeptuneDBClusterParameterGroup{}, errors.New("resource not found")
}
