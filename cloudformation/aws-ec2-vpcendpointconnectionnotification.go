package cloudformation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2VPCEndpointConnectionNotification AWS CloudFormation Resource (AWS::EC2::VPCEndpointConnectionNotification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html
type AWSEC2VPCEndpointConnectionNotification struct {

	// ConnectionEvents AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-connectionevents
	ConnectionEvents []string `json:"ConnectionEvents,omitempty"`

	// ConnectionNotificationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-connectionnotificationarn
	ConnectionNotificationArn string `json:"ConnectionNotificationArn,omitempty"`

	// ServiceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-serviceid
	ServiceId string `json:"ServiceId,omitempty"`

	// VPCEndpointId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointconnectionnotification.html#cfn-ec2-vpcendpointconnectionnotification-vpcendpointid
	VPCEndpointId string `json:"VPCEndpointId,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCEndpointConnectionNotification) AWSCloudFormationType() string {
	return "AWS::EC2::VPCEndpointConnectionNotification"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2VPCEndpointConnectionNotification) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2VPCEndpointConnectionNotification) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2VPCEndpointConnectionNotification) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2VPCEndpointConnectionNotification) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2VPCEndpointConnectionNotification) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSEC2VPCEndpointConnectionNotification) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2VPCEndpointConnectionNotification
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string               `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{} `json:"Metadata,omitempty"`
		DeletionPolicy DeletionPolicy         `json:"DeletionPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r._dependsOn,
		Metadata:       r._metadata,
		DeletionPolicy: r._deletionPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2VPCEndpointConnectionNotification) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2VPCEndpointConnectionNotification
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
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
		*r = AWSEC2VPCEndpointConnectionNotification(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}

// GetAllAWSEC2VPCEndpointConnectionNotificationResources retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPCEndpointConnectionNotificationResources() map[string]*AWSEC2VPCEndpointConnectionNotification {
	results := map[string]*AWSEC2VPCEndpointConnectionNotification{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2VPCEndpointConnectionNotification:
			results[name] = &resource
		case *AWSEC2VPCEndpointConnectionNotification:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPCEndpointConnectionNotification" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPCEndpointConnectionNotification
						if err := json.Unmarshal(b, &result); err == nil {
							t.Resources[name] = &result
							results[name] = &result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2VPCEndpointConnectionNotificationWithName retrieves all AWSEC2VPCEndpointConnectionNotification items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPCEndpointConnectionNotificationWithName(name string) (*AWSEC2VPCEndpointConnectionNotification, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2VPCEndpointConnectionNotification:
			return &resource, nil
		case *AWSEC2VPCEndpointConnectionNotification:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPCEndpointConnectionNotification" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2VPCEndpointConnectionNotification
						if err := json.Unmarshal(b, &result); err == nil {
							t.Resources[name] = &result
							return &result, nil
						}
					}
				}
			}
		}
	}
	return nil, errors.New("resource not found")
}
