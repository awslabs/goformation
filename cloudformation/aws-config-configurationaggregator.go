package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSConfigConfigurationAggregator AWS CloudFormation Resource (AWS::Config::ConfigurationAggregator)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html
type AWSConfigConfigurationAggregator struct {

	// AccountAggregationSources AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-accountaggregationsources
	AccountAggregationSources []AWSConfigConfigurationAggregator_AccountAggregationSource `json:"AccountAggregationSources,omitempty"`

	// ConfigurationAggregatorName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-configurationaggregatorname
	ConfigurationAggregatorName *Value `json:"ConfigurationAggregatorName,omitempty"`

	// OrganizationAggregationSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-organizationaggregationsource
	OrganizationAggregationSource *AWSConfigConfigurationAggregator_OrganizationAggregationSource `json:"OrganizationAggregationSource,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigurationAggregator) AWSCloudFormationType() string {
	return "AWS::Config::ConfigurationAggregator"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSConfigConfigurationAggregator) MarshalJSON() ([]byte, error) {
	type Properties AWSConfigConfigurationAggregator
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
func (r *AWSConfigConfigurationAggregator) UnmarshalJSON(b []byte) error {
	type Properties AWSConfigConfigurationAggregator
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
		*r = AWSConfigConfigurationAggregator(*res.Properties)
	}

	return nil
}

// GetAllAWSConfigConfigurationAggregatorResources retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigConfigurationAggregatorResources() map[string]AWSConfigConfigurationAggregator {
	results := map[string]AWSConfigConfigurationAggregator{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSConfigConfigurationAggregator:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::ConfigurationAggregator" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSConfigConfigurationAggregator{}
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

// GetAWSConfigConfigurationAggregatorWithName retrieves all AWSConfigConfigurationAggregator items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigurationAggregatorWithName(name string) (AWSConfigConfigurationAggregator, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSConfigConfigurationAggregator:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::ConfigurationAggregator" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSConfigConfigurationAggregator{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSConfigConfigurationAggregator{}, errors.New("resource not found")
}
