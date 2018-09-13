package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServiceDiscoveryInstance AWS CloudFormation Resource (AWS::ServiceDiscovery::Instance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html
type AWSServiceDiscoveryInstance struct {

	// InstanceAttributes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html#cfn-servicediscovery-instance-instanceattributes
	InstanceAttributes interface{} `json:"InstanceAttributes,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html#cfn-servicediscovery-instance-instanceid
	InstanceId *Value `json:"InstanceId,omitempty"`

	// ServiceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicediscovery-instance.html#cfn-servicediscovery-instance-serviceid
	ServiceId *Value `json:"ServiceId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServiceDiscoveryInstance) AWSCloudFormationType() string {
	return "AWS::ServiceDiscovery::Instance"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServiceDiscoveryInstance) MarshalJSON() ([]byte, error) {
	type Properties AWSServiceDiscoveryInstance
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
func (r *AWSServiceDiscoveryInstance) UnmarshalJSON(b []byte) error {
	type Properties AWSServiceDiscoveryInstance
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
		*r = AWSServiceDiscoveryInstance(*res.Properties)
	}

	return nil
}

// GetAllAWSServiceDiscoveryInstanceResources retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
func (t *Template) GetAllAWSServiceDiscoveryInstanceResources() map[string]AWSServiceDiscoveryInstance {
	results := map[string]AWSServiceDiscoveryInstance{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServiceDiscoveryInstance:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceDiscovery::Instance" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceDiscoveryInstance{}
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

// GetAWSServiceDiscoveryInstanceWithName retrieves all AWSServiceDiscoveryInstance items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServiceDiscoveryInstanceWithName(name string) (AWSServiceDiscoveryInstance, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServiceDiscoveryInstance:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ServiceDiscovery::Instance" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSServiceDiscoveryInstance{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSServiceDiscoveryInstance{}, errors.New("resource not found")
}
