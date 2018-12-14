package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSAmazonMQConfigurationAssociation AWS CloudFormation Resource (AWS::AmazonMQ::ConfigurationAssociation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html
type AWSAmazonMQConfigurationAssociation struct {

	// Broker AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-broker
	Broker string `json:"Broker,omitempty"`

	// Configuration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html#cfn-amazonmq-configurationassociation-configuration
	Configuration *AWSAmazonMQConfigurationAssociation_ConfigurationId `json:"Configuration,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAmazonMQConfigurationAssociation) AWSCloudFormationType() string {
	return "AWS::AmazonMQ::ConfigurationAssociation"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAmazonMQConfigurationAssociation) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAmazonMQConfigurationAssociation) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAmazonMQConfigurationAssociation) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAmazonMQConfigurationAssociation) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAmazonMQConfigurationAssociation) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSAmazonMQConfigurationAssociation) MarshalJSON() ([]byte, error) {
	type Properties AWSAmazonMQConfigurationAssociation
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
func (r *AWSAmazonMQConfigurationAssociation) UnmarshalJSON(b []byte) error {
	type Properties AWSAmazonMQConfigurationAssociation
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
		Metadata   map[string]interface{}
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSAmazonMQConfigurationAssociation(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}

// GetAllAWSAmazonMQConfigurationAssociationResources retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
func (t *Template) GetAllAWSAmazonMQConfigurationAssociationResources() map[string]AWSAmazonMQConfigurationAssociation {
	results := map[string]AWSAmazonMQConfigurationAssociation{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSAmazonMQConfigurationAssociation:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AmazonMQ::ConfigurationAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAmazonMQConfigurationAssociation
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

// GetAWSAmazonMQConfigurationAssociationWithName retrieves all AWSAmazonMQConfigurationAssociation items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSAmazonMQConfigurationAssociationWithName(name string) (AWSAmazonMQConfigurationAssociation, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSAmazonMQConfigurationAssociation:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::AmazonMQ::ConfigurationAssociation" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSAmazonMQConfigurationAssociation
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSAmazonMQConfigurationAssociation{}, errors.New("resource not found")
}
