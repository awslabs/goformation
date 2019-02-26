package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAppStreamImageBuilder AWS CloudFormation Resource (AWS::AppStream::ImageBuilder)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html
type AWSAppStreamImageBuilder struct {

	// AppstreamAgentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-appstreamagentversion
	AppstreamAgentVersion string `json:"AppstreamAgentVersion,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-description
	Description string `json:"Description,omitempty"`

	// DisplayName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-displayname
	DisplayName string `json:"DisplayName,omitempty"`

	// DomainJoinInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-domainjoininfo
	DomainJoinInfo *AWSAppStreamImageBuilder_DomainJoinInfo `json:"DomainJoinInfo,omitempty"`

	// EnableDefaultInternetAccess AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-enabledefaultinternetaccess
	EnableDefaultInternetAccess bool `json:"EnableDefaultInternetAccess,omitempty"`

	// ImageArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagearn
	ImageArn string `json:"ImageArn,omitempty"`

	// ImageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagename
	ImageName string `json:"ImageName,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-instancetype
	InstanceType string `json:"InstanceType,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-name
	Name string `json:"Name,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-vpcconfig
	VpcConfig *AWSAppStreamImageBuilder_VpcConfig `json:"VpcConfig,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppStreamImageBuilder) AWSCloudFormationType() string {
	return "AWS::AppStream::ImageBuilder"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppStreamImageBuilder) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSAppStreamImageBuilder) MarshalJSON() ([]byte, error) {
	type Properties AWSAppStreamImageBuilder
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DeletionPolicy DeletionPolicy `json:"DeletionPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DeletionPolicy: r._deletionPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSAppStreamImageBuilder) UnmarshalJSON(b []byte) error {
	type Properties AWSAppStreamImageBuilder
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
		*r = AWSAppStreamImageBuilder(*res.Properties)
	}

	return nil
}

// GetAllAWSAppStreamImageBuilderResources retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
func (t *Template) GetAllAWSAppStreamImageBuilderResources() map[string]AWSAppStreamImageBuilder {
	results := map[string]AWSAppStreamImageBuilder{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAppStreamImageBuilder:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppStream::ImageBuilder" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAppStreamImageBuilder
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSAppStreamImageBuilderWithName retrieves all AWSAppStreamImageBuilder items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAppStreamImageBuilderWithName(name string) (AWSAppStreamImageBuilder, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAppStreamImageBuilder:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AppStream::ImageBuilder" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAppStreamImageBuilder
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAppStreamImageBuilder{}, errors.New("resource not found")
}
