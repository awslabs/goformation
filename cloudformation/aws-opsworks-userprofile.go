package cloudformation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// AWSOpsWorksUserProfile AWS CloudFormation Resource (AWS::OpsWorks::UserProfile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html
type AWSOpsWorksUserProfile struct {

	// AllowSelfManagement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-allowselfmanagement
	AllowSelfManagement bool `json:"AllowSelfManagement,omitempty"`

	// IamUserArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-iamuserarn
	IamUserArn string `json:"IamUserArn,omitempty"`

	// SshPublicKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshpublickey
	SshPublicKey string `json:"SshPublicKey,omitempty"`

	// SshUsername AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshusername
	SshUsername string `json:"SshUsername,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksUserProfile) AWSCloudFormationType() string {
	return "AWS::OpsWorks::UserProfile"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSOpsWorksUserProfile) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSOpsWorksUserProfile) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSOpsWorksUserProfile) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSOpsWorksUserProfile) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSOpsWorksUserProfile) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSOpsWorksUserProfile) MarshalJSON() ([]byte, error) {
	type Properties AWSOpsWorksUserProfile
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
func (r *AWSOpsWorksUserProfile) UnmarshalJSON(b []byte) error {
	type Properties AWSOpsWorksUserProfile
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
		*r = AWSOpsWorksUserProfile(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}

// GetAllAWSOpsWorksUserProfileResources retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksUserProfileResources() map[string]*AWSOpsWorksUserProfile {
	results := map[string]*AWSOpsWorksUserProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSOpsWorksUserProfile:
			results[name] = &resource
		case *AWSOpsWorksUserProfile:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::UserProfile" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksUserProfile
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

// GetAWSOpsWorksUserProfileWithName retrieves all AWSOpsWorksUserProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksUserProfileWithName(name string) (*AWSOpsWorksUserProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSOpsWorksUserProfile:
			return &resource, nil
		case *AWSOpsWorksUserProfile:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::UserProfile" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksUserProfile
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
